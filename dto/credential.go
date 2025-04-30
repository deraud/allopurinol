package dto

type AuthRegisterRequest struct {
	Name     string `json:"name" binding:"required,max=255"`
	Email    string `json:"email" binding:"required,email,max=255"`
	Password string `json:"password" binding:"required,password"`
}

type AuthLoginRequest struct {
	Email    string `json:"email" binding:"required,email,max=255"`
	Password string `json:"password" binding:"required,max=255"`
}

type AuthForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email,max=255"`
}

type AuthResetPasswordRequest struct {
	Token       string `json:"token" binding:"required,max=255"`
	NewPassword string `json:"password" binding:"required,password"`
}

type AuthVerifyTokenRequest struct {
	Token string `json:"token" binding:"required,max=255"`
}

type AuthDtoChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required,max=255"`
	NewPassword string `json:"new_password" binding:"required,password"`
}
