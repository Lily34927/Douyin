package controller

import (
	"github.com/gin-gonic/gin"
	"groupwork/logic"
	"groupwork/module"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	// 1.获取参数
	// 记录 video_id 和 action_type
	voteVideo := new(module.FavoriteVideo)
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")

	voteVideo.VideoId, _ = strconv.ParseInt(videoId, 10, 64)
	action, _ := strconv.ParseInt(actionType, 10, 64)
	voteVideo.Action = int8(action)

	// 2.获取当前请求的用户的id
	userId, err := GetCurrentUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, module.UserResponse{
			Response: module.Response{StatusCode: 1, StatusMsg: "User need login"},
		})
		return
	}
	voteVideo.UserId = userId

	// 3.业务处理
	if err := logic.FavoriteAction(voteVideo); err != nil {
		c.JSON(http.StatusOK, module.Response{StatusCode: 1, StatusMsg: err.Error()})
	}

	// 4.返回响应
	c.JSON(http.StatusOK, module.Response{StatusCode: 0, StatusMsg: "FavoriteAction success"})
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	// 1.获取用户id
	user_id := c.Query("user_id")
	userId, _ := strconv.ParseInt(user_id, 10, 64)

	//2.业务处理
	data, err := logic.FavoriteList(userId)
	if err != nil {
		c.JSON(http.StatusOK, module.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 3.返回响应
	c.JSON(http.StatusOK, VideoListResponse{
		Response: module.Response{
			StatusCode: 0,
		},
		VideoList: data.AllVideoes,
	})
}
