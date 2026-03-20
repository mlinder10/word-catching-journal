package types

import (
	"time"
)

type Visibility string

const (
	Private Visibility = "private"
	Public  Visibility = "public"
)

func (v *Visibility) String() *string {
	return (*string)(v)
}

type PostMapable interface {
	ToPost() Post
}

type ProfileMapable interface {
	ToProfile() Profile
}

type Post struct {
	ID            string     `json:"id"`
	Visibility    Visibility `json:"visibility"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	LikeCount     int        `json:"likeCount"`
	BookmarkCount int        `json:"bookmarkCount"`
	IsLiked       bool       `json:"isLiked"`
	IsBookmarked  bool       `json:"isBookmarked"`

	Word          string   `json:"word"`
	Definition    string   `json:"definition"`
	PartOfSpeech  string   `json:"partOfSpeech"`
	Pronunciation string   `json:"pronunciation"`
	Synonyms      []string `json:"synonyms"`
	Antonyms      []string `json:"antonyms"`
	Example       string   `json:"example"`

	UserID      string  `json:"userId"`
	Username    string  `json:"username"`
	Displayname string  `json:"displayname"`
	ImageURL    *string `json:"imageUrl"`
	Color       string  `json:"color"`
}

type Profile struct {
	ID          string  `json:"id"`
	Username    string  `json:"username"`
	Displayname string  `json:"displayname"`
	ImageUrl    *string `json:"imageUrl"`
	Color       string  `json:"color"`
}

type Definition struct {
	Word          string   `json:"word"`
	Definition    string   `json:"definition"`
	PartOfSpeech  string   `json:"partOfSpeech"`
	Pronunciation string   `json:"pronunciation"`
	Synonyms      []string `json:"synonyms"`
	Antonyms      []string `json:"antonyms"`
	Example       string   `json:"example"`
}
