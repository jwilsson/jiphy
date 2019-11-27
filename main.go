package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	defaultResp := events.APIGatewayProxyResponse{
		Body:       "",
		Headers:    map[string]string{},
		StatusCode: 200,
	}

	log.Printf("Received body: %v", request.Body)

	body, _ := url.ParseQuery(request.Body)
	slackURL, _ := url.QueryUnescape(body.Get("response_url"))
	text := body.Get("text")

	if text == "" {
		return defaultResp, nil
	}

	if text == "list" {
		title := []string{"*All GIFs in Jiphy*"}
		imageKeys := append(title, getImageKeys()...)

		msg := buildSection(
			strings.Join(imageKeys, "\n"),
		)

		sendMessage(slackURL, msg)

		return defaultResp, nil
	}

	image := getImage(text)

	if image == nil {
		imageURL := "https://media.giphy.com/media/l0Iy2hYDgmCjMufzq/giphy-downsized.gif"
		title := fmt.Sprintf("Couldn't find \"%s\"", text)

		msg := buildImage(title, imageURL, "ephemeral")

		sendMessage(slackURL, msg)

		return defaultResp, nil
	}

	title := fmt.Sprintf("%s sent \"%s\"", body.Get("user_name"), text)
	msg := buildImage(title, image.Image, "in_channel")

	sendMessage(slackURL, msg)

	return defaultResp, nil
}

func main() {
	lambda.Start(handleRequest)
}
