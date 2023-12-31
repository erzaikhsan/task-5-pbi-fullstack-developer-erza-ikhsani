package deleteUser

import (
	model "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/models"
)

type Service interface {
	DeleteUserService(input *InputDeleteUser) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceDelete(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) DeleteUserService(input *InputDeleteUser) (*model.EntityUsers, string) {

	users := model.EntityUsers{
		ID: input.ID,
	}

	resultCreateUser, errCreateUser := s.repository.DeleteUserRepository(&users)

	return resultCreateUser, errCreateUser
}
