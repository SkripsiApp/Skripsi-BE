package request

type AdminRegister struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}