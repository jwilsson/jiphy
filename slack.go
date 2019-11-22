package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type BlockTitle struct {
	Emoji bool   `json:"emoji"`
	Text  string `json:"text"`
	Type  string `json:"type"`
}

type Block struct {
	AltText  string     `json:"alt_text"`
	ImageUrl string     `json:"image_url"`
	Title    BlockTitle `json:"title"`
	Type     string     `json:"type"`
}

type Response struct {
	Blocks       []*Block `json:"blocks"`
	ResponseType string   `json:"response_type"`
}

func buildMessage(title string, imageUrl string) []*Block {
	blockTitle := BlockTitle{
		Emoji: true,
		Text:  title,
		Type:  "plain_text",
	}

	blocks := []*Block{
		&Block{
			AltText:  title,
			ImageUrl: imageUrl,
			Title:    blockTitle,
			Type:     "image",
		},
	}

	return blocks
}

func sendMessage(url string, responseType string, blocks []*Block) error {
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
