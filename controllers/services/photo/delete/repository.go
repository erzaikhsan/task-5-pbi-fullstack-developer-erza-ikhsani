package deletePhoto

import (
	model "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/models"
	"gorm.io/gorm"
)

type Repository interface {
	DeletePhotoRepository(input *model.EntityPhotos) (*model.EntityPhotos, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryDelete(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) DeletePhotoRepository(input *model.EntityPhotos) (*model.EntityPhotos, string) {

	var photo model.EntityPhotos
	db := r.db.Model(&photo)
	errorCode := make(chan string, 1)

	checkPhotoId := db.Debug().Select("*").Where("id = ?", input.ID).Find(&photo)

	if checkPhotoId.RowsAffected < 1 {
		errorCode <- "DELETE_PHOTO_NOT_FOUND_404"
		return &photo, <-errorCode
	}

	checkPhotoAccess := db.Debug().Select("*").Where("id = ?", input.ID).Where("user_id = ? ", input.UserId).Find(&photo)

	if checkPhotoAccess.RowsAffected < 1 {
		errorCode <- "DELETE_PHOTO_FAILED_403"
		return &photo, <-errorCode
	}

	deletePhotoId := db.Debug().Select("*").Where("id = ?", input.ID).Find(&photo).Delete(&photo)

	if deletePhotoId.Error != nil {
		errorCode <- "DELETE_PHOTO_FAILED_403"
		return &photo, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &photo, <-errorCode
}
