package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jwilsson/go-bot-utils"
)

func handleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	s, err := utils.ParseBody(request.Body)
	if err != nil {
		return utils.CreateResponse(500), err
	}

	if s.Text == "" {
		return utils.CreateResponse(200), nil
	}

	if s.Text == "list" {
		images, err := getImages(os.Getenv("DYNAMO_TABLE_NAME"))
		if err != nil {
			return utils.CreateResponse(500), err
		}

		msg := createList(images)

		utils.SendMessage(s.ResponseURL, msg)

		return utils.CreateResponse(200), nil
	}

	image, err := getImage(s.Text, os.Getenv("DYNAMO_TABLE_NAME"))
	if err != nil {
		return utils.CreateResponse(500), err
	}

	if image == nil {
		image = &Image{
			GiphyURL:  "https://giphy.com/gifs/stonehampress-funny-horse-l0Iy2hYDgmCjMufzq",
			ImageName: "gif",
			ImageURL:  "https://media.giphy.com/media/l0Iy2hYDgmCjMufzq/giphy-downsized.gif",
		}

		msg := createImage(image, s.UserName, "ephemeral")

		utils.SendMessage(s.ResponseURL, msg)

		return utils.CreateResponse(200), nil
	}

	msg := createImage(image, s.UserName, "in_channel")

	utils.SendMessage(s.ResponseURL, msg)

	return utils.CreateResponse(200), nil
}

func main() {
	lambda.Start(handleRequest)
}
