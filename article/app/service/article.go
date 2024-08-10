package service

import (
	"victorzhou123/vicblog/article/app/dto"
	articledmsvc "victorzhou123/vicblog/article/domain/article/service"
	categorydmsvc "victorzhou123/vicblog/article/domain/category/service"
	tagdmsvc "victorzhou123/vicblog/article/domain/tag/service"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
)

type ArticleAppService interface {
	AddArticle(*dto.AddArticleCmd) error
	DeleteArticle(user cmprimitive.Username, articleId cmprimitive.Id) error
}

type articleAppService struct {
	article articledmsvc.ArticleService
	cate    categorydmsvc.CategoryService
	tag     tagdmsvc.TagService
}

func NewArticleAggService(
	article articledmsvc.ArticleService,
	cate categorydmsvc.CategoryService,
	tag tagdmsvc.TagService,
) ArticleAppService {
	return &articleAppService{
		article: article,
		cate:    cate,
		tag:     tag,
	}
}

func (s *articleAppService) AddArticle(cmd *dto.AddArticleCmd) error {

	// new article
	articleId, err := s.article.AddArticle(&articledmsvc.ArticleCmd{
		Owner:   cmd.Owner,
		Title:   cmd.Title,
		Summary: cmd.Summary,
		Content: cmd.Content,
		Cover:   cmd.Cover,
	})
	if err != nil {
		return err
	}

	article := cmprimitive.NewIdByUint(articleId)

	// make relationship with tag
	if err := s.tag.BuildRelationWithArticle(article, cmd.Tags); err != nil {
		return err
	}

	// make relationship with category
	if err := s.cate.BuildRelationWithArticle(article, cmd.Category); err != nil {
		return err
	}

	return nil
}

func (s *articleAppService) DeleteArticle(user cmprimitive.Username, articleId cmprimitive.Id) error {

	// delete article
	if err := s.article.Delete(user, articleId); err != nil {
		return err
	}

	// remove relation with tags
	if err := s.tag.RemoveRelationWithArticle(articleId); err != nil {
		return err
	}

	// remove relation with category
	if err := s.cate.RemoveRelationWithArticle(articleId); err != nil {
		return err
	}

	return nil
}
