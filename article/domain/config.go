package domain

import (
	"github.com/victorzhou123/vicblog/article/domain/category"
	"github.com/victorzhou123/vicblog/article/domain/tag"
)

type Config struct {
	Category category.Config `json:"category"`
	Tag      tag.Config      `json:"tag"`
}
