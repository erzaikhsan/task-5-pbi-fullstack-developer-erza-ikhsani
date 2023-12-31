package deleteUser

import (
	model "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/models"
	"gorm.io/gorm"
)

type Repository interface {
	DeleteUserRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryDelete(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) DeleteUserRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	checkUserId := db.Debug().Select("*").Where("id = ?", input.ID).Find(&users)

	if checkUserId.RowsAffected < 1 {
		errorCode <- "DELETE_USER_NOT_FOUND_404"
		return &users, <-errorCode
	}

	deleteUserId := db.Debug().Select("*").Where("id = ?", input.ID).Find(&users).Delete(&users)

	if deleteUserId.Error != nil {
		errorCode <- "DELETE_USER_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
