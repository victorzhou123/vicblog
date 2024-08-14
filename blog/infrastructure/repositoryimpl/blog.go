package repositoryimpl

import (
	"github.com/victorzhou123/vicblog/blog/domain/entity"
	"github.com/victorzhou123/vicblog/blog/domain/repository"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

type blogInfoImpl struct {
	Config
}

func NewBlogInfoImpl(cfg *Config) repository.Blog {
	return &blogInfoImpl{
		Config: *cfg,
	}
}

func (impl *blogInfoImpl) GetBlogInfo() (entity.Blog, error) {
	return impl.toBlog()
}

func (impl *blogInfoImpl) toBlog() (blog entity.Blog, err error) {

	if blog.Logo, err = cmprimitive.NewUrlx(impl.Config.Logo); err != nil {
		return
	}

	if blog.Name, err = cmprimitive.NewTitle(impl.Name); err != nil {
		return
	}

	if blog.Author, err = cmprimitive.NewUsername(impl.Author); err != nil {
		return
	}

	if blog.Introduction, err = cmprimitive.NewArticleContent(impl.Introduction); err != nil {
		return
	}

	if blog.Avatar, err = cmprimitive.NewUrlx(impl.Avatar); err != nil {
		return
	}

	if blog.GithubHomepage, err = cmprimitive.NewUrlx(impl.GithubHomepage); err != nil {
		return
	}

	if blog.GiteeHomepage, err = cmprimitive.NewUrlx(impl.GiteeHomepage); err != nil {
		return
	}

	if blog.CsdnHomepage, err = cmprimitive.NewUrlx(impl.CsdnHomepage); err != nil {
		return
	}

	if blog.ZhihuHomepage, err = cmprimitive.NewUrlx(impl.ZhihuHomepage); err != nil {
		return
	}

	return
}
