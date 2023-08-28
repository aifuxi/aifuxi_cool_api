package mysql

import "api.aifuxi.cool/db/mysql/models"

type Querier interface {
	CreateUser(arg CreateUserParams) (models.User, error)
	UpdateUser(id int64, arg UpdateUserParams) (models.User, error)
	GetUser(id int64) (models.User, error)
	DeleteUser(id int64) error
	ListUsers(arg ListUsersParams) ([]models.User, int64, error)
}
