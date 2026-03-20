package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/mlinder10/wcj/types"
	"github.com/mlinder10/wcj/utils/logger"
)

func ParsePageQuery(r *http.Request) int {
	page := r.URL.Query().Get("page")
	if page == "" {
		return 0
	}
	i, err := strconv.Atoi(page)
	if err != nil {
		return 0
	}
	if i < 1 {
		return 0
	}
	return i - 1
}

func ParseJSON[T any](r *http.Request, payload T) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	// TODO: uncomment in production
	// return json.NewDecoder(r.Body).Decode(payload)

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, payload)
	if err != nil {
		fmt.Println("body: " + string(bytes))
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, payload any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(payload)
}

func WriteGzip(w http.ResponseWriter, status int, payload []byte) error {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Content-Encoding", "gzip")
	w.WriteHeader(status)
	_, err := w.Write(payload)
	return err
}

func WriteError(w http.ResponseWriter, err types.APIError, internal ...error) {
	fmt.Println(err.WrappedError.Error())
	logger.Write(err.WrappedError.Error())
	for _, e := range internal {
		fmt.Println(e)
		logger.Write(e.Error())
	}
	WriteJSON(w, err.Code, map[string]string{"error": err.Message})
}

func RandomColor() string {
	return fmt.Sprintf("#%06x", rand.Intn(0xFFFFFF))
}

const codeChars = "abcdefghijklmnopqrstuvwxyz0123456789"

func ConfirmationCode() string {
	code := make([]byte, 6)
	for i := range code {
		code[i] = codeChars[rand.Intn(len(codeChars))]
	}
	return string(code)
}
