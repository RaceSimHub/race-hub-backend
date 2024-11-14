package request

type PostUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type PostUserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
