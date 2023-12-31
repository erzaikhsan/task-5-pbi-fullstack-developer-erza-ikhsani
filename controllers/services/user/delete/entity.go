package deleteUser

type InputDeleteUser struct {
	ID string `validate:"required,uuid"`
}
