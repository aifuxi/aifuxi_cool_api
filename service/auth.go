package service

import (
	"api.aifuxi.cool/dao/mysql"
	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/internal"
	"api.aifuxi.cool/models"
	"api.aifuxi.cool/myerror"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(data *dto.SignUpDTO) (*models.User, error) {
	return CreateUser(&dto.CreateUserDTO{
		Nickname:   data.Nickname,
		Email:      data.Email,
		Password:   data.Password,
		RePassword: data.RePassword,
	})
}

func SignIn(data *dto.SignInDTO) (string, error) {
	// 1. 根据email查找用户
	if exists := mysql.UserExistsByEmail(data.Email); !exists {
		return "", myerror.ErrorUserNotFound
	}

	user, err := mysql.GetUserByEmail(data.Email)
	if err != nil {
		return "", err
	}
	// 2. 判断密码对不对
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return "", myerror.ErrorIncorrectPassword
	}

	// 3. 生成 token
	token, err := internal.GenToken(data.Email)

	return token, err
}
