package structs

type UserRaw struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserPreAuth struct {
	ID string
	Email string
	PasswordHash string
}

type UserAuthenticated struct {
	ID    string
	Email string
}
type User struct {
	Name         string
	Email        string
	PasswordHash string
}
