package updatePhoto

import (
	"time"
)

type InputUpdatePhoto struct {
	ID        string `validate:"required,uuid"`
	Title     string `form:"title" json:"title" validate:"required"`
	Caption   string `form:"caption" json:"caption" validate:"required"`
	PhotoUrl  string `json:"photo_url" validate:"required"`
	UserId    string `validate:"required,uuid"`
	UpdatedAt time.Time
}
