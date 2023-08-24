package mysql

import (
	"fmt"
	"time"

	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/internal"
	"api.aifuxi.cool/models"
	"api.aifuxi.cool/myerror"
)

func GetUsers(data *dto.GetUsersDTO) (*[]models.User, int64, error) {
	users := new([]models.User)
	var total int64
	var nicknameLike, emailLike string

	order := fmt.Sprintf("%s %s", data.OrderBy, data.Order)
	if len(data.Nickname) > 0 {
		nicknameLike = "%" + data.Nickname + "%"
	}

	if len(data.Email) > 0 {
		emailLike = "%" + data.Email + "%"
	}

	if len(data.Nickname) == 0 && len(data.Email) == 0 {
		nicknameLike = "%" + data.Nickname + "%"
		emailLike = "%" + data.Email + "%"
	}

	err := db.Order(order).Scopes(isDeletedRecord, Paginate(data.Page, data.PageSize)).Where(
		db.Where("nickname LIKE ?", nicknameLike).Or("email LIKE ?", emailLike),
	).Find(users).Error
	if err != nil {
		return nil, total, err
	}

	err = db.Model(models.User{}).Scopes(isDeletedRecord).Where(
		db.Where("nickname LIKE ?", nicknameLike).Or("email LIKE ?", emailLike),
	).Count(&total).Error
	if err != nil {
		return nil, total, err
	}

	return users, total, nil
}

func GetUserByID(id int64) (*models.User, error) {
	user := new(models.User)
	err := db.Scopes(isDeletedRecord).First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	user := new(models.User)
	err := db.Scopes(isDeletedRecord).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUserByID(id int64, data *dto.UpdateUserDTO) error {
	var user = &models.User{
		ID: id,
	}
	err := db.Model(user).Scopes(isDeletedRecord).Updates(
		models.User{
			Nickname: data.Nickname,
			Avatar:   data.Avatar,
			Password: data.Password,
		}).Limit(1).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteUserByID(id int64) error {
	if !UserExistsByID(id) {
		return myerror.ErrorUserNotFound
	}

	err := db.Model(models.User{}).Where("id = ?", id).Limit(1).Update("deleted_at", time.Now().Local().Format(time.DateTime)).Error
	if err != nil {
		return err
	}

	return nil
}

func UserExistsByEmail(email string) bool {
	user := new(models.User)
	db.Scopes(isDeletedRecord).Where("email = ?", email).First(&user)
	return user.ID != 0
}

func UserExistsByID(id int64) bool {
	user := new(models.User)
	db.Scopes(isDeletedRecord).First(&user, id)
	return user.ID != 0
}

func CreateUser(data *dto.CreateUserDTO) (*models.User, error) {
	if exists := UserExistsByEmail(data.Email); exists {
		return nil, myerror.ErrorUserExists
	}

	id, err := internal.GenSnowflakeID()
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       id,
		Nickname: data.Nickname,
		Avatar:   data.Avatar,
		Email:    data.Email,
		Password: data.Password,
	}
	err = db.Model(&models.User{}).Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
