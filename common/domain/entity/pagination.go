package entity

import (
	"victorzhou123/vicblog/common/domain/primitive"
)

type Pagination struct {
	CurPage  primitive.CurPage
	PageSize primitive.PageSize
}
