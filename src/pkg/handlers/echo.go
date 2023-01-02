package handlers

import (
	"log"
	"time"

	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func EchoRequestHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	if request.Body == "" {
		return makeErrorResponse(ErrMissingBody), nil
	}

	var req EchoRequest
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return makeErrorResponse(ErrMalformedBody), nil
	}

	if req.Echo == "" {
		return makeErrorResponse(ErrMissingEcho), nil
	}

	res := EchoResponse{
		Time: time.Now().UTC().Format("2006-01-02 15:04:05"),
		Echo: req.Echo,
	}
	b, err := json.Marshal(res)
	if err != nil {
		return makeErrorResponse(err), nil
	}

	return makeResponse(200, string(b)), nil
}
