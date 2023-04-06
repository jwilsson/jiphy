package main

import (
	"context"
	"sort"

	"github.com/aws/aws-sdk-go-v2/config"

	utils "github.com/jwilsson/go-bot-utils"
)

type Image struct {
	GiphyURL  string `json:"giphy_url" dynamodbav:"giphy_url"`
	ImageName string `json:"image_name" dynamodbav:"image_name"`
	ImageURL  string `json:"image_url" dynamodbav:"image_url"`
}

func getImages(ctx context.Context, tableName string) (images []Image, err error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return images, err
	}

	err = utils.GetDynamodbData(ctx, cfg, tableName, &images)
	if err != nil {
		return images, err
	}

	sort.Slice(images, func(i int, j int) bool {
		return images[i].ImageName < images[j].ImageName
	})

	return images, nil
}
