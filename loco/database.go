package loco

import (
	"github.com/awslabs/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type Database struct {
	dynamoDB *dynamodbiface.DynamoDBAPI
}

func NewDatabase(dynamoDB *dynamodbiface.DynamoDBAPI) *Database {
	return &Database{dynamoDB: dynamoDB}
}
