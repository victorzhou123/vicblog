package primitive

import (
	"github.com/victorzhou123/vicblog/common/util"
	"github.com/victorzhou123/vicblog/common/validator"
)

type Text interface {
	Text() string
	Byte() []byte
}

type text string

func NewKeyWord(v string) (Text, error) {
	v = util.XssEscape(v)

	if err := validator.IsKeyWord(v); err != nil {
		return nil, err
	}

	return text(v), nil
}

func NewTitle(v string) (Text, error) {
	v = util.XssEscape(v)

	if err := validator.IsTitle(v); err != nil {
		return nil, err
	}

	return text(v), nil
}

func NewArticleContent(v string) (Text, error) {
	// v = util.XssEscape(v)

	if err := validator.IsArticleContent(v); err != nil {
		return nil, err
	}

	return text(v), nil
}

func NewCommentContent(v string) (Text, error) {
	v = util.XssEscape(v)

	if err := validator.IsCommentContent(v); err != nil {
		return nil, err
	}

	return text(v), nil
}

// be careful!: this function can only be used in output article content build
func NewOutPutArticleContent(v string) Text {
	return text(v)
}

func (t text) Text() string {
	return string(t)
}

func (t text) Byte() []byte {
	return []byte(t)
}
