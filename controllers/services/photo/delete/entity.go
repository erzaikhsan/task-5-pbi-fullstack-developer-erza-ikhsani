package deletePhoto

type InputDeletePhoto struct {
	ID     string `validate:"required,uuid"`
	UserId string `validate:"required,uuid"`
}
