package resultPhoto

type InputResultPhoto struct {
	UserId string `validate:"required,uuid"`
}
