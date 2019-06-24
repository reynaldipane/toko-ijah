package helper

import (
	"encoding/json"
	"net/http"
)

const (
	//ContentJSON will tell that the response that the application send is in application/json
	ContentJSON = "application/json"
	//ContentText will tell that the response that the application send is in application/text
	ContentText = "application/text"
)

// BuildResponse will build response for user's request
// like setting up the header, status_code, and body
func BuildResponse(
	w http.ResponseWriter,
	contentType string,
	statusCode int,
	body string,
) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	w.Write([]byte(body))
}

// BuildResponseWithError will be used if there is error happened
// from user's request, so we will tell them the appropiate
// error message that we will return from this function
func BuildResponseWithError(
	w http.ResponseWriter,
	contentType string,
	statusCode int,
	message string,
) {
	errorMessage, _ := json.Marshal(map[string]string{
		"error_message": message,
	})

	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	w.Write([]byte(errorMessage))
}
