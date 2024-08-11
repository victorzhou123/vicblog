package service

import (
	"mime/multipart"

	"victorzhou123/vicblog/article/domain/picture/entity"
	"victorzhou123/vicblog/article/domain/picture/repository"
	cmdmerror "victorzhou123/vicblog/common/domain/error"
	cmprimitive "victorzhou123/vicblog/common/domain/primitive"
	"victorzhou123/vicblog/common/log"
)

type FileService interface {
	Upload(cmprimitive.Username, *multipart.FileHeader) (FileUrlDto, error)
}

type fileService struct {
	repo repository.Picture
}

func NewFileService(repo repository.Picture) FileService {
	return &fileService{
		repo: repo,
	}
}

func (s *fileService) Upload(
	user cmprimitive.Username, file *multipart.FileHeader,
) (FileUrlDto, error) {

	pictureName, err := entity.NewPictureName(file.Filename)
	if err != nil {
		return FileUrlDto{}, err
	}

	mf, err := file.Open()
	if err != nil {
		return FileUrlDto{}, err
	}

	picture := entity.Picture{
		Name: pictureName,
		Data: mf,
		Size: file.Size,
	}

	if picture.OverSizeLimited() {

		log.Errorf("user %s upload file over size limited", user.Username())

		return FileUrlDto{}, cmdmerror.NewInvalidParam("picture size over limited")
	}

	url, err := s.repo.Upload(user.Username(), file.Filename, picture)
	if err != nil {

		log.Errorf("user %s upload file failed, err: %s", user.Username(), err.Error())

		return FileUrlDto{}, err
	}

	return FileUrlDto{url}, nil
}
