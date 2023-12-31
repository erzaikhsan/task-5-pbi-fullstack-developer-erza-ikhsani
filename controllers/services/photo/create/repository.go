package createPhoto

import (
	model "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreatePhotoRepository(input *model.EntityPhotos) (*model.EntityPhotos, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreatePhotoRepository(input *model.EntityPhotos) (*model.EntityPhotos, string) {

	var photos model.EntityPhotos
	db := r.db.Model(&photos)
	errorCode := make(chan string, 1)

	checkPhotoExist := db.Debug().Select("*").Where("user_id = ?", input.UserId).Find(&photos)

	if checkPhotoExist.RowsAffected > 0 {
		errorCode <- "CREATE_PHOTO_CONFLICT_409"
		return &photos, <-errorCode
	}

	photos.Title = input.Title
	photos.Caption = input.Caption
	photos.PhotoUrl = input.PhotoUrl
	photos.UserId = input.UserId

	addNewPhoto := db.Debug().Create(&photos)
	db.Commit()

	if addNewPhoto.Error != nil {
		errorCode <- "CREATE_PHOTO_FAILED_403"
		return &photos, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &photos, <-errorCode
}
