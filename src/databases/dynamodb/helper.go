package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type getItemInfo struct {
	HashKeyName   string
	HashKeyValue  string
	RangeKeyName  *string
	RangeKeyValue *string
}

func putItem(item interface{}, tableName string) error {
	av, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	return err
}

func getItem(getInfo *getItemInfo, tableName string) (*dynamodb.GetItemOutput, error) {
	keyMap := map[string]*dynamodb.AttributeValue{
		getInfo.HashKeyName: {
			S: aws.String(getInfo.HashKeyValue),
		},
	}
	if getInfo.RangeKeyName != nil && getInfo.RangeKeyValue != nil {
		keyMap[*getInfo.RangeKeyName] = &dynamodb.AttributeValue{
			S: aws.String(*getInfo.RangeKeyValue),
		}
	}
	return svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       keyMap,
	})
}
