package orm

import (
	"time"

	"api.aifuxi.cool/db/orm/models"
)

type CreateUserParams struct {
	Nickname string
	Avatar   string
	Email    string
	Password string
}

type UpdateUserParams struct {
	Nickname string
	Avatar   string
	Password string
}

type ListUsersParams struct {
	Nickname string
	Email    string

	PaginationParams
	OrderParams
}

// 创建用户
func (q *Queries) CreateUser(arg CreateUserParams) (models.User, error) {
	user := models.User{
		Nickname: arg.Nickname,
		Avatar:   arg.Avatar,
		Email:    arg.Email,
		Password: arg.Password,
	}

	err := q.db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// 获取用户列表
func (q *Queries) ListUsers(arg ListUsersParams) ([]models.User, int64, error) {
	var users []models.User
	var count int64
	var queryDB = q.db.Model(models.User{}).Scopes(DeletedRecord)

	if len(arg.Nickname) > 0 {
		queryDB.Where("nickname LIKE ?", "%"+arg.Nickname+"%")
	}

	if len(arg.Email) > 0 {
		queryDB.Where("email LIKE ?", "%"+arg.Email+"%")
	}

	queryDB = queryDB.Count(&count)

	// order := fmt.Sprintf("%s %s", arg.OrderBy, arg.Order)
	err := queryDB.Scopes(Paginate(arg.Page, arg.PageSize)).Find(&users).Error
	if err != nil {
		return nil, count, err
	}

	return users, count, nil
}

// 获取用户
func (q *Queries) GetUser(id int64) (models.User, error) {
	var user models.User

	err := q.db.Scopes(DeletedRecord).First(&user, id).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// 更新用户
func (q *Queries) UpdateUser(id int64, arg UpdateUserParams) (models.User, error) {
	user := models.User{
		ID: id,
	}
	err := q.db.Model(&user).Scopes(DeletedRecord).Updates(
		models.User{
			Nickname: arg.Nickname,
			Avatar:   arg.Avatar,
			Password: arg.Password,
		}).Limit(1).Error
	if err != nil {
		return models.User{}, err
	}

	err = q.db.First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// 删除用户
func (q *Queries) DeleteUser(id int64) error {
	user := models.User{
		ID: id,
	}
	now := time.Now()

	err := q.db.Model(&user).Updates(
		models.User{
			DeletedAt: &now,
		}).Limit(1).Error
	if err != nil {
		return err
	}

	return nil
}
