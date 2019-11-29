package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func createResponse(statusCode int) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		Body:       "",
		Headers:    map[string]string{},
		StatusCode: statusCode,
	}
}

func handleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received body: %v", request.Body)

	body, err := url.ParseQuery(request.Body)
	if err != nil {
		return createResponse(500), err
	}

	slackURL, err := url.QueryUnescape(body.Get("response_url"))
	if err != nil {
		return createResponse(500), err
	}

	text := body.Get("text")

	if text == "" {
		return createResponse(200), nil
	}

	if text == "list" {
		images, err := getImages(os.Getenv("DYNAMO_TABLE_NAME"))
		if err != nil {
			return createResponse(500), err
		}

		msg := createList(images)

		sendMessage(slackURL, msg)

		return createResponse(200), nil
	}

	image, err := getImage(text, os.Getenv("DYNAMO_TABLE_NAME"))
	if err != nil {
		return createResponse(500), err
	}

	if image == nil {
		imageURL := "https://media.giphy.com/media/l0Iy2hYDgmCjMufzq/giphy-downsized.gif"
		title := fmt.Sprintf("Couldn't find \"%s\"", text)

		msg := createImage(title, imageURL, "ephemeral")

		sendMessage(slackURL, msg)

		return createResponse(200), nil
	}

	title := fmt.Sprintf("%s sent \"%s\"", body.Get("user_name"), text)
	msg := createImage(title, image.ImageURL, "in_channel")

	sendMessage(slackURL, msg)

	return createResponse(200), nil
}

func main() {
	lambda.Start(handleRequest)
}
