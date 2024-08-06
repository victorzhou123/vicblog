package repositoryimpl

import (
	"victorzhou123/vicblog/article/domain/tag/entity"
	"victorzhou123/vicblog/article/domain/tag/repository"
	cmdmerror "victorzhou123/vicblog/common/domain/error"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
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

func (impl *tagRepoImpl) GetTagList(opt cmrepo.PageListOpt) ([]entity.Tag, int, error) {
	dos := []TagDO{}

	options := mysql.PaginationOpt{
		CurPage:  opt.CurPage,
		PageSize: opt.PageSize,
	}

	total, err := impl.GetRecordByPagination(&TagDO{}, &dos, options)
	if err != nil {
		if cmrepo.IsErrorResourceNotExists(err) {
			return nil, 0, cmdmerror.NewNotFound(cmdmerror.ErrorCodeResourceNotFound, "")
		}

		return nil, 0, err
	}

	tags := make([]entity.Tag, len(dos))
	for i := range dos {
		if tags[i], err = dos[i].toTag(); err != nil {
			return nil, 0, err
		}
	}

	return tags, total, nil
}

func (impl *tagRepoImpl) Delete(id cmprimitive.Id) error {
	do := TagDO{}
	do.ID = id.IdNum()

	err := impl.Impl.Delete(&TagDO{}, &do)
	if cmrepo.IsErrorResourceNotExists(err) {
		return cmdmerror.NewNotFound(cmdmerror.ErrorCodeResourceNotFound, "")
	}

	return err
}
