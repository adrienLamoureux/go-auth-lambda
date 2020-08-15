package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type keyItemInfo struct {
	HashKeyName   string
	HashKeyValue  string
	RangeKeyName  *string
	RangeKeyValue *string
	IndexName     *string
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

func getItem(keyInfo *keyItemInfo, tableName string) (*dynamodb.GetItemOutput, error) {
	keyMap := map[string]*dynamodb.AttributeValue{
		keyInfo.HashKeyName: {
			S: aws.String(keyInfo.HashKeyValue),
		},
	}
	if keyInfo.RangeKeyName != nil && keyInfo.RangeKeyValue != nil {
		keyMap[*keyInfo.RangeKeyName] = &dynamodb.AttributeValue{
			S: aws.String(*keyInfo.RangeKeyValue),
		}
	}
	return svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       keyMap,
	})
}

func getItems(keyInfo *keyItemInfo, tableName string) (*dynamodb.ScanOutput, error) {
	filt := expression.Name(keyInfo.HashKeyName).Equal(expression.Value(keyInfo.HashKeyValue))
	if keyInfo.RangeKeyName != nil && keyInfo.RangeKeyValue != nil {
		filt = filt.And(expression.Name(*keyInfo.RangeKeyName).Equal(expression.Value(*keyInfo.RangeKeyValue)))
	}
	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		return nil, err
	}
	return svc.Scan(&dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String(tableName),
	})
}

func deleteItem(keyInfo *keyItemInfo, tableName string) error {
	keyMap := map[string]*dynamodb.AttributeValue{
		keyInfo.HashKeyName: {
			S: aws.String(keyInfo.HashKeyValue),
		},
	}
	if keyInfo.RangeKeyName != nil && keyInfo.RangeKeyValue != nil {
		keyMap[*keyInfo.RangeKeyName] = &dynamodb.AttributeValue{
			S: aws.String(*keyInfo.RangeKeyValue),
		}
	}
	_, err := svc.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key:       keyMap,
	})
	return err
}
