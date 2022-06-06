package controller

import (
	"dousheng/common"
	"dousheng/model"
	"dousheng/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type UserLoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UserLoginResponse struct {
	common.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	common.Response
	User UserInfo `json:"user"`
}

type UserInfo struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

func Register(c *gin.Context) {
	var request UserLoginRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: common.ParamInvalid,
			StatusMsg:  "Parameter parsing error",
		})
		return
	}
	token := request.Username + ":" + request.Password
	// 参数验证

	// 检查用户是否存在
	if exist := service.IsUserExisted(request.Username); exist {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: common.UserHasExisted,
			StatusMsg:  "User has existed",
		})
		return
	}
	//
	newUser := model.User{
		Username: request.Username,
		Password: request.Password,
	}
	service.CreateUser(&newUser)
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: common.Response{
			StatusCode: common.OK,
			StatusMsg:  "success",
		},
		UserId: newUser.ID,
		Token:  token,
	})
}

func Login(c *gin.Context) {
	var request UserLoginRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: common.ParamInvalid,
			StatusMsg:  "Parameter parsing error",
		})
		return
	}
	// 检查用户是否存在
	u, _ := service.GetUserByUsername(request.Username)
	if u.Username != request.Username {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: common.UserNotExisted,
			StatusMsg:  "User not existed",
		})
		return
	}
	// 登录状态保持 session? jwt?
	if u.Password == request.Password {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{
				StatusCode: common.OK,
				StatusMsg:  "success",
			},
			UserId: u.ID,
			Token:  u.Username + ":" + u.Password,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{
				StatusCode: common.WrongPassword,
				StatusMsg:  "Wrong password",
			},
		})
	}
}

func User(c *gin.Context) {
	token := c.Query("token")
	strs := strings.Split(token, ":")
	username := strs[0]
	u, _ := service.GetUserByUsername(username)
	if u.Username != username {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: common.UserNotExisted,
			StatusMsg:  "User not existed",
		})
		return
	}
	UserInfo := UserInfo{
		Id:            u.ID,
		Name:          u.Username,
		FollowCount:   99,
		FollowerCount: 66,
		IsFollow:      false,
	}
	c.JSON(http.StatusOK, UserResponse{
		Response: common.Response{
			StatusCode: common.OK,
		},
		User: UserInfo,
	})
}
