package controller

import (
	"dousheng/common"
	"dousheng/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	common.Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	//为方便测试注销该句
	//videos := service.GetAllVideos()
	c.JSON(http.StatusOK, FeedResponse{
		Response: common.Response{
			StatusCode: 0,
		},

		// 方便评论测试，改为了测试数据
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})

	c.JSON(http.StatusOK, PublishListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
