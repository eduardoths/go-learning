package migrations

type Users struct {
	Base
	Name         string
	Email        string
	PasswordHash string
}
