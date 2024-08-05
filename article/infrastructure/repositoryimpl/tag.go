package repositoryimpl

import (
	"victorzhou123/vicblog/article/domain/tag/repository"
	cmdmerror "victorzhou123/vicblog/common/domain/error"
	cmrepo "victorzhou123/vicblog/common/domain/repository"
	"victorzhou123/vicblog/common/infrastructure/mysql"
)

func NewTagRepo(db mysql.Impl) repository.Tag {
	tableNameTag = db.TableName()

	if err := mysql.AutoMigrate(&TagDO{}); err != nil {
		return nil
	}

	return &tagRepoImpl{db}
}

type tagRepoImpl struct {
	mysql.Impl
}

func (impl *tagRepoImpl) AddBatches(tagNames repository.TagNames) error {
	names := tagNames.Names

	dos := make([]TagDO, len(names))
	for i := range names {
		dos[i].Name = names[i].TagName()
	}

	err := impl.Impl.Add(&dos)
	if cmrepo.IsErrorConstraintViolated(err) || cmrepo.IsErrorDuplicateCreating(err) {
		return cmdmerror.NewInvalidParam(err.Error())
	}

	return err
}
