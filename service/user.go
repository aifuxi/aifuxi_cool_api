package service

import (
	"github.com/aifuxi/aifuxi_cool_api/dao/mysql"
	"github.com/aifuxi/aifuxi_cool_api/dto"
	"github.com/aifuxi/aifuxi_cool_api/models"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers() (*[]models.User, error) {
	return mysql.GetUsers()
}

func CreateUser(data *dto.CreateUserDTO) (*models.User, error) {
	// encrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	data.Password = string(hashedPassword)

	return mysql.CreateUser(data)
}

func GetUserByID(id int64) (*models.User, error) {
	return mysql.GetUserByID(id)
}

func UpdateUserByID(id int64, data *dto.UpdateUserDTO) error {
	if !mysql.UserExistsByID(id) {
		return mysql.ErrorUserNotFound
	}

	return mysql.UpdateUserByID(id, data)
}

func DeleteUserByID(id int64) error {
	return mysql.DeleteUserByID(id)
}
