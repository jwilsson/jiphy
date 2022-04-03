package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

func createImageMessage(image *Image, userName string, command string, responseType string) slack.Message {
	msg := slack.NewBlockMessage()

	imageTitle := slack.NewTextBlockObject("plain_text", image.ImageName, false, false)
	imageBlock := slack.NewImageBlock(image.ImageURL, image.ImageName, "", imageTitle)
	msg = slack.AddBlockMessage(msg, imageBlock)

	text := fmt.Sprintf("Posted by %s using %s", userName, command)
	textBlock := slack.NewTextBlockObject("plain_text", text, false, false)
	contextBlock := slack.NewContextBlock("footer", textBlock)
	msg = slack.AddBlockMessage(msg, contextBlock)

	msg.Msg.Text = image.ImageName
	msg.Msg.ResponseType = responseType

	return msg
}

func createListMessage(images []Image) slack.Message {
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
