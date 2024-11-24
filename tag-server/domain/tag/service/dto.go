package service

import (
	"github.com/victorzhou123/vicblog/common/controller/rpc"
	cmdto "github.com/victorzhou123/vicblog/common/domain/dto"
	cment "github.com/victorzhou123/vicblog/common/domain/entity"
	"github.com/victorzhou123/vicblog/tag-server/domain/tag/entity"
)

type TagListDto struct {
	cment.PaginationStatus

	Tags []entity.TagWithRelatedArticleAmount
}

func (dto *TagListDto) toProto() *rpc.TagList {

	tags := make([]*rpc.TagWithRelatedArticleAmount, len(dto.Tags))
	for i := range tags {
		tags[i] = toProtoTagWithRelatedArticleAmount(dto.Tags[i])
	}

	return &rpc.TagList{
		PaginationStatus: cmdto.ToProtoPaginationStatus(dto.PaginationStatus),
		Tags:             tags,
	}
}

func toProtoTagWithRelatedArticleAmount(ta entity.TagWithRelatedArticleAmount) *rpc.TagWithRelatedArticleAmount {
	return &rpc.TagWithRelatedArticleAmount{
		Tag:                  toProtoTag(ta.Tag),
		RelatedArticleAmount: int64(ta.RelatedArticleAmount.Amount()),
	}
}

func toProtoTag(tag entity.Tag) *rpc.Tag {
	return &rpc.Tag{
		Id:        tag.Id.Id(),
		Name:      tag.Name.TagName(),
		CreatedAt: tag.CreatedAt.TimeUnix(),
	}
}

func toProtoTagNames(names *rpc.TagNames) ([]entity.TagName, error) {
	var err error
	tagNames := make([]entity.TagName, len(names.Names))
	for i := range tagNames {
		if tagNames[i], err = entity.NewTagName(names.Names[i].GetName()); err != nil {
			return nil, err
		}
	}

	return tagNames, nil
}
