import { readFile, readdir, writeFile } from 'node:fs/promises';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

import ts from 'typescript';

const scriptDirectory = path.dirname(fileURLToPath(import.meta.url));
const frontendDirectory = path.resolve(scriptDirectory, '..');
const sourceDirectory = path.join(frontendDirectory, 'src');
const baselinePath = path.join(scriptDirectory, 'ui-english-baseline.txt');
const updateBaseline = process.argv.includes('--update');

const visibleAttributes = new Set([
    'aria-label',
    'cancelText',
    'confirmText',
    'description',
    'filterPlaceholder',
    'label',
    'placeholder',
    'sheetTitle',
    'title',
    'tooltip',
]);
const visibleProperties = new Set([
    'description',
    'errorMessage',
    'label',
    'message',
    'placeholder',
    'title',
    'tooltip',
]);
const visibleVariables = new Set([
    'description',
    'emptyText',
    'errorMessage',
    'label',
    'message',
    'placeholder',
    'title',
]);

const normalize = (value) => value.replaceAll(/\s+/g, ' ').trim();
const containsEnglish = (value) => /[A-Za-z]{2,}/.test(value);

const collectFiles = async (directory) => {
    const entries = await readdir(directory, { withFileTypes: true });
    const files = [];

    for (const entry of entries) {
        const fullPath = path.join(directory, entry.name);

        if (entry.isDirectory()) {
            files.push(...(await collectFiles(fullPath)));
        } else if (
            entry.name.endsWith('.tsx') &&
            !entry.name.endsWith('.test.tsx') &&
            !entry.name.endsWith('.spec.tsx')
        ) {
            files.push(fullPath);
        }
    }

    return files;
};

const expressionStrings = (node) => {
    if (ts.isStringLiteral(node) || ts.isNoSubstitutionTemplateLiteral(node)) {
        return [node.text];
    }

    if (ts.isTemplateExpression(node)) {
        return [node.getText()];
    }

    if (ts.isConditionalExpression(node)) {
        return [...expressionStrings(node.whenTrue), ...expressionStrings(node.whenFalse)];
    }

    if (ts.isParenthesizedExpression(node)) {
        return expressionStrings(node.expression);
    }

    return [];
};

const propertyName = (node) => {
    if (!node) {
        return undefined;
    }

    if (ts.isIdentifier(node) || ts.isStringLiteral(node)) {
        return node.text;
    }

    return undefined;
};

const scanFile = async (filePath) => {
    const sourceText = await readFile(filePath, 'utf8');
    const sourceFile = ts.createSourceFile(filePath, sourceText, ts.ScriptTarget.Latest, true, ts.ScriptKind.TSX);
    const relativePath = path.relative(frontendDirectory, filePath).replaceAll('\\', '/');
    const findings = [];

    const addFinding = (kind, value) => {
        const text = normalize(value);

        if (text && containsEnglish(text)) {
            findings.push(`${relativePath}\t${kind}\t${text}`);
        }
    };

    const visit = (node) => {
        if (ts.isJsxText(node)) {
            addFinding('jsx-text', node.text);
        } else if (ts.isJsxAttribute(node) && visibleAttributes.has(node.name.text)) {
            if (node.initializer && ts.isStringLiteral(node.initializer)) {
                addFinding(`attribute:${node.name.text}`, node.initializer.text);
            } else if (node.initializer && ts.isJsxExpression(node.initializer) && node.initializer.expression) {
                for (const value of expressionStrings(node.initializer.expression)) {
                    addFinding(`attribute:${node.name.text}`, value);
                }
            }
        } else if (ts.isJsxExpression(node) && node.expression && !ts.isJsxAttribute(node.parent)) {
            for (const value of expressionStrings(node.expression)) {
                addFinding('jsx-expression', value);
            }
        } else if (ts.isPropertyAssignment(node) && visibleProperties.has(propertyName(node.name))) {
            for (const value of expressionStrings(node.initializer)) {
                addFinding(`property:${propertyName(node.name)}`, value);
            }
        } else if (
            ts.isVariableDeclaration(node) &&
            ts.isIdentifier(node.name) &&
            visibleVariables.has(node.name.text) &&
            node.initializer
        ) {
            for (const value of expressionStrings(node.initializer)) {
                addFinding(`variable:${node.name.text}`, value);
            }
        } else if (ts.isCallExpression(node) && ts.isPropertyAccessExpression(node.expression)) {
            const receiver = node.expression.expression.getText(sourceFile);
            const method = node.expression.name.text;

            if (receiver === 'toast' && ['error', 'info', 'success', 'warning'].includes(method) && node.arguments[0]) {
                for (const value of expressionStrings(node.arguments[0])) {
                    addFinding(`toast:${method}`, value);
                }
            }
        }

        ts.forEachChild(node, visit);
    };

    visit(sourceFile);

    return findings;
};

const files = await collectFiles(sourceDirectory);
const currentFindings = [...new Set((await Promise.all(files.map(scanFile))).flat())].sort();

if (updateBaseline) {
    const content = [
        '# Generated by `pnpm i18n:baseline`.',
        '# Each entry is: relative path<TAB>detection kind<TAB>English text.',
        '# Remove entries by localizing the UI; do not add new entries without reviewing them.',
        ...currentFindings,
        '',
    ].join('\n');

    await writeFile(baselinePath, content, 'utf8');
    console.log(`Updated UI English baseline with ${currentFindings.length} findings.`);
    process.exit(0);
}

const baselineText = await readFile(baselinePath, 'utf8');
const baselineFindings = baselineText
    .split(/\r?\n/)
    .map((line) => line.trim())
    .filter((line) => line && !line.startsWith('#'))
    .sort();
const baselineSet = new Set(baselineFindings);
const currentSet = new Set(currentFindings);
const added = currentFindings.filter((finding) => !baselineSet.has(finding));
const removed = baselineFindings.filter((finding) => !currentSet.has(finding));

if (added.length === 0 && removed.length === 0) {
    console.log(`UI English baseline is current (${currentFindings.length} known findings).`);
    process.exit(0);
}

if (added.length > 0) {
    console.error('\nNew user-visible English requires translation or an explicit baseline review:');
    added.forEach((finding) => console.error(`+ ${finding}`));
}

if (removed.length > 0) {
    console.error('\nBaseline entries no longer found; regenerate it after confirming the localization changes:');
    removed.forEach((finding) => console.error(`- ${finding}`));
}

console.error('\nRun `pnpm i18n:baseline` only after reviewing every reported change.');
process.exit(1);
