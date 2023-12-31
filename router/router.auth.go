package route

import (
	handlerLogin "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/handlers/auth/login"
	handlerRegister "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/handlers/auth/register"
	loginAuth "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/services/auth/login"
	registerAuth "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/services/auth/register"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	/**
	@description All Handler Auth
	*/
	LoginRepository := loginAuth.NewRepositoryLogin(db)
	loginService := loginAuth.NewServiceLogin(LoginRepository)
	loginController := handlerLogin.NewHandlerLogin(loginService)

	registerRepository := registerAuth.NewRepositoryRegister(db)
	registerService := registerAuth.NewServiceRegister(registerRepository)
	registerController := handlerRegister.NewHandlerRegister(registerService)

	/**
	@description All Auth Route
	*/
	groupRoute := route.Group("/api/v1")
	groupRoute.POST("/register", registerController.RegisterHandler)
	groupRoute.POST("/login", loginController.LoginHandler)
}
