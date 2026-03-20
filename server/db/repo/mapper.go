package repo

import (
	"encoding/json"
	"time"

	"github.com/mlinder10/wcj/types"
)

// post

func newPost(
	id, visibility, userID, word, definition,
	partOfSpeech, pronunciation, example, username,
	displayName, color string, createdAt, updatedAt time.Time,
	likeCount, bookmarkCount int, isLiked, isBookmarked bool,
	synonyms, antonyms []string, imageUrl *string) types.Post {
	return types.Post{
		ID:            id,
		Visibility:    types.Visibility(visibility),
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		LikeCount:     likeCount,
		BookmarkCount: bookmarkCount,
		IsLiked:       isLiked,
		IsBookmarked:  isBookmarked,

		Word:          word,
		Definition:    definition,
		PartOfSpeech:  partOfSpeech,
		Pronunciation: pronunciation,
		Synonyms:      synonyms,
		Antonyms:      antonyms,
		Example:       example,

		UserID:      userID,
		Username:    username,
		Displayname: displayName,
		ImageURL:    imageUrl,
		Color:       color,
	}
}

func (row GetRecentPostsRow) ToPost() types.Post {
	var synonyms []string
	if err := json.Unmarshal([]byte(row.Synonyms), &synonyms); err != nil {
		synonyms = []string{}
	}
	var antonyms []string
	if err := json.Unmarshal([]byte(row.Antonyms), &antonyms); err != nil {
		antonyms = []string{}
	}
	return newPost(
		row.ID,
		row.Visibility,
		row.UserID,
		row.Word,
		row.Definition,
		row.PartOfSpeech,
		row.Pronunciation,
		row.Example,
		row.Username,
		row.Displayname,
		row.Color,
		row.CreatedAt,
		row.UpdatedAt,
		int(row.LikeCount),
		int(row.BookmarkCount),
		row.IsLiked == 1,
		row.IsBookmarked == 1,
		synonyms,
		antonyms,
		row.ImageUrl,
	)
}

func (row GetFollowingPostsRow) ToPost() types.Post {
	var synonyms []string
	if err := json.Unmarshal([]byte(row.Synonyms), &synonyms); err != nil {
		synonyms = []string{}
	}
	var antonyms []string
	if err := json.Unmarshal([]byte(row.Antonyms), &antonyms); err != nil {
		antonyms = []string{}
	}
	return newPost(
		row.ID,
		row.Visibility,
		row.UserID,
		row.Word,
		row.Definition,
		row.PartOfSpeech,
		row.Pronunciation,
		row.Example,
		row.Username,
		row.Displayname,
		row.Color,
		row.CreatedAt,
		row.UpdatedAt,
		int(row.LikeCount),
		int(row.BookmarkCount),
		row.IsLiked == 1,
		row.IsBookmarked == 1,
		synonyms,
		antonyms,
		row.ImageUrl,
	)
}

func (row GetProfilePostsRow) ToPost() types.Post {
	var synonyms []string
	if err := json.Unmarshal([]byte(row.Synonyms), &synonyms); err != nil {
		synonyms = []string{}
	}
	var antonyms []string
	if err := json.Unmarshal([]byte(row.Antonyms), &antonyms); err != nil {
		antonyms = []string{}
	}
	return newPost(
		row.ID,
		row.Visibility,
		row.UserID,
		row.Word,
		row.Definition,
		row.PartOfSpeech,
		row.Pronunciation,
		row.Example,
		row.Username,
		row.Displayname,
		row.Color,
		row.CreatedAt,
		row.UpdatedAt,
		int(row.LikeCount),
		int(row.BookmarkCount),
		row.IsLiked == 1,
		row.IsBookmarked == 1,
		synonyms,
		antonyms,
		row.ImageUrl,
	)
}

// profile

func newProfile(id, username, displayname, color string, imageUrl *string) types.Profile {
	return types.Profile{
		ID:          id,
		Username:    username,
		Displayname: displayname,
		ImageUrl:    imageUrl,
		Color:       color,
	}
}

func (row GetProfileRow) ToProfile() types.Profile {
	return newProfile(row.ID, row.Username, row.Displayname, row.Color, row.ImageUrl)
}

func (row GetFollowingRow) ToProfile() types.Profile {
	return newProfile(row.ID, row.Username, row.Displayname, row.Color, row.ImageUrl)
}

func (row GetFollowersRow) ToProfile() types.Profile {
	return newProfile(row.ID, row.Username, row.Displayname, row.Color, row.ImageUrl)
}
