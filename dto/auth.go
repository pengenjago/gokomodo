package dto

type LoginRes struct {
	AccessToken string `json:"accessToken"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	LoginAs  string
}

type UserAuth struct {
	Role  string
	Id    string
	Email string
}
