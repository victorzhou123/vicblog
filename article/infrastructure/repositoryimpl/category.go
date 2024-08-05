package repositoryimpl

import (
	"victorzhou123/vicblog/article/domain/category/entity"
	"victorzhou123/vicblog/article/domain/category/repository"
	cmdmerror "victorzhou123/vicblog/common/domain/error"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	cmrepo "victorzhou123/vicblog/common/domain/repository"
	"victorzhou123/vicblog/common/infrastructure/mysql"
)

func NewCategoryRepo(db mysql.Impl) repository.Category {
	tableNameCategory = db.TableName()

	if err := mysql.AutoMigrate(&CategoryDO{}); err != nil {
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

	return impl.Add(&categoryDo)
}

func (impl *categoryRepoImpl) GetCategoryList(opt cmrepo.PageListOpt) ([]entity.Category, int, error) {
	categoryDos := []CategoryDO{}

	options := mysql.PaginationOpt{
		CurPage:  opt.CurPage,
		PageSize: opt.PageSize,
	}

	total, err := impl.GetRecordByPagination(&CategoryDO{}, &categoryDos, options)
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

func (impl *categoryRepoImpl) DelCategory(id cmprimitive.Id) error {
	do := &CategoryDO{}
	do.ID = id.IdNum()

	if err := impl.DeleteByPrimaryKey(&do); err != nil {
		return cmdmerror.New(cmdmerror.ErrorCodeResourceNotFound, "")
	}

	return nil
}
