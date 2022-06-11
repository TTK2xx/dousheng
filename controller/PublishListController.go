package controller

import (
	"dousheng/common"
	"dousheng/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PublishListRequest struct {
	Token string `form:"token" json:"token" binding:"required"`
	Title string `form:"title" json:"title" binding:"required"`
}

type PublishListResponse struct {
	common.Response
	VideoList []model.Video `json:"video_list"`
}

func PublishList(c *gin.Context) { //我发布的视频列表
	//videos := service.GetAllVideos()
	//c.JSON(http.StatusOK, FeedResponse{
	//	Response: common.Response{
	//		StatusCode: common.OK,
	//		StatusMsg:  "Publish Success!",
	//	},
	//	VideoList: videos,
	//})
	c.JSON(http.StatusOK, PublishListResponse{
		Response: common.Response{
			StatusCode: common.OK,
			StatusMsg:  "",
		},
		VideoList: DemoVideos,
	})
}
