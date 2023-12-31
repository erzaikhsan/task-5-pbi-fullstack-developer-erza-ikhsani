package handlerDeletePhoto

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	deletePhoto "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/services/photo/delete"
	helper "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/helpers"
	goValidator "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/helpers/goValidator"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service deletePhoto.Service
}

func NewHandlerDeletePhoto(service deletePhoto.Service) *handler {
	return &handler{service: service}
}

func (h *handler) DeletePhotoHandler(ctx *gin.Context) {

	var input deletePhoto.InputDeletePhoto
	input.ID = ctx.Param("photoId")

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

	input.UserId = id

	_, errDeletePhoto := h.service.DeletePhotoService(&input)

	switch errDeletePhoto {

	case "DELETE_PHOTO_NOT_FOUND_404":
		helper.APIResponse(ctx, "Photo data is not exist or deleted", http.StatusForbidden, http.MethodDelete, nil)
		return

	case "DELETE_PHOTO_FAILED_403":
		helper.APIResponse(ctx, "Delete photo data failed", http.StatusForbidden, http.MethodDelete, nil)
		return

	default:
		helper.APIResponse(ctx, "Delete photo data successfully", http.StatusOK, http.MethodDelete, nil)
	}
}
