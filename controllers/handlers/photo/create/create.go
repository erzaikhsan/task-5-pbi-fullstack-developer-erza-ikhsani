package handlerCreatePhoto

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/dgrijalva/jwt-go"
	createPhoto "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/controllers/services/photo/create"
	helper "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/helpers"
	goValidator "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/helpers/goValidator"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service createPhoto.Service
}

func NewHandlerCreatePhoto(service createPhoto.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreatePhotoHandler(ctx *gin.Context) {

	var input createPhoto.InputCreatePhoto
	ctx.ShouldBind(&input)

	config := goValidator.ErrorConfig{
		Options: []goValidator.ErrorMetaConfig{
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
		helper.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
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

	_, errCreatePhoto := h.service.CreatePhotoService(&input)

	switch errCreatePhoto {

	case "CREATE_PHOTO_CONFLICT_409":
		helper.APIResponse(ctx, "Photo already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "CREATE_PHOTO_FAILED_403":
		helper.APIResponse(ctx, "Create new photo account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		helper.APIResponse(ctx, "Create new photo account successfully", http.StatusCreated, http.MethodPost, nil)
	}
}
