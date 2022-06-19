package router

import (
	"dousheng/controller"
	"dousheng/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	//// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	//hello test
	apiRouter.GET("/hello/", controller.Hello)

	//// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", middleware.JWTAuth(), controller.User)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", middleware.JWTAuth(), controller.Publish)
	apiRouter.GET("/publish/list/", middleware.JWTAuth(), controller.PublishList)
	//
	//// extra apis - I
	apiRouter.POST("/favorite/action/", middleware.JWTAuth(), controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", middleware.JWTAuth(), controller.FavoriteList)
	apiRouter.POST("/comment/action/", middleware.JWTAuth(), controller.CommentAction)
	apiRouter.GET("/comment/list/", middleware.JWTAuth(), controller.CommentList)
	//
	//// extra apis - II
	apiRouter.POST("/relation/action/", middleware.JWTAuth(), controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", middleware.JWTAuth(), controller.FollowList)
	apiRouter.GET("/relation/follower/list/", middleware.JWTAuth(), controller.FollowerList)
}
