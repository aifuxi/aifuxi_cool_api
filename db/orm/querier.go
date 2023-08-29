package db

type Querier interface {
	ListUsers(arg ListUsersParams) ([]User, int64, error)
	CreateUser(arg CreateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
