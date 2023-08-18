package service

import (
	"github.com/aifuxi/aifuxi_cool_api/dao/mysql"
	"github.com/aifuxi/aifuxi_cool_api/dto"
	"github.com/aifuxi/aifuxi_cool_api/models"
)

func GetUsers() (*[]models.User, error) {
	return mysql.GetUsers()
}

func CreateUser(data *dto.CreateUserDTO) (*models.User, error) {
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
