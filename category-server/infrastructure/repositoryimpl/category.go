package repositoryimpl

import (
	"github.com/victorzhou123/vicblog/category-server/domain/category/entity"
	"github.com/victorzhou123/vicblog/category-server/domain/category/repository"
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
	cmdmerror "github.com/victorzhou123/vicblog/common/domain/error"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	"github.com/victorzhou123/vicblog/common/infrastructure/mysql"
)

func NewCategoryRepo(db mysql.Impl) repository.Category {

	if err := mysql.AutoMigrate(&CategoryDO{}, &CategoryArticleDO{}); err != nil {
		return nil
	}

	return &categoryRepoImpl{db}
}

type categoryRepoImpl struct {
	mysql.Impl
}

func (impl *categoryRepoImpl) AddCategory(name entity.CategoryName) error {
	categoryDo := CategoryDO{}
	categoryDo.Name = name.CategoryName()

	return impl.Add(&CategoryDO{}, &categoryDo)
}

func (impl *categoryRepoImpl) GetCategory(cateId cmprimitive.Id) (entity.Category, error) {

	do := CategoryDO{}

	if err := impl.Impl.GetByPrimaryKey(&CategoryDO{}, &do); err != nil {
		return entity.Category{}, err
	}

	return do.toCategory()
}

func (impl *categoryRepoImpl) GetCategoryListByPagination(opt cment.Pagination) ([]entity.Category, int, error) {
	categoryDos := []CategoryDO{}

	options := mysql.PaginationOpt{
		CurPage:  opt.CurPage.CurPage(),
		PageSize: opt.PageSize.PageSize(),
	}

	total, err := impl.GetRecordsByPagination(&CategoryDO{}, &CategoryDO{}, &categoryDos, options)
	if err != nil {
		return nil, 0, err
	}

	cates := make([]entity.Category, len(categoryDos))
	for i := range categoryDos {
		if cates[i], err = (categoryDos)[i].toCategory(); err != nil {
			return nil, 0, err
		}
	}

	return cates, total, nil
}

func (impl *categoryRepoImpl) GetCategoryList(amount cmprimitive.Amount) ([]entity.Category, error) {
	// convert amount to size
	var size int
	if amount == nil {
		size = -1
	} else {
		size = amount.Amount()
	}

	// get category records
	dos := []CategoryDO{}
	if err := impl.GetLimitRecords(&CategoryDO{}, &CategoryDO{}, &dos, size); err != nil {
		if cmdmerror.IsNotFound(err) {
			return []entity.Category{}, nil
		}

		return nil, err
	}

	cates := make([]entity.Category, len(dos))

	var err error
	for i := range dos {
		if cates[i], err = dos[i].toCategory(); err != nil {
			return nil, err
		}
	}

	return cates, err
}

func (impl *categoryRepoImpl) GetTotalNumberOfCategories() (cmprimitive.Amount, error) {
	var count int64
	if err := impl.Impl.GetTotalNumber(&CategoryDO{}, &CategoryDO{}, &count); err != nil {
		return nil, err
	}

	return cmprimitive.NewAmount(int(count))
}

func (impl *categoryRepoImpl) DelCategory(id cmprimitive.Id) error {
	do := &CategoryDO{}
	do.ID = id.IdNum()

	if err := impl.DeleteByPrimaryKey(&CategoryDO{}, &do); err != nil {
		return cmdmerror.New(cmdmerror.ErrorCodeResourceNotFound, "")
	}

	return nil
}
