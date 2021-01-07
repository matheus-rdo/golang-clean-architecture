package repository

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"github.com/matheushr97/golang-clean-architecture/app"
	"github.com/matheushr97/golang-clean-architecture/domain"
	"github.com/matheushr97/golang-clean-architecture/infra/database"
)

type bookRepository struct {
	database.DynamoRepository
}

// NewBookRepository creates a new book repository using dynamoDB as database
func NewBookRepository() domain.BookRepository {
	return &bookRepository{
		DynamoRepository: database.NewDynamoRepository(app.ENV.BookTableName),
	}
}

func (repository *bookRepository) Create(ctx context.Context, book domain.Book) (res *domain.Book, err error) {
	now := time.Now()
	book.ID = uuid.New().String()
	book.CreatedAt = now
	book.UpdatedAt = now

	if _, err := repository.PutItem(book); err != nil {
		return nil, err
	}

	return &book, nil
}

func (repository *bookRepository) Fetch(ctx context.Context) (res *[]domain.Book, err error) {
	scanInput := dynamodb.ScanInput{
		TableName: &repository.TableName,
	}
	result, err := repository.Database.ScanWithContext(ctx, &scanInput)
	if err != nil {
		return nil, err
	}
	var books []domain.Book
	if err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &books); err != nil {
		return nil, err
	}

	return &books, nil
}

func (repository *bookRepository) GetByID(ctx context.Context, id string) (*domain.Book, error) {
	itemQuery := &dynamodb.GetItemInput{
		TableName: &repository.TableName,
		Key: map[string]*dynamodb.AttributeValue{
			"bookId": {
				S: &id,
			},
		},
	}
	result, err := repository.Database.GetItemWithContext(ctx, itemQuery)
	if err != nil {
		return nil, err
	}
	var book domain.Book
	if err = dynamodbattribute.UnmarshalMap(result.Item, &book); err != nil {
		return nil, err
	}

	return &book, nil
}
