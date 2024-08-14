package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJson(request *http.Request, payload any) error {
	if request.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder((request.Body)).Decode(payload)
}

func WriteJson(response http.ResponseWriter, status int, payload any) error {
	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(status)
	return json.NewEncoder(response).Encode(payload)
}

func WriteError(response http.ResponseWriter, status int, err error) {
	WriteJson(response, status, map[string]string{"error": err.Error()})
}
