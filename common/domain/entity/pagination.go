package entity

import (
	"victorzhou123/vicblog/common/domain/primitive"
)

type Pagination struct {
	CurPage  primitive.CurPage
	PageSize primitive.PageSize
}

func (p *Pagination) ToPaginationStatus(total int) PaginationStatus {
	return PaginationStatus{
		Pagination: *p,
		Total:      total,
		PageCount:  total/p.PageSize.PageSize() + 1,
	}
}

type PaginationStatus struct {
	Pagination

	Total     int
	PageCount int
}
