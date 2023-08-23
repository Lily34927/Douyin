package router

import (
	"github.com/gin-gonic/gin"
	"groupwork/controller"
	"groupwork/module"
)

func InitRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)

	apiRouter.GET("/comment/list/", controller.CommentList)

	// 使用中间件 -- 查询用户的信息，并返回
	apiRouter.Use(module.AuthMiddleWare())
	{
		apiRouter.GET("/user/", controller.UserInfo)
		apiRouter.POST("/publish/action/", controller.Publish)
		apiRouter.GET("/publish/list/", controller.PublishList)

		// extra apis - I
		apiRouter.POST("/favorite/action/", controller.FavoriteAction)
		apiRouter.GET("/favorite/list/", controller.FavoriteList)
		apiRouter.POST("/comment/action/", controller.CommentAction)

		// extra apis - II
		apiRouter.POST("/relation/action/", controller.RelationAction)
		apiRouter.GET("/relation/follow/list/", controller.FollowList)
		apiRouter.GET("/relation/follower/list/", controller.FollowerList)
		apiRouter.GET("/relation/friend/list/", controller.FriendList)
		apiRouter.GET("/message/chat/", controller.MessageChat)
		apiRouter.POST("/message/action/", controller.MessageAction)
	}
}
