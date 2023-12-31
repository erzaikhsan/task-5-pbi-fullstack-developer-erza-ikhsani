package updateUser

import (
	"time"
)

type InputUpdateUser struct {
	ID        string `validate:"required,uuid"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Password  string `json:"password" validate:"required"`
	UpdatedAt time.Time
}
