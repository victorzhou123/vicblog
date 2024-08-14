package md2html

import "github.com/victorzhou123/vicblog/common/domain/primitive"

type Md2Html interface {
	Render(primitive.Text) primitive.Text
}
