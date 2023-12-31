package createPhoto

import (
	model "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/models"
)

type Service interface {
	CreatePhotoService(input *InputCreatePhoto) (*model.EntityPhotos, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreatePhotoService(input *InputCreatePhoto) (*model.EntityPhotos, string) {

	photos := model.EntityPhotos{
		Title:    input.Title,
		Caption:  input.Caption,
		PhotoUrl: input.PhotoUrl,
		UserId:   input.UserId,
	}

	resultCreatePhoto, errCreatePhoto := s.repository.CreatePhotoRepository(&photos)

	return resultCreatePhoto, errCreatePhoto
}
