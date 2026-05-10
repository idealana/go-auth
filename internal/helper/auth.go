package helper

import (
	"strings"
)

func ParseBearerToken(header string) string {
    parts := strings.SplitN(strings.TrimSpace(header), " ", 2)
    if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
        return ""
    }
    return strings.TrimSpace(parts[1])
}
