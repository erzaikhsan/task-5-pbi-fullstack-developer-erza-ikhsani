package resultPhoto

import (
	model "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/models"
	"gorm.io/gorm"
)

type Repository interface {
	ResultPhotoRepository(input *model.EntityPhotos) (*model.EntityPhotos, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryResult(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ResultPhotoRepository(input *model.EntityPhotos) (*model.EntityPhotos, string) {

	var photo model.EntityPhotos
	db := r.db.Model(&photo)
	errorCode := make(chan string, 1)

	resultPhoto := db.Debug().Select("*").Where("user_id = ?", input.UserId).Find(&photo)

	if resultPhoto.RowsAffected < 1 {
		errorCode <- "RESULT_PHOTO_NOT_FOUND_404"
		return &photo, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &photo, <-errorCode
}
