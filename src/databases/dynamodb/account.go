package dynamodb

import (
	"time"

	"github.com/adrienLamoureux/go-auth-lambda/src/databases"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type AccountDynamoDB struct {
}

func (accountDynamoDB *AccountDynamoDB) CreateAccount(accInfo *databases.AccountInfo) error {
	timeNow := time.Now().Unix()
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
	timeNow := time.Now().Unix()
	return putItem(&accountEmailTableData{
		AccID:    accEmailInfo.AccID,
		Email:    accEmailInfo.Email,
		CreateTm: timeNow,
		UpdateTm: timeNow,
	}, accountTableDefaultName)
}

func (accountDynamoDB *AccountDynamoDB) GetAccountEmailInfo(email string) (*databases.AccountEmailInfo, error) {
	result, err := getItem(&keyItemInfo{
		HashKeyName:  "email",
		HashKeyValue: email,
	}, accountEmailTableDefaultName)
	if err != nil {
		return nil, err
	}
	if result == nil {
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
	result, err := getItem(&keyItemInfo{
		HashKeyName:  "accId",
		HashKeyValue: accID,
	}, accountTableDefaultName)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}

	item := accountTableData{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}

	return item.toAbstract(), nil
}

func (accountDynamoDB *AccountDynamoDB) GetAccountFavMoviesInfo(accID string) ([]*databases.AccountFavMovieInfo, error) {
	result, err := getItems(&keyItemInfo{
		HashKeyName:  "accId",
		HashKeyValue: accID,
	}, accountFavMovieTableDefaultName)
	if err != nil {
		return []*databases.AccountFavMovieInfo{}, err
	}
	if result == nil {
		return []*databases.AccountFavMovieInfo{}, nil
	}

	accountFavMovieInfoList := make([]*databases.AccountFavMovieInfo, len(result.Items))
	for i, item := range result.Items {
		accountFavMovieData := accountFavMovieTableData{}
		err = dynamodbattribute.UnmarshalMap(item, &accountFavMovieData)
		if err != nil {
			return []*databases.AccountFavMovieInfo{}, err
		}
		accountFavMovieInfoList[i] = accountFavMovieData.toAbstract()
	}
	return accountFavMovieInfoList, nil
}

func (accountDynamoDB *AccountDynamoDB) CreateAccountFavMovie(accountFavMovieInfo *databases.AccountFavMovieInfo) error {
	timeNow := time.Now().Unix()
	return putItem(&accountFavMovieTableData{
		AccID:    accountFavMovieInfo.AccID,
		MovieID:  accountFavMovieInfo.MovieID,
		CreateTm: timeNow,
		UpdateTm: timeNow,
	}, accountFavMovieTableDefaultName)
}

func (accountDynamoDB *AccountDynamoDB) DeleteAccountFavMovie(accID, movieID string) error {
	movieKeyName := "movieId"
	return deleteItem(&keyItemInfo{
		HashKeyName:   "accId",
		HashKeyValue:  accID,
		RangeKeyName:  &movieKeyName,
		RangeKeyValue: &movieID,
	}, accountFavMovieTableDefaultName)
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

type accountFavMovieTableData struct {
	AccID    string `json:"accId"`
	MovieID  string `json:"movieId"`
	CreateTm int64  `json:"createTm"`
	UpdateTm int64  `json:"updateTm"`
}

func (accountFavMovieData *accountFavMovieTableData) toAbstract() *databases.AccountFavMovieInfo {
	return &databases.AccountFavMovieInfo{
		AccID:   accountFavMovieData.AccID,
		MovieID: accountFavMovieData.MovieID,
	}
}
