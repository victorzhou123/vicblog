package service

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/victorzhou123/vicblog/category-server/domain/category/entity"
	"github.com/victorzhou123/vicblog/common/controller/rpc"
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

type categoryServer struct {
	rpc.UnimplementedCategoryServiceServer

	category CategoryService
}

func NewCategoryRpcServer(cate CategoryService) rpc.CategoryServiceServer {
	return &categoryServer{
		category: cate,
	}
}

func (s *categoryServer) AddCategory(ctx context.Context, name *rpc.CategoryName) (*emptypb.Empty, error) {

	cateName, err := entity.NewCategoryName(name.GetName())
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, s.category.AddCategory(cateName)
}

func (s *categoryServer) ListCategoryByPagination(ctx context.Context, pagination *rpc.Pagination) (*rpc.CategoryList, error) {

	curPage, err := cmprimitive.NewCurPageWithString(pagination.GetCurPage())
	if err != nil {
		return &rpc.CategoryList{}, err
	}

	pageSize, err := cmprimitive.NewPageSizeWithString(pagination.GetPageSize())
	if err != nil {
		return &rpc.CategoryList{}, err
	}

	dto, err := s.category.ListCategoryByPagination(&cment.Pagination{
		CurPage:  curPage,
		PageSize: pageSize,
	})

	return dto.toProto(), err
}

func (s *categoryServer) ListCategories(ctx context.Context, amount *rpc.Amount) (*rpc.CategoriesWithRelatedArticleAmount, error) {

	am, err := cmprimitive.NewAmount(int(amount.GetAmount()))
	if err != nil {
		return nil, err
	}

	cas, err := s.category.ListCategories(am)
	if err != nil {
		return nil, err
	}

	categories := make([]*rpc.CategoryWithRelatedArticleAmount, len(cas))
	for i := range categories {
		categories[i] = toProtoCategoryWithRelatedArticleAmount(cas[i])
	}

	return &rpc.CategoriesWithRelatedArticleAmount{
		Categories: categories,
	}, nil
}

func (s *categoryServer) GetArticleCategory(ctx context.Context, id *rpc.Id) (*rpc.Category, error) {

	cate, err := s.category.GetArticleCategory(cmprimitive.NewId(id.GetId()))
	if err != nil {
		return nil, err
	}

	return &rpc.Category{
		Id:        cate.Id.Id(),
		Name:      cate.Name.CategoryName(),
		CreatedAt: cate.CreatedAt.TimeUnix(),
	}, nil
}

func (s *categoryServer) GetTotalNumberOfCategories(ctx context.Context, empty *emptypb.Empty) (*rpc.Amount, error) {

	amount, err := s.category.GetTotalNumberOfCategories()
	if err != nil {
		return nil, err
	}

	return &rpc.Amount{Amount: int64(amount.Amount())}, nil
}

func (s *categoryServer) DelCategory(ctx context.Context, id *rpc.Id) (*emptypb.Empty, error) {
	return nil, s.category.DelCategory(cmprimitive.NewId(id.GetId()))
}

func (s *categoryServer) GetRelationWithArticle(ctx context.Context, id *rpc.Id) (*rpc.Id, error) {

	identity, err := s.category.GetRelationWithArticle(cmprimitive.NewId(id.GetId()))
	if err != nil {
		return nil, err
	}

	return &rpc.Id{Id: identity.Id()}, nil
}

func (s *categoryServer) GetRelatedArticleIdsThroughCateId(ctx context.Context, id *rpc.Id) (*rpc.RespGetRelatedArticleIdsThroughCateId, error) {

	articleIds, err := s.category.GetRelatedArticleIdsThroughCateId(cmprimitive.NewId(id.GetId()))
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

func (s *categoryServer) BuildRelationWithArticle(ctx context.Context, req *rpc.ArticleIdAndCateId) (*emptypb.Empty, error) {
	return nil, s.category.BuildRelationWithArticle(
		cmprimitive.NewId(req.GetArticleId().GetId()),
		cmprimitive.NewId(req.GetCateId().GetId()),
	)
}

func (s *categoryServer) RemoveRelationWithArticle(ctx context.Context, id *rpc.Id) (*emptypb.Empty, error) {
	return nil, s.category.RemoveRelationWithArticle(cmprimitive.NewId(id.GetId()))
}
