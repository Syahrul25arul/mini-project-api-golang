package helper

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func BcryptPassword(passwordSalt string) string {
	newPassword, _ := bcrypt.GenerateFromPassword([]byte(passwordSalt), 8)
	return string(newPassword)
}

func ShowMessage(httpStatus int, message string, w http.ResponseWriter) {
	w.WriteHeader(httpStatus)
	fmt.Fprintf(w, message)
}

func ShowDataWithTypeJson(data interface{}, w http.ResponseWriter, code int) {
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func ShowDataWithTypeXml(data interface{}, w http.ResponseWriter, code int) {
	w.Header().Add("Content-type", "application/xml")
	xml.NewEncoder(w).Encode(data)
}
