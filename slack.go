package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

func createImageMessage(image *Image, userName string, command string, responseType string) slack.Message {
	imageTitle := slack.NewTextBlockObject("plain_text", image.ImageName, false, false)
	imageBlock := slack.NewImageBlock(image.ImageURL, image.ImageName, "", imageTitle)

	text := fmt.Sprintf("Posted by %s using %s", userName, command)
	textBlock := slack.NewTextBlockObject("plain_text", text, false, false)
	contextBlock := slack.NewContextBlock("footer", textBlock)

	msg := slack.NewBlockMessage(imageBlock, contextBlock)
	msg.Msg.Text = image.ImageName
	msg.Msg.ResponseType = responseType

	return msg
}

func createListMessage(images []Image) slack.Message {
	blocks := make([]slack.Block, len(images))

	for i, image := range images {
		text := fmt.Sprintf("<%s|*%s*>", image.GiphyURL, image.ImageName)
		textBlock := slack.NewTextBlockObject("mrkdwn", text, false, false)
		accessory := slack.NewAccessory(
			slack.NewImageBlockElement(image.ImageURL, image.ImageName),
		)

		blocks[i] = slack.NewSectionBlock(textBlock, nil, accessory)
	}

	msg := slack.NewBlockMessage(blocks...)
	msg.Msg.ResponseType = "ephemeral"

	return msg
}
