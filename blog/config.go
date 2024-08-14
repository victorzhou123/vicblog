package blog

import "github.com/victorzhou123/vicblog/blog/infrastructure/repositoryimpl"

type Config struct {
	BlogInfo repositoryimpl.Config `json:"blog_info"`
}
