package domain

import (
	"github.com/victorzhou123/vicblog/article/domain/category"
)

type Config struct {
	Category category.Config `json:"category"`
}
