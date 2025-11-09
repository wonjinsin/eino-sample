package database

import (
	"context"
	"time"

	"github.com/cloudwego/eino-ext/components/model/ollama"
)

func NewOllamaLLM() (*ollama.ChatModel, error) {
	ctx := context.Background()
	return ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		// Basic Configuration
		BaseURL: "http://localhost:11434", // Ollama service address
		Timeout: 30 * time.Second,         // Request timeout

		// Model Configuration
		Model: "gemma3:1b", // Model name
	})
}
