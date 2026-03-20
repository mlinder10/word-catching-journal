package types

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	Displayname string `json:"displayname"`
	Password    string `json:"password"`
}

type CreatePostRequest struct {
	Content    content `json:"content"`
	Visibility string  `json:"visibility"`
}

type content struct {
	Word          string   `json:"word"`
	Definition    string   `json:"definition"`
	PartOfSpeech  string   `json:"partOfSpeech"`
	Pronunciation string   `json:"pronunciation"`
	Synonyms      []string `json:"synonyms"`
	Antonyms      []string `json:"antonyms"`
	Example       string   `json:"example"`
}

type VerifyEmailRequest struct {
	Code string `json:"code"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type ResetPasswordRequest struct {
	Email string `json:"email"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password"`
}
