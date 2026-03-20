package dictionary

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/mlinder10/wcj/llm"
	"github.com/mlinder10/wcj/types"
)

const url = "https://api.dictionaryapi.dev/api/v2/entries/en/"

type InvalidWord struct{}

func (e InvalidWord) Error() string {
	return "No definitions found"
}

type successResponse struct {
	Word     string `json:"word"`
	Phonetic string `json:"phonetic"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string   `json:"definition"`
			Synonyms   []string `json:"synonyms"`
			Antonyms   []string `json:"antonyms"`
		} `json:"definitions"`
		Synonyms []string `json:"synonyms"`
		Antonyms []string `json:"antonyms"`
	} `json:"meanings"`
}

func (s successResponse) toDefinitions() []types.Definition {
	result := []types.Definition{}
	for _, m := range s.Meanings {
		for _, d := range m.Definitions {
			synonyms := m.Synonyms
			if len(d.Synonyms) > 0 {
				synonyms = d.Synonyms
			}
			antonyms := m.Antonyms
			if len(d.Antonyms) > 0 {
				antonyms = d.Antonyms
			}
			result = append(result, types.Definition{
				Word:          s.Word,
				Definition:    d.Definition,
				PartOfSpeech:  m.PartOfSpeech,
				Pronunciation: s.Phonetic,
				Synonyms:      synonyms,
				Antonyms:      antonyms,
			})
		}
	}
	return result
}

func Define(word string) ([]types.Definition, error) {
	res, err := http.Get(url + word)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, InvalidWord{}
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var successBody []successResponse
	err = json.Unmarshal(body, &successBody)
	if err != nil {
		return nil, err
	}

	result := []types.Definition{}
	for _, r := range successBody {
		defs := r.toDefinitions()
		result = append(result, defs...)
	}

	return result, nil
}

func suggestWordsPrompt(word string) string {
	return fmt.Sprintf(`Act as a dictionary API. You will be given a misspelled word. Return a list of words that are similar to the misspelled word.
Word: %s
Output: return a JSON array of strings (ex. ["word1", "word2", "word3"])`, word)
}

func Suggest(ctx context.Context, llm llm.LLM, word string) ([]string, error) {
	output, err := llm.Prompt(ctx, suggestWordsPrompt(word))
	if err != nil {
		return nil, err
	}

	var suggestions []string
	if err = json.Unmarshal([]byte(output.Text), &suggestions); err != nil {
		return nil, err
	}

	return suggestions, nil
}
