package handlers

import (
	"github.com/adrienLamoureux/go-auth-lambda/src/databases"
	"github.com/adrienLamoureux/go-auth-lambda/src/databases/dynamodb"
)

var accountDatabase databases.IAccountDatabase

// TODO: Use config to switch of DB system instead, for example
func init() {
	accountDatabase = &dynamodb.AccountDynamoDB{}
}
