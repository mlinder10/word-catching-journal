package types

import "time"

type UserResponse struct {
	ID          string    `json:"id"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	Displayname string    `json:"displayname"`
	Color       string    `json:"color"`
	ImageURL    *string   `json:"imageUrl"`
	CreatedAt   time.Time `json:"createdAt"`
}

type PostResponse struct {
	Posts       []Post `json:"posts"`
	CurrentPage int    `json:"currentPage"`
	TotalPages  int    `json:"totalPages"`
}

type ProfileResponse struct {
	Users       []Profile `json:"users"`
	CurrentPage int       `json:"currentPage"`
	TotalPages  int       `json:"totalPages"`
}

type DefinitionResonseType string

const (
	DefinitionResponseTypeDefinitions DefinitionResonseType = "definitions"
	DefinitionResonseTypeSuggestions  DefinitionResonseType = "suggestions"
)

type DefineMeaningsResponse struct {
	Type        DefinitionResonseType `json:"type"`
	Definitions []Definition          `json:"definitions"`
}

type DefineSuggestionsResponse struct {
	Type        DefinitionResonseType `json:"type"`
	Suggestions []string              `json:"suggestions"`
}
