package main

import (
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	StatusCode int
	Body       string
}

func Handler() (Response, error) {
	return Response{
		Body:       "Okay so your other function also executed successfully!",
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
