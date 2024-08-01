package validator

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
)

const (
	regexUsername = `^[a-zA-Z0-9_]{3,8}$`
	regexPassword = `^.{8,14}$` // #nosec G101
	regexEmail    = `^\w+(-+.\w+)*@\w+(-.\w+)*.\w+(-.\w+)*$`
	regexTitle    = `^.{8,255}$` // #nosec G101
)

var (
	// username: letters, digitals or underline only and 3 to 8 characters allowed
	regexCompUsername = regexp.MustCompile(regexUsername)

	// password: 8 to 16 letters or numbers allowed
	regexCompPassword = regexp.MustCompile(regexPassword)

	regexCompEmail = regexp.MustCompile(regexEmail)

	regexCompTitle = regexp.MustCompile(regexEmail)
)

type validateCmd struct {
	s     string
	regex *regexp.Regexp
	issue string
}

func IsUsername(v string) error {
	return validate(&validateCmd{v, regexCompUsername, "username"})
}

func IsPassword(v string) error {
	return validate(&validateCmd{v, regexCompPassword, "password"})
}

func IsEmail(v string) error {
	return validate(&validateCmd{v, regexCompEmail, "email"})
}

func IsTitle(v string) error {
	return validate(&validateCmd{v, regexCompTitle, "title"})
}

func validate(cmd *validateCmd) error {
	if cmd.s == "" {
		return errors.New("empty input cannot be validate")
	}

	if !cmd.regex.MatchString(cmd.s) {
		return fmt.Errorf("validate %s failed", cmd.issue)
	}

	return nil
}

func IsUrl(v string) error {
	if _, err := url.Parse(v); err != nil {
		return fmt.Errorf("input is not an url")
	}

	return nil
}
