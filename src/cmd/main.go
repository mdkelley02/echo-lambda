package main

import (
	"fmt"

	"github.com/echo-lambda/pkg/handlers"

	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	startTime := time.Now()
	fmt.Printf("Starting Echo Request at %s", startTime.Format(time.RFC3339))
	lambda.Start(handlers.EchoRequestHandler)
}
