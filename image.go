package main

import (
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Image struct {
	GiphyURL  string `json:"giphy_url" dynamodbav:"giphy_url"`
	ImageName string `json:"image_name" dynamodbav:"image_name"`
	ImageURL  string `json:"image_url" dynamodbav:"image_url"`
}

func getImages(tableName string) ([]Image, error) {
	svc := dynamodb.New(session.New())
	result, err := svc.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})

	if err != nil {
		return nil, err
	}

	var images []Image

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &images)
	if err != nil {
		return nil, err
	}

	sort.Slice(images, func(i int, j int) bool {
		return images[i].ImageName < images[j].ImageName
	})

	return images, nil
}

func getImage(input string, tableName string) (*Image, error) {
	expressionAttributeValues := map[string]*dynamodb.AttributeValue{
		":n": {
			S: aws.String(input),
		},
	}

	svc := dynamodb.New(session.New())
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
