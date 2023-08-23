package repository

import "errors"

var (
	// 用户
	ErrorUserExist = errors.New("User already exist")
	ErrorUserInfo  = errors.New("User doesn't exist or Error password")

	ErrorUserNotLogin = errors.New("User doesn't login")
	ErrorRegister     = errors.New("User registration failed")

	// 视频
	ErrorVideoExist = errors.New("Video already exist")

	ErrorInvalidVideoFormat = errors.New("Invalid Video Format")

	ErrorGenPicture = errors.New("Picture generate error")
)
