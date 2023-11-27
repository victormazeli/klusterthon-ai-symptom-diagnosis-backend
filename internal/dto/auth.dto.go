package dto

type LoginDTO struct {
	Email    string `json:"email" mod:"trim" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type RegisterDTO struct {
	Email    string `json:"email" mod:"trim" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	//FullName string `json:"full_name" mod:"trim" binding:"required"`
}
