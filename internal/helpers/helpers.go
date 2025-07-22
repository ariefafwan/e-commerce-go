package helpers

import (
	"log"
	"strings"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashed)
}

func GenerateSlug(input string) string {
	input = strings.ToLower(input)

	var slug strings.Builder
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			slug.WriteRune(r)
		} else if r == ' ' || r == '-' || r == '_' {
			slug.WriteRune('-')
		}
	}

	return strings.Join(strings.FieldsFunc(slug.String(), func(r rune) bool {
		return r == '-'
	}), "-")
}