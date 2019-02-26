package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Attachment struct {
	AuthorName string `json:"author_name"`
	Fallback   string `json:"fallback"`
	ImageURL   string `json:"image_url"`
}

type Response struct {
	Attachments  []*Attachment `json:"attachments"`
	ResponseType string        `json:"response_type"`
}

func handleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received body: %v", request.Body)

	body, _ := url.ParseQuery(request.Body)
	text := body.Get("text")
	userName := body.Get("user_name")

	if text == "" {
		return events.APIGatewayProxyResponse{
			Body:       "",
			Headers:    map[string]string{},
			StatusCode: 200,
		}, nil
	}

	image := GetImage(text)

	var attachments []*Attachment
	if image != nil {
		attachments = []*Attachment{
			&Attachment{
				AuthorName: fmt.Sprintf("Sent by %s", userName),
				Fallback:   text,
				ImageURL:   image.Image,
			},
		}
	} else {
		attachments = []*Attachment{
			&Attachment{
				AuthorName: fmt.Sprintf("Sent by %s", userName),
				Fallback:   fmt.Sprintf("Nothing found for *%s* :confused:", text),
				ImageURL:   "https://media.giphy.com/media/5L3qVEnmZi8Yo/giphy-downsized.gif",
			},
		}
	}

	slackURL, _ := url.QueryUnescape(body.Get("response_url"))
	jsonBody, _ := json.Marshal(Response{
		ResponseType: "in_channel",
		Attachments:  attachments,
	})

	log.Printf("Posting %s to %s", jsonBody, slackURL)

	http.Post(slackURL, "application/json", bytes.NewBuffer(jsonBody))

	return events.APIGatewayProxyResponse{
		Body:       "",
		Headers:    map[string]string{},
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
