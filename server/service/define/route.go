package define

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mlinder10/wcj/dictionary"
	"github.com/mlinder10/wcj/llm"
	"github.com/mlinder10/wcj/types"
	"github.com/mlinder10/wcj/utils"
)

func GET_DefineWord(w http.ResponseWriter, r *http.Request) {
	word := mux.Vars(r)["word"]
	if len(word) == 0 {
		utils.WriteError(w, types.Errors.FailedToParseRequest(errors.New("word is required")))
		return
	}

	defs, err := dictionary.Define(word)
	if err == nil {
		utils.WriteJSON(w, 200, types.DefineMeaningsResponse{
			Type:        types.DefinitionResponseTypeDefinitions,
			Definitions: defs,
		})
		return
	}

	if !errors.Is(err, dictionary.InvalidWord{}) {
		utils.WriteError(w, types.Errors.InternalServerError(err))
		return
	}

	// word is invalid, suggest other words
	suggestions, err := dictionary.Suggest(r.Context(), llm.Gemini, word)
	if err != nil {
		utils.WriteError(w, types.Errors.InternalServerError(err))
		return
	}

	utils.WriteJSON(w, 200, types.DefineSuggestionsResponse{
		Type:        types.DefinitionResonseTypeSuggestions,
		Suggestions: suggestions,
	})
}
