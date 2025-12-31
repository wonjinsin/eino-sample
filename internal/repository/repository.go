package repository

import (
	"context"
)

// BasicChatRepository defines the interface for basic chat data access
type BasicChatRepository interface {
	AskBasicChat(ctx context.Context, msg string) (string, error)
	AskBasicPromptTemplateChat(ctx context.Context, msg string) (string, error)
	AskBasicParallelChat(ctx context.Context, msg string) (string, error)
	AskBasicBranchChat(ctx context.Context, msg string) (string, error)
	AskWithTool(ctx context.Context, msg string) (string, error)
	AskWithGraph(ctx context.Context, msg string) (string, error)
	AskWithGraphWithBranch(ctx context.Context, _ string) (string, error)
}
