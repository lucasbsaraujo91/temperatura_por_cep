package utils

import (
	"io"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func SanitizeCity(city string) string {
	return strings.ReplaceAll(city, " ", "+")
}

// removeAccents remove acentos de uma string
func RemoveAccents(str string) string {
	t := transform.NewReader(strings.NewReader(str), norm.NFD)
	normalized, _ := io.ReadAll(t)
	result := string(normalized)

	var sb strings.Builder
	for _, r := range result {
		if unicode.Is(unicode.Mn, r) {
			continue // Ignora caracteres de acento
		}
		sb.WriteRune(r)
	}
	return sb.String()
}
