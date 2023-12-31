package helper

import (
	goValidator "github.com/erzaikhsan/task-5-pbi-btpns-erza-ikhsani/helpers/goValidator"
	"github.com/go-playground/validator/v10"
)

func GoValidator(s interface{}, config []goValidator.ErrorMetaConfig) (interface{}, int) {
	var validate *validator.Validate
	validators := goValidator.NewValidator(validate)
	bind := goValidator.NewBindValidator(validators)

	errResponse, errCount := bind.BindValidator(s, config)
	return errResponse, errCount
}
