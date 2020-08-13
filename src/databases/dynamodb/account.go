package dynamodb

import (
	"time"

	"github.com/adrienLamoureux/go-auth-lambda/src/databases"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type AccountDynamoDB struct {
}

func (accountDynamoDB *AccountDynamoDB) CreateAccount(accInfo *databases.AccountInfo) error {
	timeNow := time.Now().UnixNano()
	return putItem(&accountTableData{
		AccID:     accInfo.AccID,
		Email:     accInfo.Email,
		Password:  accInfo.Password,
		FirstName: accInfo.FirstName,
		LastName:  accInfo.LastName,
		CreateTm:  timeNow,
		UpdateTm:  timeNow,
	}, accountTableDefaultName)
}

func (accountDynamoDB *AccountDynamoDB) CreateAccountEmail(accEmailInfo *databases.AccountEmailInfo) error {
	timeNow := time.Now().UnixNano()
	return putItem(&accountEmailTableData{
		AccID:    accEmailInfo.AccID,
		Email:    accEmailInfo.Email,
		CreateTm: timeNow,
		UpdateTm: timeNow,
	}, accountTableDefaultName)
}

func (accountDynamoDB *AccountDynamoDB) GetAccountEmailByEmail(email string) (*databases.AccountEmailInfo, error) {
	result, err := getItem(&getItemInfo{
		HashKeyName:  "email",
		HashKeyValue: email,
	}, accountEmailTableDefaultName)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	item := accountEmailTableData{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}

	return item.toAbstract(), nil
}

func (accountDynamoDB *AccountDynamoDB) GetAccountInfo(accID string) (*databases.AccountInfo, error) {
	result, err := getItem(&getItemInfo{
		HashKeyName:  "accId",
		HashKeyValue: accID,
	}, accountTableDefaultName)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	item := accountTableData{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}

	return item.toAbstract(), nil
}

type accountTableData struct {
	AccID     string `json:"accId"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	CreateTm  int64  `json:"createTm"`
	UpdateTm  int64  `json:"updateTm"`
}

func (accountData *accountTableData) toAbstract() *databases.AccountInfo {
	return &databases.AccountInfo{
		AccID:     accountData.AccID,
		Email:     accountData.Email,
		Password:  accountData.Password,
		FirstName: accountData.FirstName,
		LastName:  accountData.LastName,
	}
}

type accountEmailTableData struct {
	AccID    string `json:"accId"`
	Email    string `json:"email"`
	CreateTm int64  `json:"createTm"`
	UpdateTm int64  `json:"updateTm"`
}

func (accountData *accountEmailTableData) toAbstract() *databases.AccountEmailInfo {
	return &databases.AccountEmailInfo{
		AccID: accountData.AccID,
		Email: accountData.Email,
	}
}
