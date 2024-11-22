package article

import "github.com/victorzhou123/vicblog/article/domain"

type Config struct {
	Domain domain.Config `json:"domain"`
}
