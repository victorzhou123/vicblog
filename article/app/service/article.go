package service

import (
	"victorzhou123/vicblog/article/app/dto"
	articledmsvc "victorzhou123/vicblog/article/domain/article/service"
	categorydmsvc "victorzhou123/vicblog/article/domain/category/service"
	tagdmsvc "victorzhou123/vicblog/article/domain/tag/service"
	cmappdto "victorzhou123/vicblog/common/app/dto"
	"victorzhou123/vicblog/common/domain/entity"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	cmrepo "victorzhou123/vicblog/common/domain/repository"
	"victorzhou123/vicblog/common/infrastructure/mysql"
	"victorzhou123/vicblog/common/log"
)

type ArticleAppService interface {
	GetArticleById(articleId cmprimitive.Id) (dto.ArticleWithTagCateDto, error)
	GetArticle(*dto.GetArticleCmd) (dto.ArticleDetailDto, error)
	GetArticleList(*dto.GetArticleListCmd) (dto.ArticleListDto, error)
	PaginationListArticle(*dto.ListAllArticlesCmd) (dto.ArticleDetailsListDto, error)

	AddArticle(*dto.AddArticleCmd) error

	DeleteArticle(user cmprimitive.Username, articleId cmprimitive.Id) error

	UpdateArticle(*dto.UpdateArticleCmd) error
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

func (s *articleAppService) GetArticleById(articleId cmprimitive.Id) (dto.ArticleWithTagCateDto, error) {

	// get article (has parsed content to html)
	article, err := s.article.GetArticleByIdWithContentParsed(articleId)
	if err != nil {
		return dto.ArticleWithTagCateDto{}, err
	}

	// get tags by article id
	tags, err := s.tag.GetArticleTag(articleId)
	if err != nil {
		return dto.ArticleWithTagCateDto{}, err
	}

	// get categories by article id
	cates, err := s.cate.GetArticleCategory(articleId)
	if err != nil {
		return dto.ArticleWithTagCateDto{}, err
	}

	return dto.ToArticleWithTagCateDto(article, tags, cates), nil
}

func (s *articleAppService) GetArticle(cmd *dto.GetArticleCmd) (dto.ArticleDetailDto, error) {

	// get article
	article, err := s.article.GetArticle(&cmd.GetArticleCmd)
	if err != nil {

		log.Errorf("user %s get article %s failed, err: %s",
			cmd.User.Username(), cmd.ArticleId.Id(), err.Error())

		return dto.ArticleDetailDto{}, err
	}

	// get tags and category of article
	tagIds, cateId, err := s.getArticleTagsAndCategoryById(article.Id)
	if err != nil {
		return dto.ArticleDetailDto{}, err
	}

	return dto.ToArticleDetailDto(article, tagIds, cateId), nil
}

func (s *articleAppService) GetArticleList(cmd *dto.GetArticleListCmd) (dto.ArticleListDto, error) {

	articleSummaryDto, err := s.article.GetArticleList(&articledmsvc.ArticleListCmd{
		Pagination: *cmd.ToPagination(),
		User:       cmd.User,
	})
	if err != nil {
		return dto.ArticleListDto{}, err
	}

	return dto.ToArticleListDto(articleSummaryDto.PaginationStatus, articleSummaryDto.Articles), nil
}

func (s *articleAppService) PaginationListArticle(cmd *dto.ListAllArticlesCmd) (dto.ArticleDetailsListDto, error) {

	// list articles
	articleSummaryDto, err := s.article.PaginationListArticle(&entity.Pagination{
		CurPage: cmd.CurPage, PageSize: cmd.PageSize,
	})
	if err != nil {
		return dto.ArticleDetailsListDto{}, err
	}

	articleDetailListDtos := make([]dto.ArticleDetailListDto, len(articleSummaryDto.Articles))
	for i := range articleSummaryDto.Articles {

		tags, err := s.tag.GetArticleTag(articleSummaryDto.Articles[i].Id)
		if err != nil {
			return dto.ArticleDetailsListDto{}, err
		}

		category, err := s.cate.GetArticleCategory(articleSummaryDto.Articles[i].Id)
		if err != nil {
			return dto.ArticleDetailsListDto{}, err
		}

		articleDetailListDtos[i] = dto.ToArticleDetailListDto(articleSummaryDto.Articles[i], category, tags)
	}

	return dto.ArticleDetailsListDto{
		PaginationDto: cmappdto.ToPaginationDto(articleSummaryDto.PaginationStatus),
		Articles:      articleDetailListDtos,
	}, nil
}

func (s *articleAppService) AddArticle(cmd *dto.AddArticleCmd) error {

	// transaction begin
	if err := s.tx.Begin(); err != nil {

		log.Errorf("transaction begin error, err: %s", err.Error())

		return err
	}

	// new article
	articleId, err := s.article.AddArticle(cmd.ToArticleInfo())
	if err != nil {
		return err
	}

	// make relationship with tag
	if err := s.tag.BuildRelationWithArticle(articleId, cmd.Tags); err != nil {

		log.Errorf("user %s build tag relation with article failed, err: %s",
			cmd.Owner.Username(), err.Error())

		return err
	}

	// make relationship with category
	if err := s.cate.BuildRelationWithArticle(articleId, cmd.Category); err != nil {

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

func (s *articleAppService) UpdateArticle(cmd *dto.UpdateArticleCmd) error {

	// transaction begin
	if err := s.tx.Begin(); err != nil {

		log.Errorf("transaction begin error, err: %s", err.Error())

		return err
	}

	// update article
	if err := s.article.UpdateArticle(cmd.Id, cmd.ToArticleInfo()); err != nil {

		log.Errorf("user %s update article %s failed, err: %s",
			cmd.AddArticleCmd.Owner.Username(), cmd.Id.Id(), err.Error())

		return err
	}

	// remove relation with tags
	if err := s.tag.RemoveRelationWithArticle(cmd.Id); err != nil {

		log.Errorf("user %s remove all tags relations of article %s failed, err: %s",
			cmd.AddArticleCmd.Owner.Username(), cmd.Id.Id(), err.Error())

		return err
	}

	// make relationship with tags
	if err := s.tag.BuildRelationWithArticle(cmd.Id, cmd.AddArticleCmd.Tags); err != nil {

		log.Errorf("user %s build tag relation with article failed, err: %s",
			cmd.AddArticleCmd.Owner.Username(), err.Error())

		return err
	}

	// remove relation with category
	if err := s.cate.RemoveRelationWithArticle(cmd.Id); err != nil {

		log.Errorf("user %s remove all cates relations of article %s failed, err: %s",
			cmd.AddArticleCmd.Owner.Username(), cmd.Id.Id(), err.Error())

		return err
	}

	// make relationship with category
	if err := s.cate.BuildRelationWithArticle(cmd.Id, cmd.AddArticleCmd.Category); err != nil {

		log.Errorf("user %s build category relation with article failed, err: %s",
			cmd.AddArticleCmd.Owner.Username(), err.Error())

		return err
	}

	// transaction commit
	if err := s.tx.Commit(); err != nil {

		log.Errorf("transaction commit error, err: %s", err.Error())

		return err
	}

	log.Infof("user %s update article %s success",
		cmd.AddArticleCmd.Owner.Username(), cmd.Id.Id())

	return nil
}

func (s *articleAppService) getArticleTagsAndCategoryById(articleId cmprimitive.Id) ([]cmprimitive.Id, cmprimitive.Id, error) {
	// get relation tags
	tagIds, err := s.tag.GetRelationWithArticle(articleId)
	if err != nil {

		log.Errorf("get all tags of article %s failed, err: %s",
			articleId.Id(), err.Error())

		return nil, nil, err
	}

	// get relation category
	cateId, err := s.cate.GetRelationWithArticle(articleId)
	if err != nil {

		log.Errorf("get category of article %s failed, err: %s",
			articleId.Id(), err.Error())

		return nil, nil, err
	}

	return tagIds, cateId, nil
}
