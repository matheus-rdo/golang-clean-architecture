package database

import (
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/matheushr97/golang-clean-architecture/app"
	"github.com/matheushr97/golang-clean-architecture/infra/aws"
)

var once sync.Once
var instance *dynamodb.DynamoDB

// GetDynamoClient get dynamodb client instance
func getDynamoClient() *dynamodb.DynamoDB {
	once.Do(func() {
		sess := aws.GetAwsSession(app.ENV.AwsRegion)
		instance = dynamodb.New(sess)
	})
	return instance
}

// DynamoRepository DynamoDB repository type
type DynamoRepository struct {
	Database  *dynamodb.DynamoDB
	TableName string
}

// NewDynamoRepository init dynamodb repository
func NewDynamoRepository(tableName string) DynamoRepository {
	if tableName == "" {
		log.Fatal("DynamoDB Table name is missing")
	}
	return DynamoRepository{
		Database:  getDynamoClient(),
		TableName: tableName,
	}
}

// PutItem puts item on the database
func (repository DynamoRepository) PutItem(item interface{}) (*dynamodb.PutItemOutput, error) {
	info, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: &repository.TableName,
	}

	return repository.Database.PutItem(input)
}
