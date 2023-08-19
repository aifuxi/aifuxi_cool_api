package mysql

import (
	"time"

	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/internal"
	"api.aifuxi.cool/models"
	"api.aifuxi.cool/myerror"
)

func GetUsers() (*[]models.User, error) {
	users := new([]models.User)

	err := db.Where("deleted_at is null").Find(users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(id int64) (*models.User, error) {
	user := new(models.User)
	err := db.Where("deleted_at is null").First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	user := new(models.User)
	err := db.Where("deleted_at is null and email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUserByID(id int64, data *dto.UpdateUserDTO) error {
	var user = &models.User{
		ID: id,
	}
	err := db.Model(user).Where("deleted_at is null").Updates(
		models.User{
			Nickname: data.Nickname,
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
	db.Where("deleted_at is null and email = ?", email).First(&user)
	return user.ID != 0
}

func UserExistsByID(id int64) bool {
	user := new(models.User)
	db.Where("deleted_at is null").First(&user, id)
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
		Email:    data.Email,
		Password: data.Password,
	}
	err = db.Model(&models.User{}).Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
