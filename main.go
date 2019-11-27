package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

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
		title := []string{"*All GIFs in Jiphy*"}
		imageKeys := append(title, getImageKeys()...)

		msg := buildSection(
			strings.Join(imageKeys, "\n"),
		)

		sendMessage(slackURL, msg)

		return createResponse(200), nil
	}

	image := getImage(text)

	if image == nil {
		imageURL := "https://media.giphy.com/media/l0Iy2hYDgmCjMufzq/giphy-downsized.gif"
		title := fmt.Sprintf("Couldn't find \"%s\"", text)

		msg := buildImage(title, imageURL, "ephemeral")

		sendMessage(slackURL, msg)

		return createResponse(200), nil
	}

	title := fmt.Sprintf("%s sent \"%s\"", body.Get("user_name"), text)
	msg := buildImage(title, image.Image, "in_channel")

	sendMessage(slackURL, msg)

	return createResponse(200), nil
}

func main() {
	lambda.Start(handleRequest)
}
