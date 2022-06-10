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

func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, PublishListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
