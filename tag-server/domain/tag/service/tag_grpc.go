package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/victorzhou123/vicblog/common/controller/rpc"
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/tag-server/domain/tag/repository"
)

type tagServer struct {
	rpc.UnimplementedTagServiceServer

	tag TagService
}

func NewTagRpcServer(tag TagService) rpc.TagServiceServer {
	return &tagServer{
		tag: tag,
	}
}

func (s *tagServer) AddTags(ctx context.Context, names *rpc.TagNames) (*emptypb.Empty, error) {

	tagNames, err := toProtoTagNames(names)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, s.tag.AddTags(repository.TagNames{Names: tagNames})
}

func (s *tagServer) ListTagByPagination(ctx context.Context, pagination *rpc.Pagination) (*rpc.TagList, error) {

	curPage, err := cmprimitive.NewCurPageWithString(pagination.GetCurPage())
	if err != nil {
		return &rpc.TagList{}, err
	}

	pageSize, err := cmprimitive.NewPageSizeWithString(pagination.GetPageSize())
	if err != nil {
		return &rpc.TagList{}, err
	}

	dto, err := s.tag.ListTagByPagination(&cment.Pagination{
		CurPage:  curPage,
		PageSize: pageSize,
	})

	return dto.toProto(), err
}

func (s *tagServer) ListTags(ctx context.Context, amount *rpc.Amount) (*rpc.TagsWithRelatedArticleAmount, error) {

	am, err := cmprimitive.NewAmount(int(amount.GetAmount()))
	if err != nil {
		return nil, err
	}

	cas, err := s.tag.ListTags(am)
	if err != nil {
		return nil, err
	}

	tags := make([]*rpc.TagWithRelatedArticleAmount, len(cas))
	for i := range tags {
		tags[i] = toProtoTagWithRelatedArticleAmount(cas[i])
	}

	return &rpc.TagsWithRelatedArticleAmount{
		Tags: tags,
	}, nil
}

func (s *tagServer) GetArticleTag(ctx context.Context, id *rpc.Id) (*rpc.Tags, error) {

	tags, err := s.tag.GetArticleTag(cmprimitive.NewId(id.GetId()))
	if err != nil {
		return nil, err
	}

	ts := make([]*rpc.Tag, len(tags))
	for i := range ts {
		ts[i] = toProtoTag(tags[i])
	}

	return &rpc.Tags{
		Tags: ts,
	}, nil
}

func (s *tagServer) GetTotalNumberOfTags(ctx context.Context, empty *emptypb.Empty) (*rpc.Amount, error) {

	amount, err := s.tag.GetTotalNumberOfTags()
	if err != nil {
		return nil, err
	}

	return &rpc.Amount{Amount: int64(amount.Amount())}, nil
}

func (s *tagServer) Delete(ctx context.Context, id *rpc.Id) (*emptypb.Empty, error) {
	return nil, s.tag.Delete(cmprimitive.NewId(id.GetId()))
}

func (s *tagServer) GetRelationWithArticle(ctx context.Context, id *rpc.Id) (*rpc.Ids, error) {

	identities, err := s.tag.GetRelationWithArticle(cmprimitive.NewId(id.GetId()))
	if err != nil {
		return nil, err
	}

	ids := make([]*rpc.Id, len(identities))
	for i := range ids {
		ids[i] = &rpc.Id{Id: identities[i].Id()}
	}

	return &rpc.Ids{Ids: ids}, nil
}

func (s *tagServer) GetRelatedArticleIdsThroughCateId(ctx context.Context, id *rpc.Id) (*rpc.RespGetRelatedArticleIdsThroughCateId, error) {

	articleIds, err := s.tag.GetRelatedArticleIdsThroughTagId(cmprimitive.NewId(id.GetId()))
	if err != nil {
		return nil, err
	}

	ids := make([]*rpc.Id, len(articleIds))
	for i := range ids {
		ids[i] = &rpc.Id{
			Id: articleIds[i].Id(),
		}
	}

	return &rpc.RespGetRelatedArticleIdsThroughCateId{
		Ids: ids,
	}, nil
}

func (s *tagServer) BuildRelationWithArticle(ctx context.Context, req *rpc.ArticleIdAndTagIds) (*emptypb.Empty, error) {

	ids := make([]cmprimitive.Id, len(req.GetTagIds()))
	for i := range ids {
		ids[i] = cmprimitive.NewId(req.TagIds[i].GetId())
	}

	return nil, s.tag.BuildRelationWithArticle(
		cmprimitive.NewId(req.GetArticleId().GetId()),
		ids,
	)
}

func (s *tagServer) RemoveRelationWithArticle(ctx context.Context, id *rpc.Id) (*emptypb.Empty, error) {
	return nil, s.tag.RemoveRelationWithArticle(cmprimitive.NewId(id.GetId()))
}
