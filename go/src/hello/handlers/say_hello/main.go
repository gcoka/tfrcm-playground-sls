package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is function
type Response struct {
	Message string `json:"message"`
}

// Handler is function
func Handler() (Response, error) {
	return Response{
		Message: "Hello from Gemcook!",
	}, nil
}

func main() {
	lambda.Start(Handler)
}
