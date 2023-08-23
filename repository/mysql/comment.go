package mysql

import (
	"errors"
	"github.com/jinzhu/gorm"
	"groupwork/module"
)

func InsertComment(commentRecord *module.Comment) (err error) {
	err = DB.Create(commentRecord).Error
	return
}

func DeleteComment(commentRecord *module.Comment) (err error) {
	err = DB.Where("comment_id=?", commentRecord.CommentId).Delete(&module.Comment{}).Error
	return
}

func CommentsQuery(videoId int64) (commentlist *module.CommentList, err error) {
	commentlist = new(module.CommentList)
	// 评论按时间顺序逆置显示
	var comments []module.Comment
	err = DB.Where("video_id=?", videoId).Order("created_at desc").Find(&comments).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { // find未找到
		return nil, errors.New("Video doest's have comments")
	}
	commentlist.AllComments = comments
	return
}

func UpdateCommentCount(commentRecord *module.Comment) (err error) {
	if commentRecord.Action == 1 {
		return addCommentCount(commentRecord)
	}
	return minusCommentCount(commentRecord)
}

func addCommentCount(commentRecord *module.Comment) (err error) {
	video := findCommentCount(commentRecord.VideoId)
	return DB.Model(video).Update("comment_count", video.CommentCount+1).Error
	//DB.Model(voteVideo).Update("favorite_count", favoriteCount+1)
	//fmt.Printf("video.FavoriteCount=%d\n", voteVideo.FavoriteCount)
}

func minusCommentCount(commentRecord *module.Comment) (err error) {
	video := findCommentCount(commentRecord.VideoId)
	return DB.Model(video).Update("comment_count", video.CommentCount-1).Error
}

func findCommentCount(videoId int64) (video *module.Video) {
	video = new(module.Video)
	DB.Where("video_id=?", videoId).First(video)
	return
}
