package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/nlopes/slack"
)

func buildImage(title string, imageUrl string, responseType string) slack.Message {
	blockTitle := slack.NewTextBlockObject("plain_text", title, false, false)
	block := slack.NewImageBlock(imageUrl, title, "", blockTitle)
	msg := slack.NewBlockMessage(block)

	msg.Msg.ResponseType = responseType

	return msg
}

func buildSection(content string) slack.Message {
	blockText := slack.NewTextBlockObject("mrkdwn", content, false, false)
	block := slack.NewSectionBlock(blockText, nil, nil)
	msg := slack.NewBlockMessage(block)

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
