package controller

import (
	"dousheng/common"
	"dousheng/model"
	"dousheng/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CommentActionRequest struct {
	Token       string `form:"token" json:"token" binding:"required"`
	VideoID     int64  `form:"video_id" json:"video_id" binding:"required"`
	ActionType  int32  `form:"action_type" json:"action_type" binding:"required, oneof=1 2"`
	CommentText string `form:"comment_text" json:"comment_text" `
	CommentID   int64  `form:"comment_id" json:"comment_id" `
}

type CommentActionResponse struct {
	common.Response
	Comment model.Comment `json:"comment"`
}

func CommentAction(c *gin.Context) {
	var request CommentActionRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: common.ParamInvalid,
			StatusMsg:  "Parameter parsing error",
		})
		return
	}

	choose := request.ActionType
	if choose == 1 { // 发布评论
		newComment := model.Comment{
			VideoID:    request.VideoID,
			Content:    request.CommentText,
			CreateDate: time.Now().Format("01-02"),
			// 等待user信息获取函数
			//CommentUser: ,
		}

		service.CreateComment(&newComment)
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: common.Response{
				StatusCode: common.OK,
				StatusMsg:  "success",
			},
			Comment: newComment,
		})
	} else if choose == 2 { // 删除评论
		service.DeleteComment(request.CommentID)
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: common.Response{
				StatusCode: common.OK,
				StatusMsg:  "success",
			},
		})
	}

}

type CommentListRequest struct {
	Token   string `form:"token" json:"token" binding:"required"`
	VideoID int64  `form:"video_id" json:"video_id" binding:"required"`
}

type CommentListResponse struct {
	common.Response
	CommentList []model.Comment `json:"comment_list"`
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

	vid := request.VideoID

	newCommentList := service.GetCommentByVideoID(vid)
	c.JSON(http.StatusOK, CommentListResponse{
		Response: common.Response{
			StatusCode: common.OK,
			StatusMsg:  "success",
		},
		CommentList: newCommentList,
	})
}
