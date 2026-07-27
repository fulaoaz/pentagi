import type { Locale } from './types';

import { defaultLocale, isLocale } from './types';

export const LOCALE_STORAGE_KEY = 'locale';

export const resolveLocalePreference = (stored: unknown, preferredLanguages: readonly string[]): Locale => {
    if (isLocale(stored)) {
        return stored;
    }

    for (const tag of preferredLanguages) {
        if (isLocale(tag)) {
            return tag;
        }

        const normalized = tag.toLowerCase();

        if (normalized.startsWith('zh')) {
            return 'zh-CN';
        }

        if (normalized.startsWith('en')) {
            return 'en';
        }
    }

    return defaultLocale;
};

export const resolveBrowserLocale = (storageKey = LOCALE_STORAGE_KEY): Locale =>
    resolveLocalePreference(localStorage.getItem(storageKey), globalThis.navigator?.languages ?? []);
