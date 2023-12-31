package resultPhoto

import (
	model "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/models"
)

type Service interface {
	ResultPhotoService(input *InputResultPhoto) (*model.EntityPhotos, string)
}

type service struct {
	repository Repository
}

func NewServiceResults(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ResultPhotoService(input *InputResultPhoto) (*model.EntityPhotos, string) {

	photo := model.EntityPhotos{
		UserId: input.UserId,
	}

	resultPhoto, errPhoto := s.repository.ResultPhotoRepository(&photo)

	return resultPhoto, errPhoto
}
