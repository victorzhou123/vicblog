package service

import (
	"context"
	"strconv"
	"time"

	"github.com/victorzhou123/vicblog/category-server/domain/category/entity"
	"github.com/victorzhou123/vicblog/common/controller/rpc"
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"google.golang.org/protobuf/types/known/emptypb"
)

type categoryClient struct {
	expire time.Duration
	rpcSvc rpc.CategoryServiceServer
}

func NewCategoryServer(rpcSvc rpc.CategoryServiceServer) CategoryService {
	return &categoryClient{
		rpcSvc: rpcSvc,
	}
}

func (s *categoryClient) AddCategory(name entity.CategoryName) error {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	_, err := s.rpcSvc.AddCategory(ctx, &rpc.CategoryName{
		Name: name.CategoryName(),
	})

	return err
}

func (s *categoryClient) ListCategoryByPagination(pagination *cment.Pagination) (CategoryListDto, error) {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	categoryList, err := s.rpcSvc.ListCategoryByPagination(ctx, &rpc.Pagination{
		CurPage:  strconv.Itoa(pagination.CurPage.CurPage()),
		PageSize: strconv.Itoa(pagination.PageSize.PageSize()),
	})
	if err != nil {
		return CategoryListDto{}, err
	}

	return toCategoryListDto(categoryList)
}

func (s *categoryClient) ListCategories(amount cmprimitive.Amount) ([]entity.CategoryWithRelatedArticleAmount, error) {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	categories, err := s.rpcSvc.ListCategories(ctx, &rpc.Amount{Amount: int64(amount.Amount())})
	if err != nil {
		return nil, err
	}

	cs := make([]entity.CategoryWithRelatedArticleAmount, len(categories.Categories))
	for i := range cs {
		cs[i] = toCategoryWithRelatedArticleAmount(categories.Categories[i])
	}

	return cs, nil
}

func (s *categoryClient) GetArticleCategory(articleId cmprimitive.Id) (entity.Category, error) {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	category, err := s.rpcSvc.GetArticleCategory(ctx, &rpc.Id{Id: articleId.Id()})
	if err != nil {
		return entity.Category{}, err
	}

	categoryName, err := entity.NewCategoryName(category.GetName())
	if err != nil {
		return entity.Category{}, err
	}

	return entity.Category{
		Id:        cmprimitive.NewId(category.Id),
		Name:      categoryName,
		CreatedAt: cmprimitive.NewTimeXWithUnix(category.GetCreatedAt()),
	}, nil
}

func (s *categoryClient) GetTotalNumberOfCategories() (cmprimitive.Amount, error) {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	amount, err := s.rpcSvc.GetTotalNumberOfCategories(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	return cmprimitive.NewAmount(int(amount.GetAmount()))
}

func (s *categoryClient) DelCategory(id cmprimitive.Id) error {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	_, err := s.rpcSvc.DelCategory(ctx, &rpc.Id{Id: id.Id()})

	return err
}

func (s *categoryClient) GetRelationWithArticle(articleId cmprimitive.Id) (cateId cmprimitive.Id, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	id, err := s.rpcSvc.GetRelationWithArticle(ctx, &rpc.Id{Id: articleId.Id()})
	if err != nil {
		return nil, err
	}

	return cmprimitive.NewId(id.GetId()), nil
}

func (s *categoryClient) GetRelatedArticleIdsThroughCateId(cateId cmprimitive.Id) ([]cmprimitive.Id, error) {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	resp, err := s.rpcSvc.GetRelatedArticleIdsThroughCateId(ctx, &rpc.Id{Id: cateId.Id()})
	if err != nil {
		return nil, err
	}

	ids := make([]cmprimitive.Id, len(resp.Ids))
	for i := range ids {
		ids[i] = cmprimitive.NewId(resp.Ids[i].GetId())
	}

	return ids, nil
}

func (s *categoryClient) BuildRelationWithArticle(articleId, cateId cmprimitive.Id) error {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	_, err := s.rpcSvc.BuildRelationWithArticle(ctx, &rpc.ReqBuildRelationWithArticle{
		ArticleId: &rpc.Id{Id: articleId.Id()},
		CateId:    &rpc.Id{Id: cateId.Id()},
	})

	return err
}

func (s *categoryClient) RemoveRelationWithArticle(articleId cmprimitive.Id) error {

	ctx, cancel := context.WithTimeout(context.Background(), s.expire)
	defer cancel()

	_, err := s.rpcSvc.RemoveRelationWithArticle(ctx, &rpc.Id{Id: articleId.Id()})

	return err
}
