package blog

import "victorzhou123/vicblog/blog/infrastructure/repositoryimpl"

type Config struct {
	BlogInfo repositoryimpl.Config `json:"blog_info"`
}
