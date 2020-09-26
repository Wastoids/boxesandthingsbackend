package storage

import (
	"time"

	"github.com/Wastoids/boxesandthingsbackend/domain"
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

type thing struct {
	ID          string
	BoxID       string
	Name        string
	Description string
	ExpiresOn   time.Time
}

const tableName = "boxesAndThings"
const gsiName = "thingsByBox"

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

func (d dynamo) getBoxesByEmail(email string) ([]*domain.Box, error) {
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

func (d dynamo) getThingsByBox(boxID string) ([]*domain.Thing, error) {
	params := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		IndexName:              aws.String(gsiName),
		KeyConditionExpression: aws.String("#boxid = :boxvalue"),
		ExpressionAttributeNames: map[string]*string{
			"#boxid": aws.String("boxID"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":boxvalue": {S: aws.String(boxID)},
		},
	}
	resp, err := d.dynamoConn.Query(params)
	if err != nil {
		logrus.Error("could not get the things by box", err)
		return nil, err
	}

	var things []thing
	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &things)
	if err != nil {
		logrus.Error("could not unmarshal the response into thing entities", err)
		return nil, err
	}

	return toThingDomainModels(things), nil
}

func toBoxDomainModels(boxes []box) []*domain.Box {
	var result []*domain.Box

	for _, b := range boxes {
		result = append(result, toBoxDomainModel(b))
	}

	return result
}

func toBoxDomainModel(box box) *domain.Box {
	return &domain.Box{
		Description: box.Description,
		ID:          box.ID,
		Name:        box.Name,
		ParentBoxID: box.ParentBoxID,
	}
}

func toThingDomainModels(things []thing) []*domain.Thing {
	var result []*domain.Thing

	for _, t := range things {
		result = append(result, toThingDomainModel(t))
	}

	return result
}

func toThingDomainModel(thing thing) *domain.Thing {
	return &domain.Thing{
		ID:          thing.ID,
		BoxID:       thing.BoxID,
		Name:        thing.Name,
		Description: thing.Description,
		ExpiresOn:   thing.ExpiresOn,
	}
}
