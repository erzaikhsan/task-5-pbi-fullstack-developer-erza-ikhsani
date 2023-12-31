package updateUser

import (
	model "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/models"
	"gorm.io/gorm"
)

type Repository interface {
	UpdateUserRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateUserRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.ID = input.ID

	checkUserAccount := db.Debug().Select("*").Where("id = ?", input.ID).Find(&users)

	if checkUserAccount.RowsAffected < 1 {
		errorCode <- "UPDATE_USER_NOT_FOUND_404"
		return &users, <-errorCode
	}

	users.ID = input.ID
	users.Username = input.Username
	users.Email = input.Email
	users.Password = input.Password
	users.UpdatedAt = input.UpdatedAt

	updateUser := db.Debug().Select("username", "email", "password", "updated_at").Where("id = ?", input.ID).Updates(users)
	db.Commit()

	if updateUser.Error != nil {
		errorCode <- "UPDATE_USER_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}
