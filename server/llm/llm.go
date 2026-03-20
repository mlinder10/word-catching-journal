package llm

import "context"

type LLM interface {
	Prompt(ctx context.Context, prompt string) (Output, error)
}

type Output struct {
	Text             string
	PromptTokens     int
	CompletionTokens int
}
