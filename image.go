package main

import (
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

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
		return nil, err
	}

	err = utils.GetDynamodbData(s, tableName, &images)
	if err != nil {
		return nil, err
	}

	sort.Slice(images, func(i int, j int) bool {
		return images[i].ImageName < images[j].ImageName
	})

	return images, nil
}

func getImage(input string, tableName string) (*Image, error) {
	s, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	svc := dynamodb.New(s)
	expressionAttributeValues := map[string]*dynamodb.AttributeValue{
		":n": {
			S: aws.String(input),
		},
	}

	result, err := svc.Query(&dynamodb.QueryInput{
		ExpressionAttributeValues: expressionAttributeValues,
		KeyConditionExpression:    aws.String("image_name = :n"),
		TableName:                 aws.String(tableName),
	})

	if err != nil {
		return nil, err
	}

	var images []Image

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &images)
	if err != nil {
		return nil, err
	}

	if len(images) == 0 {
		return nil, nil
	}

	return &images[0], nil
}
