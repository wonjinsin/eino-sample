package langchain

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/wonjinsin/simple-chatbot/internal/repository"
	"github.com/wonjinsin/simple-chatbot/pkg/errors"
)

// cleanMarkdownJSONParser is a custom parser that cleans markdown code blocks before parsing JSON
type cleanMarkdownJSONParser[T any] struct {
	baseParser schema.MessageParser[T]
}

// Parse cleans markdown code blocks from message content and then parses it
func (p *cleanMarkdownJSONParser[T]) Parse(ctx context.Context, msg *schema.Message) (T, error) {
	var result T
	if msg == nil {
		return result, nil
	}

	// Clean markdown code blocks from content
	content := p.cleanMarkdown(msg.Content)

	// Create a temporary message with cleaned content
	cleanedMsg := &schema.Message{
		Role:      msg.Role,
		Content:   content,
		ToolCalls: msg.ToolCalls,
	}

	// Use base parser to parse the cleaned message
	return p.baseParser.Parse(ctx, cleanedMsg)
}

// cleanMarkdown removes markdown code blocks and extracts JSON
func (p *cleanMarkdownJSONParser[T]) cleanMarkdown(content string) string {
	// Remove markdown code blocks (```json ... ``` or ``` ... ```)
	markdownCodeBlockRegex := regexp.MustCompile("(?s)```(?:json)?\\s*(.*?)\\s*```")
	matches := markdownCodeBlockRegex.FindStringSubmatch(content)
	if len(matches) > 1 {
		// Extract JSON from code block
		return strings.TrimSpace(matches[1])
	}

	// If no code block, try to find JSON object in the content
	// Find the first { and match until the corresponding }
	startIdx := strings.Index(content, "{")
	if startIdx != -1 {
		braceCount := 0
		for i := startIdx; i < len(content); i++ {
			if content[i] == '{' {
				braceCount++
			} else if content[i] == '}' {
				braceCount--
				if braceCount == 0 {
					return strings.TrimSpace(content[startIdx : i+1])
				}
			}
		}
	}

	// Return trimmed content if no JSON found
	return strings.TrimSpace(content)
}

type basicChatRepo struct {
	ollamaLLM model.ToolCallingChatModel
}

// NewBasicChatRepo creates a new basic chat repository
func NewBasicChatRepo(ollamaLLM model.ToolCallingChatModel) repository.BasicChatRepository {
	return &basicChatRepo{ollamaLLM: ollamaLLM}
}

// Ask asks the LLM a question and returns the answer
func (r *basicChatRepo) AskBasicChat(ctx context.Context, _ string) (string, error) {
	messages := []*schema.Message{
		{
			Role:    schema.System,
			Content: "You are a helpful assistant.",
		},
		{
			Role:    schema.User,
			Content: "Please explain about langchain.",
		},
		{
			Role:      schema.Assistant,
			Content:   "LangChain is a library for building language model applications.",
			ToolCalls: nil,
		},
		{
			Role:    schema.User,
			Content: "Please answer the 3 main function.",
		},
	}

	resp, err := r.ollamaLLM.Generate(ctx, messages)
	if err != nil {
		return "", errors.Wrap(err, "failed to generate content")
	}
	return resp.Content, nil
}

// AskBasicPromptTemplateChat asks the LLM a question and returns the answer
func (r *basicChatRepo) AskBasicPromptTemplateChat(ctx context.Context, _ string) (string, error) {
	// Create a prompt template
	template := prompt.FromMessages(
		schema.GoTemplate,
		schema.SystemMessage("You are a JSON-only response assistant. You MUST respond with ONLY valid JSON. The response must be a single JSON object with an 'answer' field containing a plain string value. Do NOT use markdown code blocks, backticks, or any formatting. Do NOT nest JSON objects. Return ONLY the raw JSON object."),
		schema.UserMessage("Generate a report for {{.user}} on {{.date}}. Return your response as a JSON object with this exact structure: {\"answer\": \"your report text here\"}. The answer field must contain a plain string, not nested JSON. Example: {\"answer\": \"John Doe report for 2026-11-09\"}"),
	)

	// Render the template with data
	variables := map[string]any{
		"role": "assistant",
		"user": "John Doe",
		"date": time.Now().AddDate(1, 0, 0).Format("2006-01-02"),
	}

	// Create parser for json with markdown cleaning
	type JSONResponse struct {
		Answer string `json:"answer"`
	}

	// Custom parser that cleans markdown before parsing
	cleanAndParseParser := &cleanMarkdownJSONParser[*JSONResponse]{
		baseParser: schema.NewMessageJSONParser[*JSONResponse](&schema.MessageJSONParseConfig{
			ParseFrom: schema.MessageParseFromContent,
		}),
	}
	parserLambda := compose.MessageParser(cleanAndParseParser)

	chain, err := compose.NewChain[map[string]any, *JSONResponse]().
		AppendChatTemplate(template).
		AppendChatModel(r.ollamaLLM).
		AppendLambda(parserLambda).
		Compile(ctx)

	if err != nil {
		return "", errors.Wrap(err, "failed to compile chain")
	}

	result, err := chain.Invoke(ctx, variables)
	if err != nil {
		return "", errors.Wrap(err, "failed to invoke chain")
	}
	return result.Answer, nil
}
