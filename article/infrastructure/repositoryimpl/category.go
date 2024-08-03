package repositoryimpl

import (
	"victorzhou123/vicblog/article/domain/category/entity"
	"victorzhou123/vicblog/article/domain/category/repository"
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
	categoryDo := &CategoryDO{}
	categoryDo.Name = name.CategoryName()

	return impl.Add(&categoryDo)
}
