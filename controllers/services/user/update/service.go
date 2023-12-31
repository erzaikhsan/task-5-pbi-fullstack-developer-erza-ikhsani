package updateUser

import (
	helper "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/helpers"
	model "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/models"
)

type Service interface {
	UpdateUserService(input *InputUpdateUser) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceUpdate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) UpdateUserService(input *InputUpdateUser) (*model.EntityUsers, string) {

	hashPassword := helper.HashPassword(input.Password)
	users := model.EntityUsers{
		ID:       input.ID,
		Username: input.Username,
		Email:    input.Email,
		Password: hashPassword,
	}

	resultUpdateUser, errUpdateUser := s.repository.UpdateUserRepository(&users)

	return resultUpdateUser, errUpdateUser
}
