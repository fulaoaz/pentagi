import { spawnSync } from "node:child_process";
import path from "node:path";
import { fileURLToPath } from "node:url";

const repositoryDirectory = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  "..",
);
const checkOnly = process.argv.includes("--check");
const remoteArgument = process.argv.find((argument) =>
  argument.startsWith("--remote="),
);
const branchArgument = process.argv.find((argument) =>
  argument.startsWith("--branch="),
);
const remote = remoteArgument?.slice("--remote=".length) || "origin";
const upstreamBranch = branchArgument?.slice("--branch=".length) || "main";
const upstreamRef = `${remote}/${upstreamBranch}`;

const run = (command, arguments_, options = {}) => {
  const result = spawnSync(command, arguments_, {
    cwd: options.cwd ?? repositoryDirectory,
    encoding: "utf8",
    shell: process.platform === "win32",
    stdio: options.capture ? "pipe" : "inherit",
  });

  if (result.status !== 0 && !options.allowFailure) {
    process.exit(result.status ?? 1);
  }

  return options.capture ? result.stdout.trim() : "";
};

const currentBranch = run("git", ["branch", "--show-current"], {
  capture: true,
});

if (currentBranch !== "zh-CN") {
  console.error(
    `Run this script from the zh-CN branch (current branch: ${currentBranch || "detached HEAD"}).`,
  );
  process.exit(1);
}

const worktreeStatus = run("git", ["status", "--porcelain"], { capture: true });

if (worktreeStatus) {
  console.error(
    "Commit or stash the current worktree changes before syncing upstream.",
  );
  process.exit(1);
}

run("git", ["fetch", remote, upstreamBranch, "--prune"]);

const behind = Number(
  run("git", ["rev-list", "--count", `HEAD..${upstreamRef}`], {
    capture: true,
  }),
);

if (behind === 0) {
  console.log(`zh-CN is up to date with ${upstreamRef}.`);
  process.exit(0);
}

console.log(`${upstreamRef} has ${behind} new commit(s).`);

if (checkOnly) {
  process.exit(0);
}

run("git", ["merge", "--no-ff", "--no-commit", upstreamRef]);

const i18nCheck = spawnSync("node", ["scripts/check-ui-i18n.mjs"], {
  cwd: path.join(repositoryDirectory, "frontend"),
  encoding: "utf8",
  shell: process.platform === "win32",
  stdio: "inherit",
});

if (i18nCheck.status !== 0) {
  console.error(
    "Upstream was merged without committing. Translate or review the reported UI text, then commit.",
  );
  process.exit(i18nCheck.status ?? 1);
}

run("git", ["commit", "--no-edit"]);
console.log(`Merged ${upstreamRef}; run the full test suite before pushing.`);
