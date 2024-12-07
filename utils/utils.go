package utils

import (
	"github.com/go-passwd/validator"
	"net/mail"
	"regexp"
)

var alphanumeric = regexp.MustCompile("^[a-zA-Z0-9]*$")

func IsAlphanumeric(str string) bool {
	return alphanumeric.MatchString(str)
}

func IsEmail(str string) bool {
	_, err := mail.ParseAddress(str)
	return err == nil
}

func IsValidPassword(str string) bool {
	passwordValidator := validator.New(
		validator.MinLength(8, nil),
		validator.ContainsAtLeast("ABCDEFGHIJKLMNOPQRSTUVWXYZ", 1, nil),
		validator.ContainsAtLeast("abcdefghijklmnopqrstuvwxyz", 1, nil),
		validator.ContainsAtLeast("0123456789", 1, nil),
		validator.ContainsAtLeast("#$@!%&*?", 1, nil),
	)
	err := passwordValidator.Validate(str)
	return err == nil
}
