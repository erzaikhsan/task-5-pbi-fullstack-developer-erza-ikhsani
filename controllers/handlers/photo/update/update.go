package handlerUpdatePhoto

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/dgrijalva/jwt-go"
	updatePhoto "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/services/photo/update"
	helper "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/helpers"
	goValidator "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/helpers/goValidator"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service updatePhoto.Service
}

func NewHandlerUpdatePhoto(service updatePhoto.Service) *handler {
	return &handler{service: service}
}

func (h *handler) UpdatePhotoHandler(ctx *gin.Context) {

	var input updatePhoto.InputUpdatePhoto
	input.ID = ctx.Param("photoId")
	ctx.ShouldBind(&input)

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
				Field:   "Title",
				Message: "title is required on body",
			},
			goValidator.ErrorMetaConfig{
				Tag:     "required",
				Field:   "Caption",
				Message: "caption is required on body",
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

	input.UserId = id

	_, fileHeader, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": "invalid request body " + err.Error(),
		})
		return
	}

	allowedExtensions := []string{".png", ".jpg", ".jpeg", ".webp"}
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	allowed := false
	for _, allowedExt := range allowedExtensions {
		if ext == allowedExt {
			allowed = true
			break
		}
	}
	if !allowed {
		helper.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodGet, errResponse)
		return
	}

	uploadPath := "public/photos/" + fileHeader.Filename
	err = ctx.SaveUploadedFile(fileHeader, uploadPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failed",
			"message": "failed to save file " + err.Error(),
		})
		return
	}

	input.PhotoUrl = uploadPath

	_, errUpdatePhoto := h.service.UpdatePhotoService(&input)

	switch errUpdatePhoto {

	case "UPDATE_PHOTO_NOT_FOUND_404":
		helper.APIResponse(ctx, "Photo data is not exist or deleted", http.StatusNotFound, http.MethodPost, nil)

	case "UPDATE_PHOTO_FAILED_403":
		helper.APIResponse(ctx, "Update photo data failed", http.StatusForbidden, http.MethodPost, nil)

	default:
		helper.APIResponse(ctx, "Update photo data sucessfully", http.StatusOK, http.MethodPost, nil)
	}
}
