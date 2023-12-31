package handlerDeleteUser

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	deleteUser "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/services/user/delete"
	helper "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/helpers"
	goValidator "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/helpers/goValidator"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service deleteUser.Service
}

func NewHandlerDeleteUser(service deleteUser.Service) *handler {
	return &handler{service: service}
}

func (h *handler) DeleteUserHandler(ctx *gin.Context) {

	var input deleteUser.InputDeleteUser
	input.ID = ctx.Param("userId")

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
		},
	}

	errResponse, errCount := helper.GoValidator(&input, config.Options)

	if errCount > 0 {
		helper.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodDelete, errResponse)
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
		helper.APIResponse(ctx, "Delete access denied", http.StatusForbidden, http.MethodPost, nil)
		return
	}

	_, errDeleteUser := h.service.DeleteUserService(&input)

	switch errDeleteUser {

	case "DELETE_USER_NOT_FOUND_404":
		helper.APIResponse(ctx, "User data is not exist or deleted", http.StatusForbidden, http.MethodDelete, nil)
		return

	case "DELETE_USER_FAILED_403":
		helper.APIResponse(ctx, "Delete user data failed", http.StatusForbidden, http.MethodDelete, nil)
		return

	default:
		helper.APIResponse(ctx, "Delete user data successfully", http.StatusOK, http.MethodDelete, nil)
	}
}
