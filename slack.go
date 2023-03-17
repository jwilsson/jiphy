package main

import (
	"fmt"

	utils "github.com/jwilsson/go-bot-utils"
	"github.com/slack-go/slack"
	"golang.org/x/exp/slices"
)

type MessageInput struct {
	Command     string
	ImageName   string
	Images      []Image
	ResponseURL string
	UserName    string
}

func createImageMessage(image Image, userName string, command string, responseType string) slack.Message {
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

func sendMessage(input MessageInput) error {
	if input.ImageName == "list" {
		return utils.SendMessage(input.ResponseURL, createListMessage(input.Images))
	}

	i := slices.IndexFunc(input.Images, func(img Image) bool {
		return img.ImageName == input.ImageName
	})

	var responseType string
	var image Image

	if i >= 0 {
		responseType = "in_channel"
		image = input.Images[i]
	} else {
		responseType = "ephemeral"
		image = Image{
			GiphyURL:  "https://giphy.com/gifs/stonehampress-funny-horse-l0Iy2hYDgmCjMufzq",
			ImageName: "gif",
			ImageURL:  "https://media.giphy.com/media/l0Iy2hYDgmCjMufzq/giphy-downsized.gif",
		}
	}

	return utils.SendMessage(input.ResponseURL, createImageMessage(image, input.UserName, input.Command, responseType))
}
