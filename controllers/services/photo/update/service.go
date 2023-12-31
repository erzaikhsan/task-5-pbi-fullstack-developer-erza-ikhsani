package updatePhoto

import (
	model "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/models"
)

type Service interface {
	UpdatePhotoService(input *InputUpdatePhoto) (*model.EntityPhotos, string)
}

type service struct {
	repository Repository
}

func NewServiceUpdate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdatePhotoService(input *InputUpdatePhoto) (*model.EntityPhotos, string) {

	photo := model.EntityPhotos{
		ID:       input.ID,
		Title:    input.Title,
		Caption:  input.Caption,
		PhotoUrl: input.PhotoUrl,
		UserId:   input.UserId,
	}

	resultUpdatePhoto, errUpdatePhoto := s.repository.UpdatePhotoRepository(&photo)

	return resultUpdatePhoto, errUpdatePhoto
}
