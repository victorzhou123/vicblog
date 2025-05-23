package repositoryimpl

import (
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
	cmdmerror "github.com/victorzhou123/vicblog/common/domain/error"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
	cmrepo "github.com/victorzhou123/vicblog/common/domain/repository"
	"github.com/victorzhou123/vicblog/common/infrastructure/mysql"
	"github.com/victorzhou123/vicblog/tag-server/domain/tag/entity"
	"github.com/victorzhou123/vicblog/tag-server/domain/tag/repository"
)

func NewTagRepo(db mysql.Impl) repository.Tag {

	if err := mysql.AutoMigrate(&TagDO{}, &TagArticleDO{}); err != nil {
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

func (impl *tagRepoImpl) GetTagList(amount cmprimitive.Amount) ([]entity.Tag, error) {
	// convert amount to limit
	var limit int
	if amount == nil {
		limit = -1
	} else {
		limit = amount.Amount()
	}

	dos := []TagDO{}

	if err := impl.GetLimitRecords(&TagDO{}, &TagDO{}, &dos, limit); err != nil {
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

func (impl *tagRepoImpl) GetTotalNumberOfTag() (cmprimitive.Amount, error) {
	var count int64
	if err := impl.Impl.GetTotalNumber(&TagDO{}, &TagDO{}, &count); err != nil {
		return nil, err
	}

	return cmprimitive.NewAmount(int(count))
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
