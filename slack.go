package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/nlopes/slack"
)

func createImage(title string, imageUrl string, responseType string) slack.Message {
	blockTitle := slack.NewTextBlockObject("plain_text", title, false, false)
	block := slack.NewImageBlock(imageUrl, title, "", blockTitle)
	msg := slack.NewBlockMessage(block)

	msg.Msg.ResponseType = responseType

	return msg
}

func createList(images []Image) slack.Message {
	msg := slack.NewBlockMessage()

	for _, image := range images {
		text := fmt.Sprintf("<%s|*%s*>", image.GiphyURL, image.ImageName)

		textBlock := slack.NewTextBlockObject("mrkdwn", text, false, false)
		imageBlock := slack.NewImageBlockElement(image.ImageURL, image.ImageName)
		accessory := slack.NewAccessory(imageBlock)

		block := slack.NewSectionBlock(textBlock, nil, accessory)
		msg = slack.AddBlockMessage(msg, block)
	}

	msg.Msg.ResponseType = "ephemeral"

	return msg
}

func sendMessage(url string, message slack.Message) error {
	body, _ := json.Marshal(message)

	log.Printf("Posting %s to %s", body, url)

	_, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)

	return err
}
