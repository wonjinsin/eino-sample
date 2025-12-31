package usecase

import (
	"context"
)

// BasicChatService defines the interface for basic chat business logic
type BasicChatService interface {
	AskBasicChat(ctx context.Context, msg string) (string, error)
	AskBasicPromptTemplateChat(ctx context.Context, msg string) (string, error)
}
