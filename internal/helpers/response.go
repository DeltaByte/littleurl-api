package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type ErrorResponseBody struct {
	StatusCode int         `json:"code"`
	Message    string      `json:"message"`
	Details    interface{} `json:"details"`
}

func GatewayErrorResponse(code int, message string) *events.APIGatewayV2HTTPResponse {
	if message == "" {
		message = http.StatusText(code)
	}

	bodyBytes, _ := json.Marshal(ErrorResponseBody{
		StatusCode: code,
		Message:    message,
	})

	return &events.APIGatewayV2HTTPResponse{
		StatusCode: code,
		Body:       string(bodyBytes),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		IsBase64Encoded: false,
	}
}
