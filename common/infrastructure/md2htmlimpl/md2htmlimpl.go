package md2htmlimpl

import (
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"

	"github.com/victorzhou123/vicblog/common/domain/md2html"
	"github.com/victorzhou123/vicblog/common/domain/primitive"
)

type md2Html struct{}

func NewMd2Html() md2html.Md2Html {
	return &md2Html{}
}

func (m *md2Html) Render(content primitive.Text) primitive.Text {

	b := mdToHTML(content.Byte())

	return primitive.NewOutPutArticleContent(string(b))
}

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
