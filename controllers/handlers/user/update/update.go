package handlerUpdateUser

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	updateUser "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/services/user/update"
	helper "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/helpers"
	goValidator "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/helpers/goValidator"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service updateUser.Service
}

func NewHandlerUpdateUser(service updateUser.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdateUserHandler(ctx *gin.Context) {

	var input updateUser.InputUpdateUser
	input.ID = ctx.Param("userId")
	ctx.ShouldBindJSON(&input)

	config := goValidator.ErrorConfig{
		Options: []goValidator.ErrorMetaConfig{
			goValidator.ErrorMetaConfig{
				Tag:     "required",
				Field:   "ID",
				Message: "id is required on param",
			},
			goValidator.ErrorMetaConfig{
				Tag:     "uuid",
				Field:   "ID",
				Message: "params must be uuid format",
			},
			goValidator.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Username",
				Message: "username is required on body",
			},
			goValidator.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Email",
				Message: "email is required on body",
			},
			goValidator.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Password",
				Message: "password is required on body",
			},
		},
	}

	errResponse, errCount := helper.GoValidator(&input, config.Options)

	if errCount > 0 {
		helper.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}

	token, _ := ctx.Get("user")
	claims, ok := token.(jwt.MapClaims)
	if !ok {
		helper.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}
	id, idExists := claims["id"].(string)
	if !idExists {
		helper.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}

	if input.ID != id {
		helper.APIResponse(ctx, "Update access denied", http.StatusForbidden, http.MethodPost, nil)
		return
	}

	_, errUpdateUser := h.service.UpdateUserService(&input)

	switch errUpdateUser {

	case "UPDATE_USER_NOT_FOUND_404":
		helper.APIResponse(ctx, "User data is not exist or deleted", http.StatusNotFound, http.MethodPost, nil)

	case "UPDATE_USER_FAILED_403":
		helper.APIResponse(ctx, "Update user data failed", http.StatusForbidden, http.MethodPost, nil)

	default:
		helper.APIResponse(ctx, "Update user data sucessfully", http.StatusOK, http.MethodPost, nil)
	}
}
