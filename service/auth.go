package service

import (
	"api.aifuxi.cool/dao/mysql"
	"api.aifuxi.cool/dto"
	"api.aifuxi.cool/internal"
	"api.aifuxi.cool/models"
	"api.aifuxi.cool/myerror"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(arg dto.SignUpDTO) (models.User, error) {
	return CreateUser(dto.CreateUserDTO{
		Nickname:   arg.Nickname,
		Email:      arg.Email,
		Password:   arg.Password,
		RePassword: arg.RePassword,
	})
}

func SignIn(arg dto.SignInDTO) (string, models.User, error) {
	// 1. 根据email查找用户
	if exists := mysql.UserExistsByEmail(arg.Email); !exists {
		return "", models.User{}, myerror.ErrorUserNotFound
	}

	user, err := mysql.GetUserByEmail(arg.Email)
	if err != nil {
		return "", models.User{}, err
	}

	// 2. 判断密码对不对
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(arg.Password))
	if err != nil {
		return "", models.User{}, myerror.ErrorIncorrectPassword
	}

	// 3. 生成 token
	token, err := internal.GenToken(arg.Email)

	return token, user, err
}
