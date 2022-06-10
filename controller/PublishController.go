package controller

import (
	"dousheng/common"
	"dousheng/service"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	idworker "github.com/gitstliu/go-id-worker"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"golang.org/x/net/context"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
)

type PublishRequest struct {
	Token string `form:"token" json:"token" binding:"required"`
	Data  []byte `form:"data" json:"data" binding:"required"`
	Title string `form:"title" json:"title" binding:"required"`
}

type PublishResponse struct {
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

var videoFileExt = []string{"mp4", "flv"} //此处可根据需要添加格式
var idGen *idworker.IdWorker

func init() {
	idGen = &idworker.IdWorker{}
	idGen.InitIdWorker(1, 1)
}

func Publish(c *gin.Context) {
	file, errd := c.FormFile("data")
	errd, coverUrl, playUrl := UploadVideo(file)
	if errd != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: common.ParamInvalid,
			StatusMsg:  "Publish Parameter parsing error",
		})
		return
	}
	token, _ := c.GetPostForm("token")
	strs := strings.Split(token, ":")
	username := strs[0]
	title, _ := c.GetPostForm("title")
	u, _ := service.GetUserByUsername(username)
	video := Video{
		Id:            u.ID,
		Author:        *u,
		PlayUrl:       "rd5met9ed.hn-bkt.clouddn.com" + "/" + playUrl,
		CoverUrl:      "rd5met9ed.hn-bkt.clouddn.com" + "/" + coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         title,
	}
	fmt.Println("video=%#v\n", video)
	c.JSON(http.StatusOK, common.Response{
		StatusCode: common.OK,
		StatusMsg:  "Publish Success!",
	})
}

func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, PublishListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}

func IsVideoAllowed(suffix string) bool {
	for _, fileExt := range videoFileExt {
		if suffix == fileExt {
			return true
		}
	}
	return false
}

func UploadVideo(file *multipart.FileHeader) (err error, coverUrl string, playUrl string) {
	//先处理输入
	filename := file.Filename                      //获取文件名
	indexOfDot := strings.LastIndex(filename, ".") //获取文件后缀名前的.的位置
	if indexOfDot < 0 {
		return errors.New("没有获取到文件的后缀名"), coverUrl, playUrl
	}
	suffix := filename[indexOfDot+1 : len(filename)] //获取后缀名
	suffix = strings.ToLower(suffix)                 //后缀名统一小写处理
	if !IsVideoAllowed(suffix) {
		return errors.New("上传的文件不符合视频的格式"), coverUrl, playUrl
	}
	fmt.Println("刚才上传的文件后缀名：" + suffix)
	id, err := idGen.NextId()
	filename = strconv.FormatInt(id, 10)

	videoName := filename + "." + suffix //视频的文件名
	data, err := file.Open()             //data是文件内容的访问接口（重点）
	VideoFolderName := "video"
	key := VideoFolderName + "/" + videoName //key是要上传的文件访问路径
	//下面是七牛api
	//domainName := "rd5met9ed.hn-bkt.clouddn.com"
	bucket := "top-20"
	accessKey := "ANvRMQN-FX6C6abeKAYxqAq1qq9je2x1UAmlLjFA"
	secretKey := "RhH86hgmwDphJxs5jBa1yUzZM7ydAch7msd-_VSi"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	//视频封面start
	coverName := filename + "." + "jpg"           //封面的文件名
	coverFolderName := "cover"                    //七牛云中存放图片的目录名。用于与文件名拼接，组成文件路径
	coverKey := coverFolderName + "/" + coverName //封面的访问路径，我们通过此路径在七牛云空间中定位封面
	saveJpgEntry := base64.StdEncoding.EncodeToString([]byte(bucket + ":" + coverKey))
	putPolicy.PersistentOps = "vframe/jpg/offset/1|saveas/" + saveJpgEntry //取视频第1秒的截图
	//end
	putPolicy.Expires = 7200 //自定义凭证有效期（示例2小时，Expires 单位为秒，为上传凭证的有效时间）
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	//data是字节流，data := []byte("hello, this is qiniu cloud")
	//file.size是要上传的文件大小
	err = formUploader.Put(context.Background(), &ret, upToken, key, data, file.Size, &putExtra)
	if err != nil {
		return err, coverUrl, playUrl
	}
	//fmt.Println(ret.Key, ret.Hash)

	coverUrl = coverKey
	playUrl = key
	return err, coverUrl, playUrl
}
