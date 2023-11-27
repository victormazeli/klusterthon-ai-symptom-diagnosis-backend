package dto

type ChatInputDTO struct {
	Message string `json:"message" mod:"trim" binding:"required"`
}
