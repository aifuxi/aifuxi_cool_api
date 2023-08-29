package db

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

var (
	ErrUserExist = errors.New("用户已存在")
)

type ExistUserParams struct {
	ID    int64
	Email string
}

func (q *Queries) ExistUser(arg ExistUserParams) (bool, error) {
	var user User
	cond := User{
		ID:    arg.ID,
		Email: arg.Email,
	}

	err := q.db.Debug().First(&user, cond).Error
	if err != nil {
		// 如果 err 是 ErrRecordNotFound，只是记录没找到，不认为是出错了
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	if user.ID != 0 {
		return true, ErrUserExist
	}

	return false, nil
}

type ListUsersParams struct {
	Email    string
	Nickname string

	Page     int
	PageSize int
	Order    string
	OrderBy  string
}

func (q *Queries) ListUsers(arg ListUsersParams) ([]User, int64, error) {
	var users []User
	var count int64

	queryDB := q.db.Model(User{}).Scopes(isDeleted)

	if len(arg.Nickname) > 0 {
		queryDB.Where("nickname LIKE ?", "%"+arg.Nickname+"%")
	}

	if len(arg.Email) > 0 {
		queryDB.Where("email LIKE ?", "%"+arg.Email+"%")
	}

	queryDB = queryDB.Count(&count)

	order := fmt.Sprintf("%s %s", arg.OrderBy, arg.Order)
	err := queryDB.Order(order).Scopes(paginate(arg.Page, arg.PageSize)).Find(&users).Error
	if err != nil {
		return nil, count, err
	}

	return users, count, nil
}

type CreateUserParams struct {
	Nickname string
	Avatar   string
	Email    string
	Password string
}

func (q *Queries) CreateUser(arg CreateUserParams) (User, error) {
	user := User{
		Nickname: arg.Nickname,
		Avatar:   arg.Avatar,
		Email:    arg.Email,
		Password: arg.Password,
	}

	exitUserArg := ExistUserParams{Email: arg.Email}
	exist, err := q.ExistUser(exitUserArg)
	if err != nil {
		return User{}, err
	}
	if exist {
		return User{}, err
	}

	err = q.db.Create(&user).Error
	if err != nil {
		return User{}, err
	}

	return user, nil
}
