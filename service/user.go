package service

import (
	"api.aifuxi.cool/dao/mysql"
	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/models"
	"api.aifuxi.cool/myerror"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(arg dto.GetUsersDTO) ([]models.User, int64, error) {
	return mysql.GetUsers(arg)
}

func GetUserProfile(email string) (models.User, error) {
	return mysql.GetUserByEmail(email)
}

func CreateUser(arg dto.CreateUserDTO) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(arg.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	arg.Password = string(hashedPassword)

	return mysql.CreateUser(arg)
}

func GetUserByID(id int64) (models.User, error) {
	return mysql.GetUserByID(id)
}

func UpdateUserByID(id int64, arg dto.UpdateUserDTO) error {
	if !mysql.UserExistsByID(id) {
		return myerror.ErrorUserNotFound
	}

	return mysql.UpdateUserByID(id, arg)
}

func DeleteUserByID(id int64) error {
	return mysql.DeleteUserByID(id)
}
