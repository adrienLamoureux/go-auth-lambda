# Go-auth-lambda

## Getting Started

### Pre-requiert

- Go 1.13
- DynamoDBLocal (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.html)

### Run DynamoDB

Go to the DynamoDB folder to start it with (Unix)

```
java -Djava.library.path=./DynamoDBLocal_lib -jar DynamoDBLocal.jar -sharedDb

```

### Dependencies installation

```
go mod tidy
```

## Run Tests

```
TODO
```

## Run Local Server

```
go run src/dev_main.go
```

## Deployment

### Build

#### Dev environment

```
go build -o bin/handler ./src
```

#### Prod environment

```
GOOS=linux GOARCH=amd64 go build -tags lambda -o bin/handler ./src
```

### Deploy container

TODO

### Deploy API

TODO
