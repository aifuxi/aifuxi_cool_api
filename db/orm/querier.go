package db

type Querier interface {
	ListUsers(arg ListUsersParams) ([]User, int64, error)
	CreateUser(arg CreateUserParams) (User, error)
	GetUserByID(id int64) (User, error)
	UpdateUser(id int64, arg UpdateUserParams) error
	DeleteUserByID(id int64) error

	ListTags(arg ListTagsParams) ([]Tag, int64, error)
	CreateTag(arg CreateTagParams) (Tag, error)
	GetTagByID(id int64) (Tag, error)
	UpdateTag(id int64, arg UpdateTagParams) error
	DeleteTagByID(id int64) error
}

var _ Querier = (*Queries)(nil)
