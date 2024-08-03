package controller

import "victorzhou123/vicblog/article/domain/category/entity"

type reqCategory struct {
	Name string `json:"name"`
}

func (req *reqCategory) toCategoryName() (entity.CategoryName, error) {
	return entity.NewCategoryName(req.Name)
}
