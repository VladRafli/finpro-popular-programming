package dto

type ReadLoginDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type CreateUserDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type ForgotPasswordDto struct {
	Email string `json:"email" validate:"required,email"`
}

type ChangePasswordDto struct {
	OldPassword string `json:"old_password" validate:"required"`
	NewPassword string `json:"new_password" validate:"required"`
}