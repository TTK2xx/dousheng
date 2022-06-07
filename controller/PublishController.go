package controller

import (
	"dousheng/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PublishActionRequest struct {
	Token string `form:"token" json:"token" binding:"required"`
	Data  byte   `form:"data" json:"data" binding:"required"`
	Title string `form:"title" json:"title" binding:"required"`
}

type PublishActionResponse struct {
	common.Response
}

type PublishListRequest struct {
	Token string `form:"token" json:"token" binding:"required"`
	Title string `form:"title" json:"title" binding:"required"`
}

type PublishListResponse struct {
	common.Response
	VideoList []Video `json:"video_list"`
}

func Publish(c *gin.Context) {
	var request PublishActionRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: common.ParamInvalid,
			StatusMsg:  "Parameter parsing error",
		})
		return
	}
	//token := request.Token
	//strs := strings.Split(token, ":")
	//username := strs[0]
	//user, err := service.GetUserByUsername(username)
	//
	//if err != nil {
	//	c.JSON(http.StatusOK, common.Response{
	//		StatusCode: 1,
	//		StatusMsg:  err.Error(),
	//	})
	//	return
	//}
	//
	//c.JSON(http.StatusOK, common.Response{
	//	StatusCode: 0,
	//	StatusMsg:  " uploaded successfully",
	//})
}

func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, PublishListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
