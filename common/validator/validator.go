package validator

import (
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	regexUsername = `^[a-zA-Z0-9_]{3,8}$`
	regexPassword = `^.{8,14}$` // #nosec G101
	regexEmail    = `^\w+(-+.\w+)*@\w+(-.\w+)*.\w+(-.\w+)*$`
	regexTitle    = `^.{3,255}$` // #nosec G101

	articleContentLengthLimit = 40000
	articleSummaryLengthLimit = 140
	categoryNameLengthLimit   = 60
	tagNameLengthLimit        = 60
	pictureNameLengthLimit    = 200
)

var (
	// username: letters, digitals or underline only and 3 to 8 characters allowed
	regexCompUsername = regexp.MustCompile(regexUsername)

	// password: 8 to 16 letters or numbers allowed
	regexCompPassword = regexp.MustCompile(regexPassword)

	regexCompEmail = regexp.MustCompile(regexEmail)

	regexCompTitle = regexp.MustCompile(regexTitle)

	allowedPictureExts = []string{".jpg", ".jpeg", ".png"}
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

func IsArticleContent(v string) error {
	if len(v) > articleContentLengthLimit {
		return fmt.Errorf("article content must less than %d", articleContentLengthLimit)
	}

	return nil
}

func IsArticleSummary(v string) error {
	if len(v) > articleSummaryLengthLimit {
		return fmt.Errorf("article summary must less than %d", articleSummaryLengthLimit)
	}

	return nil
}

func IsCategoryName(v string) error {
	if len(v) > categoryNameLengthLimit || len(v) <= 0 {
		return fmt.Errorf("category name length must greater than 0 and less than %d", categoryNameLengthLimit)
	}

	return nil
}

func IsTagName(v string) error {
	if len(v) > tagNameLengthLimit || len(v) <= 0 {
		return fmt.Errorf("tag name length must greater than 0 and less than %d", tagNameLengthLimit)
	}

	return nil
}

func IsAmount(v int) error {
	if v < 0 {
		return errors.New("amount must greater than or equal to 0")
	}

	return nil
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

func IsPictureName(name string) error {
	if len(name) > pictureNameLengthLimit || len(name) <= 0 {
		return fmt.Errorf("picture name length must bigger than 0 and less than %d", pictureNameLengthLimit)
	}

	// get extension of file
	ext := filepath.Ext(name)
	ext = strings.ToLower(ext)

	var allowedExt bool
	for i := range allowedPictureExts {
		if ext == allowedPictureExts[i] {
			allowedExt = true

			break
		}
	}

	if !allowedExt {
		return fmt.Errorf("picture extension must be in %v", allowedPictureExts)
	}

	return nil
}
