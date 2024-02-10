package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lsendoya/cognitoDalas/internal/user/infrastructure/http"
)

func main() {
	lambda.Start(http.Handler)
}
