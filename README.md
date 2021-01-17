# Clean Architecture on Golang
My version of Uncle Bob's Clean Architecture  

## Stack
- Gin Web Framework
- DynamoDB database
- Environment configuration

## Local development

### Requirements
- Go
- [Localstack](https://github.com/localstack/localstack)
- [awslocal (CLI)](https://github.com/localstack/awscli-local)

### Create resources

Start localstack  
`docker-compose up -d`

Create DynamoDB table  
```shell script

awslocal dynamodb create-table \
    --table-name books \
    --attribute-definitions AttributeName=bookId,AttributeType=S \
    --key-schema AttributeName=bookId,KeyType=HASH \
    --billing-mode PAY_PER_REQUEST

```
### Run application
`go run .`