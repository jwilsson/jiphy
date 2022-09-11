package main

import (
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	utils "github.com/jwilsson/go-bot-utils"
	"golang.org/x/exp/slices"
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

	if s.Text == "list" {
		utils.SendMessage(s.ResponseURL, createListMessage(images))
	} else {
		i := slices.IndexFunc(images, func(img Image) bool {
			return img.ImageName == s.Text
		})

		var responseType string
		var image *Image

		if i >= 0 {
			responseType = "in_channel"
			image = &images[i]
		} else {
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
