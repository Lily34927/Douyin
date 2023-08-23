package logic

import (
	"groupwork/module"
	"groupwork/repository/mysql"
)

func CommentAction(commentRecord *module.Comment) (err error) {
	// 1. 写入评论
	if commentRecord.Action == 1 {
		commentRecord.CommentId = module.GenID()
		err = mysql.InsertComment(commentRecord)
	}

	// 2. 删除评论
	if commentRecord.Action == 2 {
		err = mysql.DeleteComment(commentRecord)
		return
	}

	// 3.更新video的CommentCount
	return mysql.UpdateCommentCount(commentRecord)
}

func CommentList(videoId int64) (commentlist *module.CommentList, err error) {
	return mysql.CommentsQuery(videoId)
}
