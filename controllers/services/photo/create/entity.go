package createPhoto

type InputCreatePhoto struct {
	Title    string `form:"title" json:"title" validate:"required"`
	Caption  string `form:"caption" json:"caption" validate:"required"`
	PhotoUrl string `json:"photo_url" validate:"required"`
	UserId   string `json:"user_id" validate:"required"`
}
