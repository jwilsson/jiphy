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

	s, err := utils.ParseBody(request.Body)
	if err != nil {
		return utils.CreateResponse(400), err
	}

	if s.Text == "" {
		return utils.CreateResponse(200), nil
	}

	images, err := getImages(os.Getenv("DYNAMO_TABLE_NAME"))
	if err != nil {
		return utils.CreateResponse(500), err
	}

	if s.Text == "list" {
		utils.SendMessage(s.ResponseURL, createListMessage(images))
	} else {
		responseType := "in_channel"
		image := findImage(images, s.Text)

		if image == nil {
			responseType = "ephemeral"
			image = &Image{
				GiphyURL:  "https://giphy.com/gifs/stonehampress-funny-horse-l0Iy2hYDgmCjMufzq",
				ImageName: "gif",
				ImageURL:  "https://media.giphy.com/media/l0Iy2hYDgmCjMufzq/giphy-downsized.gif",
			}
		}

		utils.SendMessage(s.ResponseURL, createImageMessage(image, s.UserName, s.Command, responseType))
	}

	return utils.CreateResponse(200), nil
}

func main() {
	lambda.Start(handleRequest)
}
