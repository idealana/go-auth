package middleware

import (
    "strconv"
    "strings"

    "github.com/gofiber/fiber/v3"
)

const localeKey = "locale"
const defaultLocaleKey = "defaultLocale"

type LocaleConfig struct {
    SupportLocales []string
    DefaultLocale string
}

func Locale(cfg LocaleConfig) fiber.Handler {
    supported := make(map[string]bool, len(cfg.SupportLocales))
    for _, l := range cfg.SupportLocales {
        supported[l] = true
    }

    defaultLocale := cfg.DefaultLocale

    return func(ctx fiber.Ctx) error {
        locale := parseLocaleWith(ctx.Get("Accept-Language"), supported, defaultLocale)
        ctx.Locals(localeKey, locale)
        ctx.Locals(defaultLocaleKey, defaultLocale)
        return ctx.Next()
    }
}

func GetLocale(ctx fiber.Ctx) string {
    if locale, ok := ctx.Locals(localeKey).(string); ok {
        return locale
    }

    return getDefaultLocale(ctx)
}

func getDefaultLocale(ctx fiber.Ctx) string {
    if locale, ok := ctx.Locals(defaultLocaleKey).(string); ok {
        return locale
    }

    return "en" // fallback
}

func parseLocaleWith(header string, supportedLocales map[string]bool, defaultLocale string) string {
    if header == "" {
        return defaultLocale
    }

    locale, quality := defaultLocale, -1.0

    for _, part := range strings.Split(header, ",") {
        part = strings.TrimSpace(part)
        segments := strings.Split(part, ";")
        lang := strings.ToLower(strings.TrimSpace(segments[0]))
        lang = strings.SplitN(lang, "-", 2)[0]

        q := 1.0
        for _, seg := range segments[1:] {
            seg = strings.TrimSpace(seg)
            if strings.HasPrefix(seg, "q=") {
                if val, err := strconv.ParseFloat(seg[2:], 64); err == nil {
                    q = val
                }
            }
        }

        if supportedLocales[lang] && q > quality {
            locale, quality = lang, q
        }
    }

    return locale
}
