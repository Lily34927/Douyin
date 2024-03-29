package logic

import (
	"fmt"
	"groupwork/module"
	"groupwork/repository"
	"groupwork/repository/mysql"
)

func Register(username, password string) (newUser *module.User, err error) {
	// 1.判断用户存不存在
	if err = mysql.CheckUserExist(username); err != nil {
		return
	}

	// 2. 生成UID
	userIdSequence := module.GenID()

	// 创建User实例
	newUser = &module.User{
		UserId:   userIdSequence,
		Name:     username,
		Password: password,
	}

	// 3.添加到数据库中
	if err = mysql.InsertUser(newUser); err != nil {
		// 用户注册失败
		return nil, repository.ErrorRegister
	}

	// 4.生成token
	token, err := module.GenToken(newUser.UserId, newUser.Password)
	// 获取token失败
	if err != nil {

	}
	newUser.Token = token
	return
}

func Login(username, password string) (user *module.User, err error) {
	user = &module.User{
		Name:     username,
		Password: password,
	}

	// 1.验证用户是否存在（用户名和密码）
	if err = mysql.CheckLoginUser(user); err != nil {
		return nil, repository.ErrorUserInfo
	}

	// 2.生成token
	//token := username + password
	token, err := module.GenToken(user.UserId, user.Password)
	if err != nil {
		fmt.Println("genToken failed")
		return
	}

	user.Token = token
	return
}

func UserInfo(userId int64) (user *module.User, err error) {
	user = &module.User{
		UserId: userId,
	}
	err = mysql.GetUserInfo(user)
	if err != nil {
		return nil, err
	}
	return
}
