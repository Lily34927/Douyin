package mysql

import (
	"errors"
	"github.com/jinzhu/gorm"
	"groupwork/module"
)

func ActionHandle(voteVideo *module.FavoriteVideo) (err error) {
	// 新建或删除赞操作的记录
	if voteVideo.Action == 1 {
		return createAction(voteVideo)
	}
	return deleteAction(voteVideo)

	// 投票结果记录到新的数据表中. 记录用户为该帖子投票的数据. 赞成票还是反对票
}

func GetFavoriteListByUserId(userId int64) (videolist *module.VideoList, err error) {
	// 思路：
	// 1.在favorite_videos表中，通过userId找到用户喜欢视频video_id;
	// 2.在表videos中，通过video_id获取封面，视频，赞数量等信息
	var videoFavorite []module.FavoriteVideo
	err = DB.Where("user_id=?", userId).Order("created_at desc").Find(&videoFavorite).Error
	if err != nil {
		return nil, errors.New("User doest's have favorite videos")
	}

	var videos []module.Video
	for _, data := range videoFavorite {
		var video module.Video
		err = DB.Where("video_id=?", data.VideoId).First(&video).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Query failed")
		}
		videos = append(videos, video)
	}

	videolist = new(module.VideoList)
	videolist.AllVideoes = videos
	return
}

func UpdateFavoriteCount(voteVideo *module.FavoriteVideo) (err error) {
	// 更新 video 类下的 FavoriteCount
	if voteVideo.Action == 1 {
		return addAction(voteVideo)
	}
	return minusAction(voteVideo)
}

// createAction 添加赞的视频
func createAction(voteVideo *module.FavoriteVideo) (err error) {
	return DB.Create(voteVideo).Error
}

// deleteAction 取消赞的视频
func deleteAction(voteVideo *module.FavoriteVideo) (err error) {
	return DB.Where("user_id=? and video_id=?", voteVideo.UserId, voteVideo.VideoId).
		Delete(&module.FavoriteVideo{}).Error
}

func addAction(voteVideo *module.FavoriteVideo) (err error) {
	video := findFavoriteCount(voteVideo.VideoId)
	return DB.Model(video).Update("favorite_count", video.FavoriteCount+1).Error
	//DB.Model(voteVideo).Update("favorite_count", favoriteCount+1)
	//fmt.Printf("video.FavoriteCount=%d\n", voteVideo.FavoriteCount)
}

func minusAction(voteVideo *module.FavoriteVideo) (err error) {
	video := findFavoriteCount(voteVideo.VideoId)
	return DB.Model(video).Update("favorite_count", video.FavoriteCount-1).Error
}

func findFavoriteCount(videoId int64) (video *module.Video) {
	video = new(module.Video)
	DB.Where("video_id=?", videoId).First(video)
	return
}
