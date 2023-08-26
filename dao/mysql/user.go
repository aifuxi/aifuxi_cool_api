package mysql

import (
	"fmt"
	"time"

	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/models"
	"api.aifuxi.cool/myerror"
)

func GetUsers(data dto.GetUsersDTO) ([]models.User, int64, error) {
	var users []models.User
	var count int64
	var queryDB = db.Model(models.User{}).Scopes(isDeletedRecord)

	if len(data.Nickname) > 0 {
		queryDB.Where("nickname LIKE ?", "%"+data.Nickname+"%")
	}

	if len(data.Email) > 0 {
		queryDB.Where("email LIKE ?", "%"+data.Email+"%")
	}

	queryDB = queryDB.Count(&count)

	order := fmt.Sprintf("%s %s", data.OrderBy, data.Order)
	err := queryDB.Order(order).Scopes(Paginate(data.Page, data.PageSize)).Find(&users).Error
	if err != nil {
		return nil, count, err
	}

	return users, count, nil
}

func GetUserByID(id int64) (models.User, error) {
	var user models.User

	err := db.Scopes(isDeletedRecord).First(&user, id).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User

	err := db.Scopes(isDeletedRecord).Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func UpdateUserByID(id int64, data dto.UpdateUserDTO) error {
	err := db.Model(models.User{}).Scopes(isDeletedRecord).Where("id = ?", id).Updates(
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
	var user models.User

	db.Scopes(isDeletedRecord).Where("email = ?", email).First(&user)

	return user.ID != 0
}

func UserExistsByID(id int64) bool {
	var user models.User

	db.Scopes(isDeletedRecord).First(&user, id)

	return user.ID != 0
}

func CreateUser(data dto.CreateUserDTO) (models.User, error) {
	var user models.User

	if exists := UserExistsByEmail(data.Email); exists {
		return user, myerror.ErrorUserExists
	}

	user = models.User{
		Nickname: data.Nickname,
		Avatar:   data.Avatar,
		Email:    data.Email,
		Password: data.Password,
	}

	err := db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
