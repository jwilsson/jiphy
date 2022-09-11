package main

import (
	"sort"

	"github.com/aws/aws-sdk-go/aws/session"

	utils "github.com/jwilsson/go-bot-utils"
)

type Image struct {
	GiphyURL  string `json:"giphy_url" dynamodbav:"giphy_url"`
	ImageName string `json:"image_name" dynamodbav:"image_name"`
	ImageURL  string `json:"image_url" dynamodbav:"image_url"`
}

func getImages(tableName string) (images []Image, err error) {
	s, err := session.NewSession()
	if err != nil {
		return images, err
	}

	err = utils.GetDynamodbData(s, tableName, &images)
	if err != nil {
		return images, err
	}

	sort.Slice(images, func(i int, j int) bool {
		return images[i].ImageName < images[j].ImageName
	})

	return images, nil
}
