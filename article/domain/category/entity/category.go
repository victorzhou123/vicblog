package entity

import cmprimitive "victorzhou123/vicblog/common/domain/primitive"

type Category struct {
	Id   cmprimitive.Id
	Name CategoryName
}
