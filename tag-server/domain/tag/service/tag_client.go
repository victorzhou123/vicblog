package service

import (
	"context"
	"time"

	"github.com/victorzhou123/vicblog/common/controller/rpc"
	cmdto "github.com/victorzhou123/vicblog/common/domain/dto"
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/tag-server/domain/tag/entity"
	"github.com/victorzhou123/vicblog/tag-server/domain/tag/repository"
)

type tagClient struct {
	expire time.Duration
	rpcSvc rpc.TagServiceClient
}

func NewTagServer(expire time.Duration, rpcSvc rpc.TagServiceClient) TagService {
	return &tagClient{
		expire: expire,
		rpcSvc: rpcSvc,
	}
}

func (s *tagClient) AddTags(names repository.TagNames) error {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	ns := make([]*rpc.TagName, len(names.Names))
	for i := range ns {
		ns[i] = &rpc.TagName{Name: names.Names[i].TagName()}
	}

	_, err := s.rpcSvc.AddTags(ctx, &rpc.TagNames{
		Names: ns,
	})

	return err
}

func (s *tagClient) GetArticleTag(articleId cmprimitive.Id) ([]entity.Tag, error) {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	tags, err := s.rpcSvc.GetArticleTag(ctx, &rpc.Id{Id: articleId.Id()})
	if err != nil {
		return nil, err
	}

	ts := make([]entity.Tag, len(tags.GetTags()))
	for i := range ts {
		ts[i], err = toTag(tags.Tags[i])
		if err != nil {
			return nil, err
		}
	}

	return ts, nil
}

func (s *tagClient) ListTagByPagination(pagination *cment.Pagination) (TagListDto, error) {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	tagList, err := s.rpcSvc.ListTagByPagination(ctx, cmdto.ToProtoPagination(*pagination))
	if err != nil {
		return TagListDto{}, err
	}

	return toTagListDto(tagList)
}

func (s *tagClient) ListTags(amount cmprimitive.Amount) ([]entity.TagWithRelatedArticleAmount, error) {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	ta, err := s.rpcSvc.ListTags(ctx, &rpc.Amount{Amount: int64(amount.Amount())})
	if err != nil {
		return nil, err
	}

	tas := make([]entity.TagWithRelatedArticleAmount, len(ta.GetTags()))
	for i := range tas {
		if tas[i], err = toTagWithRelatedArticleAmount(ta.Tags[i]); err != nil {
			return nil, err
		}
	}

	return tas, nil
}

func (s *tagClient) GetTotalNumberOfTags() (cmprimitive.Amount, error) {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	amount, err := s.rpcSvc.GetTotalNumberOfTags(ctx, nil)
	if err != nil {
		return nil, err
	}

	return cmprimitive.NewAmount(int(amount.GetAmount()))
}

func (s *tagClient) Delete(id cmprimitive.Id) error {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	_, err := s.rpcSvc.Delete(ctx, &rpc.Id{Id: id.Id()})

	return err
}

func (s *tagClient) GetRelationWithArticle(articleId cmprimitive.Id) ([]cmprimitive.Id, error) {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	ids, err := s.rpcSvc.GetRelationWithArticle(ctx, &rpc.Id{Id: articleId.Id()})
	if err != nil {
		return nil, err
	}

	identities := make([]cmprimitive.Id, len(ids.GetIds()))
	for i := range identities {
		identities[i] = cmprimitive.NewId(ids.Ids[i].GetId())
	}

	return identities, nil
}

func (s *tagClient) GetRelatedArticleIdsThroughTagId(tagId cmprimitive.Id) ([]cmprimitive.Id, error) {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	ids, err := s.rpcSvc.GetRelatedArticleIdsThroughTagId(ctx, &rpc.Id{Id: tagId.Id()})
	if err != nil {
		return nil, err
	}

	identities := make([]cmprimitive.Id, len(ids.GetIds()))
	for i := range identities {
		identities[i] = cmprimitive.NewId(ids.Ids[i].GetId())
	}

	return identities, nil
}

func (s *tagClient) BuildRelationWithArticle(articleId cmprimitive.Id, tagIds []cmprimitive.Id) error {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	ts := make([]*rpc.Id, len(tagIds))
	for i := range ts {
		ts[i] = &rpc.Id{Id: tagIds[i].Id()}
	}

	_, err := s.rpcSvc.BuildRelationWithArticle(ctx, &rpc.ArticleIdAndTagIds{
		ArticleId: &rpc.Id{Id: articleId.Id()},
		TagIds:    ts,
	})

	return err
}

func (s *tagClient) RemoveRelationWithArticle(articleId cmprimitive.Id) error {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	_, err := s.rpcSvc.RemoveRelationWithArticle(ctx, &rpc.Id{Id: articleId.Id()})

	return err
}
