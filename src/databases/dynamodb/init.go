package dynamodb

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var svc *dynamodb.DynamoDB

func init() {
	region := os.Getenv("DYNAMO_REGION")
	endpoint := os.Getenv("DYNAMO_ENDPOINT")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:   &region,
			Endpoint: &endpoint,
		},
	}))

	// Create DynamoDB client
	svc = dynamodb.New(sess)
}
