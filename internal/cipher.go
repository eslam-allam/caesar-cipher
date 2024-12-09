package internal

import (
	"strings"
	"unicode"
)

func Encrypt(data string, shift int) string {
	var result strings.Builder
	for _, c := range data {
		if !unicode.IsLetter(c) {
			result.WriteRune(c)
			continue
		}
		if unicode.IsUpper(c) {
			result.WriteByte(byte(int(c)+shift-65)%26 + 65)
			continue
		}
		result.WriteByte(byte(int(c)+shift-97)%26 + 97)
	}
	return result.String()
}

func Decrypt(data string, shift int) string {
	return Encrypt(data, 26-shift)
}
