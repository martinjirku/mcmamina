package text

import (
	"strings"
	"unicode"
)

func ToCamelCase(input string) string {
	splitFunc := func(c rune) bool {
		return c == '_' || c == '-' || c == ' '
	}

	parts := strings.FieldsFunc(input, splitFunc)
	for i, part := range parts {
		if i != 0 || unicode.IsUpper(rune(part[0])) {
			parts[i] = strings.ToUpper(part[:1]) + part[1:]
		}
	}
	return strings.Join(parts, "")

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
