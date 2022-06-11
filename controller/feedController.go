package controller

import (
	"dousheng/common"
	"dousheng/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FeedResponse struct {
	common.Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, PublishListResponse{
		Response: common.Response{
			StatusCode: common.OK,
			StatusMsg:  "",
		},
		VideoList: DemoVideos,
	})
	//videos := service.GetAllVideos()
	//c.JSON(http.StatusOK, FeedResponse{
	//	Response: common.Response{
	//		StatusCode: 0,
	//	},
	//	VideoList: videos,
	//	NextTime:  time.Now().Unix(),
	//})
}
