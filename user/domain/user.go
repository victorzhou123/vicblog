package domain

type User struct {
	Email Email

	Account Account
}

type Account struct {
	Username Username
	Password Password
}
