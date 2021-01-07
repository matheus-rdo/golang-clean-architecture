# Clean Architecture on Golang
Some variations of Uncle Bob's clean architecture  

## Features
- Clean Architecture (obviously) 
- Environment configuration (using Viper)
- HTTP Layer (Gin Framework)

## Create resources

Run localstack:  
`docker-compose up -d`  

Create DynamoDB table  
```shell script

awslocal dynamodb create-table \
    --table-name books \
    --attribute-definitions AttributeName=bookId,AttributeType=S \
    --key-schema AttributeName=bookId,KeyType=HASH \
    --billing-mode PAY_PER_REQUEST

```


