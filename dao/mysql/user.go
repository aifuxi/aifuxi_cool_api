package mysql

import (
	"fmt"
	"time"

	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/models"
	"api.aifuxi.cool/myerror"
)

func GetUsers(arg dto.GetUsersDTO) ([]models.User, int64, error) {
	var users []models.User
	var count int64
	var queryDB = db.Model(models.User{}).Scopes(isDeletedRecord)

	if len(arg.Nickname) > 0 {
		queryDB.Where("nickname LIKE ?", "%"+arg.Nickname+"%")
	}

	if len(arg.Email) > 0 {
		queryDB.Where("email LIKE ?", "%"+arg.Email+"%")
	}

	queryDB = queryDB.Count(&count)

	order := fmt.Sprintf("%s %s", arg.OrderBy, arg.Order)
	err := queryDB.Order(order).Scopes(Paginate(arg.Page, arg.PageSize)).Find(&users).Error
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

func UpdateUserByID(id int64, arg dto.UpdateUserDTO) error {
	err := db.Model(models.User{}).Scopes(isDeletedRecord).Where("id = ?", id).Updates(
		models.User{
			Nickname: arg.Nickname,
			Avatar:   arg.Avatar,
			Password: arg.Password,
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

func CreateUser(arg dto.CreateUserDTO) (models.User, error) {
	var user models.User

	if exists := UserExistsByEmail(arg.Email); exists {
		return user, myerror.ErrorUserExists
	}

	user = models.User{
		Nickname: arg.Nickname,
		Avatar:   arg.Avatar,
		Email:    arg.Email,
		Password: arg.Password,
	}

	err := db.Create(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
