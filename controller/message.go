package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"groupwork/module"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"
)

var tempChat = map[string][]module.Message{}

var messageIdSequence = int64(1)

type ChatResponse struct {
	module.Response
	MessageList []module.Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")

	if user, exist := usersLoginInfo[token]; exist {
		userIdB, _ := strconv.Atoi(toUserId)
		chatKey := genChatKey(user.UserId, int64(userIdB))

		atomic.AddInt64(&messageIdSequence, 1)
		curMessage := module.Message{
			Id:         messageIdSequence,
			Content:    content,
			CreateTime: time.Now().Format(time.Kitchen),
		}

		if messages, exist := tempChat[chatKey]; exist {
			tempChat[chatKey] = append(messages, curMessage)
		} else {
			tempChat[chatKey] = []module.Message{curMessage}
		}
		c.JSON(http.StatusOK, module.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, module.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")

	if user, exist := usersLoginInfo[token]; exist {
		userIdB, _ := strconv.Atoi(toUserId)
		chatKey := genChatKey(user.UserId, int64(userIdB))

		c.JSON(http.StatusOK, ChatResponse{Response: module.Response{StatusCode: 0}, MessageList: tempChat[chatKey]})
	} else {
		c.JSON(http.StatusOK, module.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
