package databases

type IAccountDatabase interface {
	CreateAccount(accInfo *AccountInfo) error
	CreateAccountEmail(accEmailInfo *AccountEmailInfo) error
	GetAccountEmailInfo(email string) (*AccountEmailInfo, error)
	GetAccountInfo(accID string) (*AccountInfo, error)
	GetAccountFavMoviesInfo(accID string) ([]*AccountFavMovieInfo, error)
	CreateAccountFavMovie(accountFavMovieInfo *AccountFavMovieInfo) error
	DeleteAccountFavMovie(accID, movieID string) error
}

type AccountFavMovieInfo struct {
	AccID   string
	MovieID string
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
