package main

import (
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	utils "github.com/jwilsson/go-bot-utils"
)

func handleRequest(request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	if !utils.VerifySecret(request, os.Getenv("SLACK_SIGNING_SECRET")) {
		return utils.CreateResponse(403), errors.New("invalid signature header")
	}

	s, err := utils.ParseBody(request)
	if err != nil {
		return utils.CreateResponse(400), err
	}

	images, err := getImages(os.Getenv("DYNAMO_TABLE_NAME"))
	if err != nil {
		return utils.CreateResponse(500), err
	}

	err = sendMessage(MessageInput{
		Command:     s.Command,
		ImageName:   s.Text,
		Images:      images,
		ResponseURL: s.ResponseURL,
		UserName:    s.UserName,
	})

	if err != nil {
		return utils.CreateResponse(500), err
	}

	return utils.CreateResponse(200), err
}

func main() {
	lambda.Start(handleRequest)
}
