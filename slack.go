package main

import (
	"fmt"

	"github.com/slack-go/slack"
)

func createImage(image *Image, userName string, command string, responseType string) slack.Message {
	msg := slack.NewBlockMessage()

	text := fmt.Sprintf("<%s|*%s*>", image.GiphyURL, image.ImageName)
	textBlock := slack.NewTextBlockObject("mrkdwn", text, false, false)
	sectionBlock := slack.NewSectionBlock(textBlock, nil, nil)
	msg = slack.AddBlockMessage(msg, sectionBlock)

	text = fmt.Sprintf("Posted by %s using %s", userName, command)
	imageTitle := slack.NewTextBlockObject("plain_text", text, false, false)
	imageBlock := slack.NewImageBlock(image.ImageURL, image.ImageName, "", imageTitle)
	msg = slack.AddBlockMessage(msg, imageBlock)

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
