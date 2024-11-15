package request

type UserRegister struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UserUpdate struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSendOTP struct {
	Email string `json:"email"`
}

type UserVerifyOTP struct {
	Email string `json:"email"`
	Otp   string `json:"otp"`
}

type UserNewPassword struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}