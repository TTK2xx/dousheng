package controller

import (
	"dousheng/common"
	"dousheng/model"
	"dousheng/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*func CommentAction(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}*/

type CommentListRequest struct {
	Token   string `form:"token" json:"token" binding:"required"`
	VideoID int64  `form:"video_id" json:"video_id" binding:"required"`
}

type CommentListResponse struct {
	common.Response
	CommentList model.Comment `json:"comment_list"`
}

func CommentList(c *gin.Context) {
	var request CommentListRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: common.ParamInvalid,
			StatusMsg:  "Parameter parsing error",
		})
		return
	}

	// 用户鉴权ID处理
	/*
	 * 待处理
	 * 待处理
	 */

	//

	vid := request.VideoID
	// 参数验证

	// 检查视频ID是否存在
	/*
	 * 待处理
	 * 待处理
	 */

	//
	newCommentList := service.GetCommentByVideoID(vid)
	c.JSON(http.StatusOK, CommentListResponse{
		Response: common.Response{
			StatusCode: common.OK,
			StatusMsg:  "success",
		},
		CommentList: newCommentList,
	})
}
