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

### Without docker

```
export DYNAMO_REGION=my_region
```

```
export DYNAMO_ENDPOINT=my_endpoint
```

```
go run src/dev_main.go
```

### With a docker image

```
docker image build -t go-auth-lambda:X.X .
```

```
docker container run --name go-auth-lambda -e DYNAMO_REGION=X -e DYNAMO_ENDPOINT=X -e AWS_ACCESS_KEY_ID=X -e AWS_SECRET_ACCESS_KEY=X -p X:7200 -d go-auth-lambda:X.X
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
