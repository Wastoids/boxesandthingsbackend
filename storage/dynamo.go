package storage

import (
	"github.com/Wastoids/boxesandthingsbackend/models"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/sirupsen/logrus"
)

type box struct {
	ID          string `json:"id" dynamodbav:"id"`
	Name        string `json:"name" dynamodbav:"name"`
	Description string `json:"description" dynamodbav:"description"`
	ParentBoxID string `json:"parentBoxID" dynamodbav:"parentBoxID"`
}

const tableName = "boxesAndThings"

type dynamo struct {
	dynamoConn *dynamodb.DynamoDB
}

func newDynamo() dynamo {
	mySession := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String("http://127.0.0.1:8000"),
		Region:   aws.String("us-east-1")}))
	svc := dynamodb.New(mySession)
	return dynamo{dynamoConn: svc}
}

func (d dynamo) getBoxesByEmail(email string) ([]*models.Box, error) {
	params := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("#email = :emailvalue AND begins_with(#myKey, :bbbbb)"),
		ExpressionAttributeNames: map[string]*string{
			"#email": aws.String("email"),
			"#myKey": aws.String("artifact_type"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":emailvalue": {S: aws.String(email)},
			":bbbbb":      {S: aws.String("box")},
		},
	}
	resp, err := d.dynamoConn.Query(params)
	if err != nil {
		logrus.Error("could not get the boxes by email", err)
		return nil, err
	}

	var boxes []box
	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &boxes)
	if err != nil {
		logrus.Error("could not unmarshal the response into box entities", err)
		return nil, err
	}

	return toBoxDomainModels(boxes), nil
}

func toBoxDomainModels(boxes []box) []*models.Box {
	var result []*models.Box

	for _, b := range boxes {
		result = append(result, toBoxDomainModel(b))
	}

	return result
}

func toBoxDomainModel(box box) *models.Box {
	return &models.Box{
		Description: box.Description,
		ID:          box.ID,
		Name:        box.Name,
		ParentBoxID: box.ParentBoxID,
	}
}
