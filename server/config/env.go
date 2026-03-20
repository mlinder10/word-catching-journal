package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	DatabaseURL string
	ServerURL   string
	ClientURL   string
	Port        string
	LogPath     string

	JWTKey                 string
	JWTSecret              string
	JWTExpirationInSeconds int

	PostsLimit int
	UsersLimit int

	GeminiAPIKey string

	EmailAddr     string
	EmailAddress  string
	EmailIdentity string
	EmailUsername string
	EmailPassword string
	EmailHost     string
}

var Env config

func init() {
	LoadEnvironment(".env")
}

func LoadEnvironment(from string) {
	godotenv.Load(from)
	Env = config{
		DatabaseURL: getEnv("DATABASE_URL", ""),
		ServerURL:   getEnv("SERVER_URL", ""),
		ClientURL:   getEnv("CLIENT_URL", ""),
		Port:        getEnv("PORT", ""),

		LogPath: getEnv("LOG_PATH", ""),

		JWTKey:                 getEnv("JWT_KEY", ""),
		JWTSecret:              getEnv("JWT_SECRET", ""),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 0),

		PostsLimit: getEnvAsInt("POSTS_LIMIT", 0),
		UsersLimit: getEnvAsInt("USERS_LIMIT", 0),

		GeminiAPIKey: getEnv("GEMINI_API_KEY", ""),

		EmailAddr:     getEnv("EMAIL_ADDR", ""),
		EmailAddress:  getEnv("EMAIL_ADDRESS", ""),
		EmailIdentity: getEnv("EMAIL_IDENTITY", ""),
		EmailPassword: getEnv("EMAIL_PASSWORD", ""),
		EmailHost:     getEnv("EMAIL_HOST", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		b, err := strconv.ParseBool(value)
		if err != nil {
			return fallback
		}
		return b
	}
	return fallback
}
