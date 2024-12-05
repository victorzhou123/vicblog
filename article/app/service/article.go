package service

import (
	"github.com/victorzhou123/vicblog/article/app/dto"
	articledmsvc "github.com/victorzhou123/vicblog/article/domain/article/service"
	categorydmsvc "github.com/victorzhou123/vicblog/category-server/domain/category/service"
	cmappdto "github.com/victorzhou123/vicblog/common/app/dto"
	cmappevent "github.com/victorzhou123/vicblog/common/app/event"
	"github.com/victorzhou123/vicblog/common/domain/entity"
	"github.com/victorzhou123/vicblog/common/domain/mq"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	cmrepo "github.com/victorzhou123/vicblog/common/domain/repository"
	"github.com/victorzhou123/vicblog/common/infrastructure/mysql"
	"github.com/victorzhou123/vicblog/common/log"
	tagdmsvc "github.com/victorzhou123/vicblog/tag-server/domain/tag/service"
)

const (
	topicAddArticleReadTimes = cmappevent.TopicAddArticleReadTimes

	fieldArticleId = cmappevent.FieldArticleId
)

type ArticleAppService interface {
	GetArticleById(articleId cmprimitive.Id) (dto.ArticleWithTagCateDto, error)
	GetArticle(*dto.GetArticleCmd) (dto.ArticleDetailDto, error)
	GetArticleList(*dto.GetArticleListCmd) (dto.ArticleListDto, error)
	GetArticleCardListByCateId(*dto.GetArticleCardListByCateIdCmd) (dto.ArticleCardListDto, error)
	GetArticleCardListByTagId(*dto.GetArticleCardListByTagIdCmd) (dto.ArticleCardListDto, error)
	GetArticleListClassifiedByMonth(*cmappdto.PaginationCmd) (dto.ArticlesClassifiedByMonthDto, error)
	PaginationListArticle(*dto.ListAllArticlesCmd) (dto.ArticleDetailsListDto, error)
	SearchArticle(*dto.SearchArticlesCmd) (dto.ArticleCardsWithSummaryDto, error)

	AddArticle(*dto.AddArticleCmd) error

	DeleteArticle(user cmprimitive.Username, articleId cmprimitive.Id) error

	UpdateArticle(*dto.UpdateArticleCmd) error
}

type articleAppService struct {
	tx        cmrepo.Transaction
	article   articledmsvc.ArticleService
	cate      categorydmsvc.CategoryService
	tag       tagdmsvc.TagService
	publisher mq.MQ
}

func NewArticleAppService(
	tx mysql.Transaction,
	article articledmsvc.ArticleService,
	cate categorydmsvc.CategoryService,
	tag tagdmsvc.TagService,
	publisher mq.MQ,
) ArticleAppService {
	return &articleAppService{
		tx:        tx,
		article:   article,
		cate:      cate,
		tag:       tag,
		publisher: publisher,
	}
}

func (s *articleAppService) GetArticleById(articleId cmprimitive.Id) (dto.ArticleWithTagCateDto, error) {

	// get article (has parsed content to html)
	article, err := s.article.GetArticleByIdWithContentParsed(articleId)
	if err != nil {
		log.Errorf("get parsed article failed, err: %s", err.Error())

		return dto.ArticleWithTagCateDto{}, err
	}

	// get tags by article id
	tags, err := s.tag.GetArticleTag(articleId)
	if err != nil {
		log.Errorf("get article tag failed, err: %s", err.Error())

		return dto.ArticleWithTagCateDto{}, err
	}

	// get categories by article id
	cates, err := s.cate.GetArticleCategory(articleId)
	if err != nil {
		log.Errorf("get article category failed, err: %s", err.Error())

		return dto.ArticleWithTagCateDto{}, err
	}

	// get prev and next article
	preNextArticle, err := s.article.GetPrevAndNextArticle(articleId)
	if err != nil {
		log.Errorf("get prev and next failed, err: %s", err.Error())

		return dto.ArticleWithTagCateDto{}, err
	}

	// publish message
	msg, err := cmappevent.ToMessage(map[string]string{fieldArticleId: article.Id.Id()})
	if err != nil {
		log.Errorf("new message failed, err: %s", err.Error())
	}
	s.publisher.Publish(topicAddArticleReadTimes, &msg)

	return dto.ToArticleWithTagCateDto(article, tags, cates, preNextArticle.Prev, preNextArticle.Next), nil
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

func (s *articleAppService) GetArticleCardListByCateId(
	cmd *dto.GetArticleCardListByCateIdCmd,
) (dto.ArticleCardListDto, error) {

	if err := cmd.Validate(); err != nil {
		return dto.ArticleCardListDto{}, err
	}

	// get article id list through category id
	articleIds, err := s.cate.GetRelatedArticleIdsThroughCateId(cmd.CategoryId)
	if err != nil {
		return dto.ArticleCardListDto{}, err
	}

	// get article card list through batch article id
	articleCardDto, err := s.article.GetArticleCardList(&articledmsvc.ArticleCardsCmd{
		Pagination: *cmd.ToPagination(),
		ArticleIds: articleIds,
	})
	if err != nil {
		return dto.ArticleCardListDto{}, err
	}

	return dto.ToArticleCardListDto(articleCardDto.PaginationStatus, articleCardDto.ArticleCards), nil
}

func (s *articleAppService) GetArticleCardListByTagId(
	cmd *dto.GetArticleCardListByTagIdCmd) (dto.ArticleCardListDto, error) {

	if err := cmd.Validate(); err != nil {
		return dto.ArticleCardListDto{}, err
	}

	// get article id list through tag id
	articleIds, err := s.tag.GetRelatedArticleIdsThroughTagId(cmd.TagId)
	if err != nil {
		return dto.ArticleCardListDto{}, err
	}

	// get article card list through batch article id
	articleCardDto, err := s.article.GetArticleCardList(&articledmsvc.ArticleCardsCmd{
		Pagination: *cmd.ToPagination(),
		ArticleIds: articleIds,
	})
	if err != nil {
		return dto.ArticleCardListDto{}, err
	}

	return dto.ToArticleCardListDto(articleCardDto.PaginationStatus, articleCardDto.ArticleCards), nil
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

func (s *articleAppService) GetArticleListClassifiedByMonth(cmd *cmappdto.PaginationCmd) (dto.ArticlesClassifiedByMonthDto, error) {

	sub, err := s.article.ListArticlesClassifiedByMonth(cmd.ToPagination())
	if err != nil {
		return dto.ArticlesClassifiedByMonthDto{}, err
	}

	archives := make([]dto.ArticleCreatedInSameMonth, len(sub.ArticleArchives))
	for i := range sub.ArticleArchives {

		articles := make([]dto.ArticleCardDto, len(sub.ArticleArchives[i].ArticleCards))
		for j := range sub.ArticleArchives[i].ArticleCards {
			articles[j] = dto.ToArticleCardDto(sub.ArticleArchives[i].ArticleCards[j])
		}

		archives[i] = dto.ArticleCreatedInSameMonth{
			Date:     sub.ArticleArchives[i].Time.TimeYearMonthOnly(),
			Articles: articles,
		}
	}

	return dto.ArticlesClassifiedByMonthDto{
		PaginationDto:       cmappdto.ToPaginationDto(sub.PaginationStatus),
		ArticlesInSameMonth: archives,
	}, nil
}

func (s *articleAppService) SearchArticle(cmd *dto.SearchArticlesCmd) (dto.ArticleCardsWithSummaryDto, error) {

	sa, err := s.article.SearchArticles(cmd.Word, cmd.ToPagination())
	if err != nil {
		return dto.ArticleCardsWithSummaryDto{}, err
	}

	return dto.ToArticleCardsWithSummaryDto(sa), nil
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
