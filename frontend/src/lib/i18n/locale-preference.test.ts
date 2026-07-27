import { describe, expect, it } from 'vitest';

import { resolveLocalePreference } from './locale-preference';

describe('resolveLocalePreference', () => {
    it('uses a stored locale before browser preferences', () => {
        expect(resolveLocalePreference('en', ['zh-CN'])).toBe('en');
    });

    it('maps Chinese language variants to Simplified Chinese', () => {
        expect(resolveLocalePreference(null, ['zh-Hans-SG', 'en-US'])).toBe('zh-CN');
    });

    it('maps English variants to English', () => {
        expect(resolveLocalePreference(undefined, ['en-GB', 'zh-CN'])).toBe('en');
    });

    it('uses the project default for unsupported languages', () => {
        expect(resolveLocalePreference(null, ['fr-FR'])).toBe('zh-CN');
    });
});
