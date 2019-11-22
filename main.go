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
		imageKeys := getImageKeys()

		imageKeys = append(title, imageKeys...)

		sectionText := strings.Join(imageKeys, "\n")
		blocks := buildSection(sectionText)

		sendMessage(slackURL, "ephemeral", blocks)

		return defaultResp, nil
	}

	image := getImage(text)

	if image == nil {
		title := fmt.Sprintf("Couldn't find \"%s\"", text)
		blocks := buildImage(title, "https://media.giphy.com/media/l0Iy2hYDgmCjMufzq/giphy-downsized.gif")

		sendMessage(slackURL, "ephemeral", blocks)

		return defaultResp, nil
	}

	title := fmt.Sprintf("%s sent \"%s\"", body.Get("user_name"), text)
	blocks := buildImage(title, image.Image)

	sendMessage(slackURL, "in_channel", blocks)

	return defaultResp, nil
}

func main() {
	lambda.Start(handleRequest)
}
