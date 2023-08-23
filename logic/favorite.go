package logic

import (
	"errors"
	"groupwork/module"
	"groupwork/repository/mysql"
)

func FavoriteAction(voteVideo *module.FavoriteVideo) (err error) {
	// 1.投票参数的限制：要么是1，要么是2
	if !(voteVideo.Action == 1 || voteVideo.Action == 2) {
		return errors.New("Parameter out of bounds")
	}

	// 2.处理投票记录
	if err := mysql.ActionHandle(voteVideo); err != nil {
		return err
	}

	// 3.更新video的FavoriteCount
	return mysql.UpdateFavoriteCount(voteVideo)
}

func FavoriteList(userId int64) (*module.VideoList, error) {
	return mysql.GetFavoriteListByUserId(userId)
}
