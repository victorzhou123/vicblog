package entity

import cmprimitive "victorzhou123/vicblog/common/domain/primitive"

type Tag struct {
	Id        cmprimitive.Id
	Name      TagName
	CreatedAt cmprimitive.Timex
}
