package main

// Things to DO
// go get github.com/aws/aws-lambda-go/lambda
// then after coding we need to compile our file in linux binary with command "GOOD=linux go build -o main"
// then zip this binary and upload to lambda
// and dont forget to add main fucntion in handler

// TODO: apply CI/CD for build as well
import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID float64 `json:"id"`
	Value string `json:"value"`
}

type Response struct {
	Message string `json:"message"`
	Ok bool `json:"ok"`
}

func Handler(request Request) (Response, error)  {
	return Response{
		Message: fmt.Sprintf("Process Request ID %f", request.ID)
		Ok: true
	}, nil
}

func main()  {
	lambda.start(Handler)
}