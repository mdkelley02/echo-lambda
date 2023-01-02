package handlers

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
)

var (
	ErrMissingBody   = errors.New("missing body")
	ErrMissingEcho   = errors.New("missing echo")
	ErrMalformedBody = errors.New("malformed body")
)

type EchoRequest struct {
	Echo string `json:"echo"`
}

type EchoResponse struct {
	Time string `json:"time"`
	Echo string `json:"echo"`
}

type EchoErrorResponse struct {
	Error string `json:"error"`
}

func makeResponse(code int, body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       body,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

func makeErrorResponse(err error) events.APIGatewayProxyResponse {
	m, err := json.Marshal(EchoErrorResponse{Error: err.Error()})
	if err != nil {
		return events.APIGatewayProxyResponse{}
	}
	return makeResponse(
		500,
		string(m),
	)
}
