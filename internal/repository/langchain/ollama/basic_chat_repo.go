package langchain

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/wonjinsin/simple-chatbot/internal/repository"
)

type basicChatRepo struct {
	ollamaLLM *ollama.LLM
}

// NewBasicChatRepo creates a new basic chat repository
func NewBasicChatRepo(ollamaLLM *ollama.LLM) repository.BasicChatRepository {
	return &basicChatRepo{ollamaLLM: ollamaLLM}
}

// Ask asks the LLM a question and returns the answer
func (r *basicChatRepo) Ask(ctx context.Context, _ string) (string, error) {
	messages := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "You are a helpful assistant."),
		llms.TextParts(llms.ChatMessageTypeHuman, "Please explain about langchain."),
		llms.TextParts(llms.ChatMessageTypeAI, "LangChain is a library for building language model applications."),
		llms.TextParts(llms.ChatMessageTypeHuman, "Please answer the 3 main function."),
	}

	resp, err := r.ollamaLLM.GenerateContent(ctx, messages)
	if err != nil {
		return "", err
	}
	var answer string
	for i, choice := range resp.Choices {
		answer += fmt.Sprintf("%d answer: %s\n", i+1, choice.Content)
	}
	return answer, nil
}
