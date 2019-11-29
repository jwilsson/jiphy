package main

import (
	"sort"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Image struct {
	GiphyUrl  string `json:"giphy_url" dynamodbav:"giphy_url"`
	Id        string `json:"id" dynamodbav:"id"`
	ImageName string `json:"image_name" dynamodbav:"image_name"`
	ImageURL  string `json:"image_url" dynamodbav:"image_url"`
}

func getImageNames(tableName string) ([]string, error) {
	svc := dynamodb.New(session.New())
	result, err := svc.Scan(&dynamodb.ScanInput{
		ProjectionExpression: aws.String("image_name"),
		TableName:            aws.String(tableName),
	})

	if err != nil {
		return nil, err
	}

	var images []Image

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &images)
	if err != nil {
		return nil, err
	}

	imageNames := make([]string, 0, len(images))

	for _, v := range images {
		imageNames = append(imageNames, "• "+v.ImageName)
	}

	sort.Strings(imageNames)

	return imageNames, nil
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
		ProjectionExpression:      aws.String("image_url"),
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
