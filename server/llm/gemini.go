package llm

import (
	"context"

	"github.com/mlinder10/wcj/config"
	"google.golang.org/genai"
)

type gemini struct {
	client *genai.Client
}

func newGemini() gemini {
	client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey: config.Env.GeminiAPIKey,
	})
	if err != nil {
		panic(err)
	}
	return gemini{client}
}

func (g gemini) Prompt(ctx context.Context, prompt string) (Output, error) {
	result, err := g.client.Models.GenerateContent(
		ctx,
		"gemini-3-flash-preview",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return Output{}, err
	}

	output := Output{
		Text:             result.Text(),
		PromptTokens:     int(result.UsageMetadata.PromptTokenCount),
		CompletionTokens: int(result.UsageMetadata.CandidatesTokenCount),
	}

	return output, nil
}

var Gemini = newGemini()
