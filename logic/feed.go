package logic

import (
	"groupwork/module"
	"groupwork/repository/mysql"
)

func Feed() (*[]module.Video, error) {
	return mysql.FeedVideos()
}
