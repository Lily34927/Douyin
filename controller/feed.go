package controller

import (
	"github.com/gin-gonic/gin"
	"groupwork/logic"
	"groupwork/module"
	"net/http"
	"time"
)

type FeedResponse struct {
	module.Response
	VideoList []module.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// 1.业务处理
	data, err := logic.Feed()
	if err != nil {
		c.JSON(http.StatusOK, module.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	// 2.返回响应
	c.JSON(http.StatusOK, FeedResponse{
		Response:  module.Response{StatusCode: 0},
		VideoList: *data,
		NextTime:  time.Now().Unix(),
	})
}
