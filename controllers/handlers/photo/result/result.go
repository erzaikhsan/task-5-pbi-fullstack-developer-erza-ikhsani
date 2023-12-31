package handlerResultPhoto

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	resultPhoto "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/services/photo/result"
	helper "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/helpers"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service resultPhoto.Service
}

func NewHandlerResultPhoto(service resultPhoto.Service) *handler {
	return &handler{service: service}
}

func (h *handler) ResultPhotoHandler(ctx *gin.Context) {
	var input resultPhoto.InputResultPhoto

	token, _ := ctx.Get("user")
	claims, ok := token.(jwt.MapClaims)
	if !ok {
		helper.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, "Token verification failed")
		return
	}
	id, idExists := claims["id"].(string)
	if !idExists {
		helper.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, "Id not found")
		return
	}

	input.UserId = id

	resultPhoto, errResultPhoto := h.service.ResultPhotoService(&input)

	switch errResultPhoto {

	case "RESULT_PHOTO_NOT_FOUND_404":
		helper.APIResponse(ctx, "Photo data is not exists", http.StatusConflict, http.MethodPost, nil)

	default:
		helper.APIResponse(ctx, "Results photo data successfully", http.StatusOK, http.MethodPost, resultPhoto)
	}
}
