package controller

import (
	"github.com/gin-gonic/gin"
	"groupwork/logic"
	"groupwork/module"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	module.Response
	CommentList []module.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	module.Response
	Comment module.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	// 1.获取参数
	commentRecord := new(module.Comment)
	// 记录 video_id，action_type，comment_text和comment_id[删除用]
	video_id := c.Query("video_id")
	commentRecord.VideoId, _ = strconv.ParseInt(video_id, 10, 64)
	action_type := c.Query("action_type")
	actionType, _ := strconv.ParseInt(action_type, 10, 64)
	commentRecord.Action = int8(actionType)
	if commentRecord.Action == 1 {
		comment_text := c.Query("comment_text")
		commentRecord.Content = comment_text
	}
	if commentRecord.Action == 2 {
		comment_id := c.Query("comment_id")
		commentRecord.CommentId, _ = strconv.ParseInt(comment_id, 10, 64)
	}

	// 2.获取当前请求的用户的id
	userId, err := GetCurrentUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, module.Response{StatusCode: 1, StatusMsg: "User need login"})
		return
	}
	commentRecord.UserId = userId

	// 3.业务处理
	if err = logic.CommentAction(commentRecord); err != nil {
		c.JSON(http.StatusOK, module.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	// 4.返回响应
	if commentRecord.Action == 1 {
		c.JSON(http.StatusOK, CommentActionResponse{Response: module.Response{StatusCode: 0},
			Comment: module.Comment{
				CommentId: commentRecord.CommentId,
				UserId:    commentRecord.UserId,
				Content:   commentRecord.Content,
				//gorm.Model.CreatedAt: commentRecord.Model.CreatedAt,
				//CreateDate: "05-01",
			}})
	}
	c.JSON(http.StatusOK, module.Response{StatusCode: 0})
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	// 1.获取参数
	id := c.Query("video_id")
	videoId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, module.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	// 2.业务处理
	commentlist, err := logic.CommentList(videoId)
	if err != nil {
		c.JSON(http.StatusOK, module.Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}

	// 3.返回响应

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    module.Response{StatusCode: 0},
		CommentList: commentlist.AllComments,
		//CommentList: DemoComments,
	})
}
