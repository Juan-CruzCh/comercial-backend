package utils

import (
	"strings"
)

func GenerarCodigo(nombre string) string {
	consonantes := ""
	nombre = strings.ToLower(nombre)
	for _, c := range nombre {
		if isConsonante(c) {
			consonantes += string(c)
		}
	}
	return strings.ToUpper(consonantes)
}

func isConsonante(c rune) bool {
	vocales := "aeiou"
	return ('a' <= c && c <= 'z') && !strings.ContainsRune(vocales, c)
}
