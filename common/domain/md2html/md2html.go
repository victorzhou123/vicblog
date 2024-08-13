package md2html

import "victorzhou123/vicblog/common/domain/primitive"

type Md2Html interface {
	Render(primitive.Text) primitive.Text
}
