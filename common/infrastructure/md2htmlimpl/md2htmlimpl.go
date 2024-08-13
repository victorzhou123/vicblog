package md2htmlimpl

import (
	"github.com/russross/blackfriday/v2"

	"victorzhou123/vicblog/common/domain/md2html"
	"victorzhou123/vicblog/common/domain/primitive"
)

type md2Html struct{}

func NewMd2Html() md2html.Md2Html {
	return &md2Html{}
}

func (m *md2Html) Render(content primitive.Text) primitive.Text {

	b := blackfriday.Run(content.Byte(), blackfriday.WithNoExtensions())

	return primitive.NewOutPutArticleContent(string(b))
}
