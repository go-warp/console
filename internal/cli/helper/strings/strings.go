package strings

import (
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ToPascalCase converts a string to PascalCase
func ToPascalCase(s string) string {
	caser := cases.Title(language.Und)

	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})

	for i, word := range words {
		words[i] = caser.String(strings.ToLower(word))
	}

	return strings.Join(words, "")
}

// ToCamelCase converts a string to camelCase
func ToCamelCase(s string) string {
	caser := cases.Title(language.Und)

	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})

	for i, word := range words {
		if i == 0 {
			words[i] = strings.ToLower(word)
		} else {
			words[i] = caser.String(strings.ToLower(word))
		}
	}

	return strings.Join(words, "")
}
