import type { ReactNode } from 'react';

import { createContext, useCallback, useEffect, useMemo, useState } from 'react';

import type { Locale, TranslationValues } from '@/lib/i18n';

import { defaultLocale, dictionaries, fallbackLocale, isLocale, translate } from '@/lib/i18n';

interface LocaleProviderState {
    locale: Locale;
    setLocale: (locale: Locale) => void;
    /** Translates `key`, interpolating `{name}` placeholders from `values`. */
    t: (key: string, values?: TranslationValues) => string;
}

export const LocaleProviderContext = createContext<LocaleProviderState>({
    locale: defaultLocale,
    setLocale: () => null,
    t: (key, values) => translate(dictionaries[fallbackLocale], dictionaries[fallbackLocale], key, values),
});

interface LocaleProviderProps {
    children: ReactNode;
    storageKey?: string;
}

/**
 * Resolves the initial locale: an explicit stored choice wins, otherwise we
 * infer from the browser and fall back to `defaultLocale` for anything we
 * don't ship a dictionary for.
 */
const resolveInitialLocale = (storageKey: string): Locale => {
    const stored = localStorage.getItem(storageKey);

    if (isLocale(stored)) {
        return stored;
    }

    const preferred = globalThis.navigator?.languages ?? [];

    for (const tag of preferred) {
        if (isLocale(tag)) {
            return tag;
        }

        // Match `zh`, `zh-Hans`, `zh-SG`, ... onto our Simplified Chinese bundle.
        if (tag.toLowerCase().startsWith('zh')) {
            return 'zh-CN';
        }

        if (tag.toLowerCase().startsWith('en')) {
            return 'en';
        }
    }

    return defaultLocale;
};

export function LocaleProvider({ children, storageKey = 'locale' }: LocaleProviderProps) {
    const [locale, setLocaleState] = useState<Locale>(() => resolveInitialLocale(storageKey));

    useEffect(() => {
        document.documentElement.lang = locale;
    }, [locale]);

    const setLocale = useCallback(
        (next: Locale) => {
            localStorage.setItem(storageKey, next);
            setLocaleState(next);
        },
        [storageKey],
    );

    const value = useMemo<LocaleProviderState>(
        () => ({
            locale,
            setLocale,
            t: (key, values) => translate(dictionaries[locale], dictionaries[fallbackLocale], key, values),
        }),
        [locale, setLocale],
    );

    return <LocaleProviderContext.Provider value={value}>{children}</LocaleProviderContext.Provider>;
}
