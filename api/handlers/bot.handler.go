package handlers

import (
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/internal/ApiResponse"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/internal/config"
	"github.com/victormazeli/klusterthon-ai-symptom-diagnosis-backend/internal/dto"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BotHandler struct {
	Db  *gorm.DB
	Cfg *config.Config
}

func (b *BotHandler) HandleChat(c *gin.Context) {
	var chatInput dto.ChatInputDTO
	if err := c.ShouldBindJSON(&chatInput); err != nil {
		ApiResponse.SendBadRequest(c, "Invalid JSON input: "+err.Error())
		return
	}

	// Create a channel to communicate the result and the modified request messages back from the Goroutine
	resultChan := make(chan struct {
		Result   string
		Messages []openai.ChatCompletionMessage
	})

	// Use a Goroutine to make the OpenAI API call concurrently
	go func() {
		client := openai.NewClient(b.Cfg.OpenAI.Key)
		req := openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "you are a helpful chatbot",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: chatInput.Message,
				},
			},
		}

		resp, err := client.CreateChatCompletion(c, req)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			resultChan <- struct {
				Result   string
				Messages []openai.ChatCompletionMessage
			}{Result: "Internal server error", Messages: nil}
			return
		}

		// Append the AI's response to the request messages
		req.Messages = append(req.Messages, resp.Choices[0].Message)

		resultChan <- struct {
			Result   string
			Messages []openai.ChatCompletionMessage
		}{Result: resp.Choices[0].Message.Content, Messages: req.Messages}
	}()

	// Here, you can continue processing other tasks while waiting for the Goroutine to complete

	// Wait for the Goroutine to finish and retrieve the result and modified request messages
	result := <-resultChan

	// Send the result to the client
	ApiResponse.SendSuccess(c, "Response successful", result.Result)
}

func (b *BotHandler) GetConversations(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	// Fetch user from database by userID
	c.JSON(http.StatusOK, gin.H{"user_id": userID, "name": "John Doe"})
}
