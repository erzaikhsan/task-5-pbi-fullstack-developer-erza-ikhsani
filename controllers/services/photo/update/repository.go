package updatePhoto

import (
	model "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/models"
	"gorm.io/gorm"
)

type Repository interface {
	UpdatePhotoRepository(input *model.EntityPhotos) (*model.EntityPhotos, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdatePhotoRepository(input *model.EntityPhotos) (*model.EntityPhotos, string) {

	var photo model.EntityPhotos
	db := r.db.Model(&photo)
	errorCode := make(chan string, 1)

	photo.ID = input.ID

	checkPhotoId := db.Debug().Select("*").Where("id = ?", input.ID).Find(&photo)

	if checkPhotoId.RowsAffected < 1 {
		errorCode <- "UPDATE_PHOTO_NOT_FOUND_404"
		return &photo, <-errorCode
	}

	checkPhotoAccess := db.Debug().Select("*").Where("id = ?", input.ID).Where("user_id = ? ", input.UserId).Find(&photo)

	if checkPhotoAccess.RowsAffected < 1 {
		errorCode <- "UPDATE_PHOTO_FAILED_403"
		return &photo, <-errorCode
	}

	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoUrl = input.PhotoUrl
	photo.UpdatedAt = input.UpdatedAt

	updatePhoto := db.Debug().Select("title", "caption", "photo_url", "updated_at").Where("id = ?", input.ID).Updates(photo)

	if updatePhoto.Error != nil {
		errorCode <- "UPDATE_PHOTO_FAILED_403"
		return &photo, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &photo, <-errorCode
}
