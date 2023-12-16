package text

import (
	"strings"
	"unicode"
)

func ToCamelCase(input string) string {
	isToUpper := false
	return strings.Map(func(r rune) rune {
		if r == '_' || r == '-' || r == ' ' {
			isToUpper = true
			return -1 // Skip this character
		}
		if isToUpper {
			isToUpper = false
			return unicode.ToUpper(r)
		}
		return unicode.ToLower(r)
	}, input)
}

func CamelToDashDelimited(input string) string {
	var sb strings.Builder
	for i, r := range input {
		if unicode.IsUpper(r) && i > 0 {
			sb.WriteRune('-')
		}
		sb.WriteRune(unicode.ToLower(r))
	}
	return sb.String()
}
