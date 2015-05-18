package loco

import (
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
)

type Database struct {
	dynamoDB *dynamodb.DynamoDB
}

func NewDatabase(dynamoDB *dynamodb.DynamoDB) *Database {
	return &Database{dynamoDB: dynamoDB}
}
