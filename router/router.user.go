package route

import (
	handlerDeleteUser "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/handlers/user/delete"
	handlerUpdateUser "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/handlers/user/update"
	deleteUser "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/services/user/delete"
	updateUser "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/services/user/update"
	middleware "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/middlewares"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler User
	*/
	updateUserRepository := updateUser.NewRepositoryUpdate(db)
	updateUserService := updateUser.NewServiceUpdate(updateUserRepository)
	updateUserController := handlerUpdateUser.NewHandlerUpdateUser(updateUserService)

	deleteUserRepository := deleteUser.NewRepositoryDelete(db)
	deleteUserService := deleteUser.NewServiceDelete(deleteUserRepository)
	deleteUserController := handlerDeleteUser.NewHandlerDeleteUser(deleteUserService)

	/**
	@description All User Route
	*/
	groupRoute := route.Group("/api/v1").Use(middleware.Auth())
	groupRoute.PUT("/users/:userId", updateUserController.UpdateUserHandler)
	groupRoute.DELETE("/users/:userId", deleteUserController.DeleteUserHandler)
}
