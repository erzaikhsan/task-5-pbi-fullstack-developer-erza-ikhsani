package deletePhoto

import (
	model "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/models"
)

type Service interface {
	DeletePhotoService(input *InputDeletePhoto) (*model.EntityPhotos, string)
}

type service struct {
	repository Repository
}

func NewServiceDelete(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) DeletePhotoService(input *InputDeletePhoto) (*model.EntityPhotos, string) {

	photo := model.EntityPhotos{
		ID:     input.ID,
		UserId: input.UserId,
	}

	resultDeletePhoto, errDeletePhoto := s.repository.DeletePhotoRepository(&photo)

	return resultDeletePhoto, errDeletePhoto
}
