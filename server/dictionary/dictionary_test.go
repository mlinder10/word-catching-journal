package dictionary

import (
	"context"
	"errors"
	"testing"

	"github.com/mlinder10/wcj/llm"
)

func TestDefine(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		_, err := Define("test")

		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("invalid word", func(t *testing.T) {
		_, err := Define("asdf")

		if err == nil {
			t.Fatal("expected error")
		}

		if ok := errors.Is(err, InvalidWord{}); !ok {
			t.Fatal("expected InvalidWord")
		}
	})
}

func TestSuggest(t *testing.T) {
	t.Run("", func(t *testing.T) {
		words, err := Suggest(context.Background(), llm.Gemini, "vociferus")

		if err != nil {
			t.Fatal(err)
		}

		if len(words) == 0 {
			t.Fatal("expected > 0")
		}

		t.Log(words)
	})
}
