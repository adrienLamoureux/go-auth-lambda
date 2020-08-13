package databases

type IAccountDatabase interface {
	CreateAccount(accInfo *AccountInfo) error
	CreateAccountEmail(accEmailInfo *AccountEmailInfo) error
	GetAccountEmailByEmail(email string) (*AccountEmailInfo, error)
	GetAccountInfo(accID string) (*AccountInfo, error)
}

type AccountInfo struct {
	AccID     string
	Email     string
	Password  string
	FirstName string
	LastName  string
}

type AccountEmailInfo struct {
	AccID string
	Email string
}
