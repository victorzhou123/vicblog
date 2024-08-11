package service

import (
	"victorzhou123/vicblog/article/app/dto"
	articledmsvc "victorzhou123/vicblog/article/domain/article/service"
	categorydmsvc "victorzhou123/vicblog/article/domain/category/service"
	tagdmsvc "victorzhou123/vicblog/article/domain/tag/service"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	cmrepo "victorzhou123/vicblog/common/domain/repository"
	"victorzhou123/vicblog/common/infrastructure/mysql"
	"victorzhou123/vicblog/common/log"
)

type ArticleAppService interface {
	GetArticle(*dto.GetArticleCmd) (dto.ArticleDetailDto, error)

	AddArticle(*dto.AddArticleCmd) error

	DeleteArticle(user cmprimitive.Username, articleId cmprimitive.Id) error
}

type articleAppService struct {
	tx      cmrepo.Transaction
	article articledmsvc.ArticleService
	cate    categorydmsvc.CategoryService
	tag     tagdmsvc.TagService
}

func NewArticleAppService(
	tx mysql.Transaction,
	article articledmsvc.ArticleService,
	cate categorydmsvc.CategoryService,
	tag tagdmsvc.TagService,
) ArticleAppService {
	return &articleAppService{
		tx:      tx,
		article: article,
		cate:    cate,
		tag:     tag,
	}
}

func (s *articleAppService) GetArticle(cmd *dto.GetArticleCmd) (dto.ArticleDetailDto, error) {

	// get article
	article, err := s.article.GetArticle(&cmd.GetArticleCmd)
	if err != nil {

		log.Errorf("user %s get article %s failed, err: %s",
			cmd.User.Username(), cmd.ArticleId.Id(), err.Error())

		return dto.ArticleDetailDto{}, err
	}

	// get relation tags
	tagIds, err := s.tag.GetRelationWithArticle(article.Id)
	if err != nil {

		log.Errorf("get all tags of article %s failed, err: %s",
			article.Id.Id(), err.Error())

		return dto.ArticleDetailDto{}, err
	}

	// get relation category
	cateId, err := s.cate.GetRelationWithArticle(article.Id)
	if err != nil {

		log.Errorf("get category of article %s failed, err: %s",
			article.Id.Id(), err.Error())

		return dto.ArticleDetailDto{}, err
	}

	return dto.ToArticleDetailDto(article, tagIds, cateId), nil
}

func (s *articleAppService) AddArticle(cmd *dto.AddArticleCmd) error {

	// transaction begin
	if err := s.tx.Begin(); err != nil {

		log.Errorf("transaction begin error, err: %s", err.Error())

		return err
	}

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

		log.Errorf("user %s build tag relation with article failed, err: %s",
			cmd.Owner.Username(), err.Error())

		return err
	}

	// make relationship with category
	if err := s.cate.BuildRelationWithArticle(article, cmd.Category); err != nil {

		log.Errorf("user %s build category relation with article failed, err: %s",
			cmd.Owner.Username(), err.Error())

		return err
	}

	// transaction commit
	if err := s.tx.Commit(); err != nil {

		log.Errorf("transaction commit error, err: %s", err.Error())

		return err
	}

	log.Infof("user %s add article success", cmd.Owner.Username())

	return nil
}

func (s *articleAppService) DeleteArticle(user cmprimitive.Username, articleId cmprimitive.Id) error {

	// transaction begin
	if err := s.tx.Begin(); err != nil {

		log.Errorf("transaction begin error, err: %s", err.Error())

		return err
	}

	// delete article
	if err := s.article.Delete(user, articleId); err != nil {

		log.Errorf("user %s delete article (articleId: %s) failed, err: %s",
			user.Username(), articleId.Id(), err.Error())

		return err
	}

	// remove relation with tags
	if err := s.tag.RemoveRelationWithArticle(articleId); err != nil {

		log.Errorf("user %s remove all tags relations of article %s failed, err: %s",
			user.Username(), articleId.Id(), err.Error())

		return err
	}

	// remove relation with category
	if err := s.cate.RemoveRelationWithArticle(articleId); err != nil {

		log.Errorf("user %s remove all cates relations of article %s failed, err: %s",
			user.Username(), articleId.Id(), err.Error())

		return err
	}

	// transaction commit
	if err := s.tx.Commit(); err != nil {

		log.Errorf("transaction commit error, err: %s", err.Error())

		return err
	}

	log.Infof("user %s delete article %s success", user.Username(), articleId.Id())

	return nil
}
