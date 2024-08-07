package repositoryimpl

import (
	"fmt"
	"victorzhou123/vicblog/article/domain/picture/entity"
	"victorzhou123/vicblog/article/domain/picture/repository"
	"victorzhou123/vicblog/common/infrastructure/oss"
)

func NewPictureImpl(o oss.OssService) repository.Picture {
	return &pictureImpl{o}
}

type pictureImpl struct {
	oss.OssService
}

func (impl *pictureImpl) Upload(username, pictureName string, picture entity.Picture) (string, error) {
	position := fmt.Sprintf("%s/%s", username, pictureName)

	return impl.UploadPicture(position, picture.Data)
}
