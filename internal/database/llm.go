package database

import (
	"github.com/tmc/langchaingo/llms/ollama"
)

func NewOllamaLLM() (*ollama.LLM, error) {
	return ollama.New(
		ollama.WithModel("gemma3:1b"),
	)
}
