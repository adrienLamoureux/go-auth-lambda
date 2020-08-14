package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var svc *dynamodb.DynamoDB

var (
	region   = "eu-central-1"
	endpoint = "http://localhost:8000"
)

func init() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:   &region,
			Endpoint: &endpoint,
		},
	}))

	// Create DynamoDB client
	svc = dynamodb.New(sess)
}
