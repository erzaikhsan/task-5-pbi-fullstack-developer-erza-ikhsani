package route

import (
	handlerCreatePhoto "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/handlers/photo/create"
	handlerDeletePhoto "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/handlers/photo/delete"
	handlerResultPhoto "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/handlers/photo/result"
	handlerUpdatePhoto "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/handlers/photo/update"
	createPhoto "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/services/photo/create"
	deletePhoto "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/services/photo/delete"
	resultPhoto "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/services/photo/result"
	updatePhoto "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/services/photo/update"
	middleware "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitPhotoRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler User
	*/
	createPhotoRepository := createPhoto.NewRepositoryCreate(db)
	createPhotoService := createPhoto.NewServiceCreate(createPhotoRepository)
	createPhotoController := handlerCreatePhoto.NewHandlerCreatePhoto(createPhotoService)

	resultPhotoRepository := resultPhoto.NewRepositoryResult(db)
	resultPhotoService := resultPhoto.NewServiceResults(resultPhotoRepository)
	resultPhotoController := handlerResultPhoto.NewHandlerResultPhoto(resultPhotoService)

	deletePhotoRepository := deletePhoto.NewRepositoryDelete(db)
	deletePhotoService := deletePhoto.NewServiceDelete(deletePhotoRepository)
	deletePhotoController := handlerDeletePhoto.NewHandlerDeletePhoto(deletePhotoService)

	updatePhotoRepository := updatePhoto.NewRepositoryUpdate(db)
	updatePhotoService := updatePhoto.NewServiceUpdate(updatePhotoRepository)
	updatePhotoController := handlerUpdatePhoto.NewHandlerUpdatePhoto(updatePhotoService)

	/**
	@description All User Route
	*/
	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.POST("/photos", createPhotoController.CreatePhotoHandler)
	groupRoute.GET("/photos", resultPhotoController.ResultPhotoHandler)
	groupRoute.DELETE("/photos/:photoId", deletePhotoController.DeletePhotoHandler)
	groupRoute.PUT("/photos/:photoId", updatePhotoController.UpdatePhotoHandler)
}
