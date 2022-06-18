package controller

import (
	"dousheng/common"
	"dousheng/middleware"
	"dousheng/model"
	"dousheng/service"
	"dousheng/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
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
	User model.UserInfo `json:"user"`
}

func Register(ctx *gin.Context) {
	var request UserLoginRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusOK, common.Response{
			StatusCode: common.ParamInvalid,
			StatusMsg:  "Parameter parsing error",
		})
		return
	}
	// 参数验证

	// 检查用户是否存在
	if exist := service.IsUserExisted(request.Username); exist {
		ctx.JSON(http.StatusOK, common.Response{
			StatusCode: common.UserHasExisted,
			StatusMsg:  "User has existed",
		})
		return
	}
	hashPassword, err := utils.PasswordHash(request.Password)
	if err != nil {
		return
	}
	newUser := model.User{
		ID:       utils.GenID(),
		Username: request.Username,
		Password: hashPassword,
	}
	service.CreateUser(&newUser)
	token := generateToken(ctx, newUser)
	ctx.JSON(http.StatusOK, UserLoginResponse{
		Response: common.Response{
			StatusCode: common.OK,
			StatusMsg:  "success",
		},
		UserId: newUser.ID,
		Token:  token,
	})
}

func Login(ctx *gin.Context) {
	var request UserLoginRequest
	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusOK, common.Response{
			StatusCode: common.ParamInvalid,
			StatusMsg:  "Parameter parsing error",
		})
		return
	}
	// 检查用户是否存在
	u, _ := service.GetUserByUsername(request.Username)
	if u.Username != request.Username {
		ctx.JSON(http.StatusOK, common.Response{
			StatusCode: common.UserNotExisted,
			StatusMsg:  "User not existed",
		})
		return
	}
	// 登录状态保持 session? jwt?
	if utils.PasswordVerify(request.Password, u.Password) {
		token := generateToken(ctx, *u)
		ctx.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{
				StatusCode: common.OK,
				StatusMsg:  "success",
			},
			UserId: u.ID,
			Token:  token,
		})
	} else {
		ctx.JSON(http.StatusOK, UserLoginResponse{
			Response: common.Response{
				StatusCode: common.WrongPassword,
				StatusMsg:  "Wrong password",
			},
		})
	}
}

func User(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(*middleware.CustomClaims)
	username := claims.UserName
	u, _ := service.GetUserByUsername(username)
	if u.Username != username {
		ctx.JSON(http.StatusOK, common.Response{
			StatusCode: common.UserNotExisted,
			StatusMsg:  "User not existed",
		})
		return
	}
	if err, userInfo := service.GetUserInfoByUserID(u.ID, u.ID); err != nil {
		ctx.JSON(http.StatusOK, UserResponse{
			Response: common.Response{
				StatusCode: common.OperationFailed,
			},
			User: userInfo,
		})
	} else {
		ctx.JSON(http.StatusOK, UserResponse{
			Response: common.Response{
				StatusCode: common.OK,
			},
			User: userInfo,
		})
	}

}

// token生成器
// md 为上面定义好的middleware中间件
func generateToken(c *gin.Context, user model.User) string {
	// 构造SignKey: 签名和解签名需要使用一个值
	j := middleware.NewJWT()

	// 构造用户claims信息(负荷)
	claims := middleware.CustomClaims{
		user.Username,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 签名过期时间
			Issuer:    "ttk",                           // 签名颁发者
		},
	}

	// 根据claims生成token对象
	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: common.OperationFailed,
			StatusMsg:  err.Error(),
		})
	}

	log.Println(token)

	return token

}
