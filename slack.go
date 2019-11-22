package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Block struct {
	Type string `json:"type"`
}

type BlockText struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

type BlockTitle struct {
	BlockText

	Emoji bool `json:"emoji"`
}

type ImageBlock struct {
	Block

	AltText  string     `json:"alt_text"`
	ImageUrl string     `json:"image_url"`
	Title    BlockTitle `json:"title"`
}

type SectionBlock struct {
	Block

	Text BlockText `json:"text"`
}

type Response struct {
	Blocks       interface{} `json:"blocks"`
	ResponseType string      `json:"response_type"`
}

func buildImage(title string, imageUrl string) []*ImageBlock {
	blockTitle := BlockTitle{
		BlockText: BlockText{
			Text: title,
			Type: "plain_text",
		},

		Emoji: false,
	}

	blocks := []*ImageBlock{
		&ImageBlock{
			Block: Block{
				Type: "image",
			},

			AltText:  title,
			ImageUrl: imageUrl,
			Title:    blockTitle,
		},
	}

	return blocks
}

func buildSection(text string) []*SectionBlock {
	blockText := BlockText{
		Text: text,
		Type: "mrkdwn",
	}

	blocks := []*SectionBlock{
		&SectionBlock{
			Block: Block{
				Type: "section",
			},

			Text: blockText,
		},
	}

	return blocks
}

func sendMessage(url string, responseType string, blocks interface{}) error {
	body, _ := json.Marshal(Response{
		Blocks:       blocks,
		ResponseType: responseType,
	})

	log.Printf("Posting %s to %s", body, url)

	_, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)

	return err
}
