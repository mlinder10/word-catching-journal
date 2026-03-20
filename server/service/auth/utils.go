package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mlinder10/wcj/types"
	"github.com/mlinder10/wcj/utils"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func verifyPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func permissionDenied(w http.ResponseWriter, err ...error) {
	for _, e := range err {
		fmt.Println(e.Error())
	}
	utils.WriteError(w, types.Errors.Unauthenticated(errors.New("unauthenticated")))
}
