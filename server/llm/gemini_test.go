package llm

import (
	"context"
	"testing"

	"github.com/mlinder10/wcj/config"
)

func TestGeminiPrompt(t *testing.T) {
	config.LoadEnvironment("../.env")
	result, err := Gemini.Prompt(context.Background(), "Hello")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(result.Text)
}
