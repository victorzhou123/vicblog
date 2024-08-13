package repositoryimpl

import (
	"victorzhou123/vicblog/article/domain/tag/entity"
	"victorzhou123/vicblog/article/domain/tag/repository"
	cment "victorzhou123/vicblog/common/domain/entity"
	cmdmerror "victorzhou123/vicblog/common/domain/error"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	cmrepo "victorzhou123/vicblog/common/domain/repository"
	"victorzhou123/vicblog/common/infrastructure/mysql"
)

func NewTagRepo(db mysql.Impl) repository.Tag {

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

	err := impl.Impl.Add(&TagDO{}, &dos)
	if cmrepo.IsErrorConstraintViolated(err) || cmrepo.IsErrorDuplicateCreating(err) {
		return cmdmerror.NewInvalidParam(err.Error())
	}

	return err
}

func (impl *tagRepoImpl) GetBatchTags(tagIds []cmprimitive.Id) ([]entity.Tag, error) {

	ids := make([]uint, len(tagIds))
	for i := range ids {
		ids[i] = tagIds[i].IdNum()
	}

	dos := []TagDO{}

	if err := impl.Impl.Model(&TagDO{}).Where(impl.InFilter("id"), ids).Find(&dos).Error; err != nil {
		return nil, cmdmerror.NewNotFound(cmdmerror.ErrorCodeResourceNotFound, "")
	}

	var err error
	tags := make([]entity.Tag, len(dos))
	for i := range dos {
		if tags[i], err = dos[i].toTag(); err != nil {
			return nil, err
		}
	}

	return tags, nil
}

func (impl *tagRepoImpl) GetTagListByPagination(opt cment.Pagination) ([]entity.Tag, int, error) {
	dos := []TagDO{}

	options := mysql.PaginationOpt{
		CurPage:  opt.CurPage.CurPage(),
		PageSize: opt.PageSize.PageSize(),
	}

	total, err := impl.GetRecordsByPagination(&TagDO{}, &TagDO{}, &dos, options)
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

func (impl *tagRepoImpl) GetAllTagList() ([]entity.Tag, error) {
	dos := []TagDO{}

	if err := impl.GetRecords(&TagDO{}, &TagDO{}, &dos); err != nil {
		if cmdmerror.IsNotFound(err) {
			return []entity.Tag{}, nil
		}

		return nil, err
	}

	tags := make([]entity.Tag, len(dos))

	var err error
	for i := range dos {
		if tags[i], err = dos[i].toTag(); err != nil {
			return nil, err
		}
	}

	return tags, err
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
