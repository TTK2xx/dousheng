package controller

import (
	"dousheng/common"
	"github.com/gin-gonic/gin"
	idworker "github.com/gitstliu/go-id-worker"
	"mime/multipart"
	"net/http"
)

type PublishRequest struct {
	Token string `form:"token" json:"token" binding:"required"`
	Data  byte   `form:"data" json:"data" binding:"required"`
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

func IsVideoAllowed(suffix string) bool {
	for _, fileExt := range videoFileExt {
		if suffix == fileExt {
			return true
		}
	}
	return false
}

func UploadVideo(file *multipart.FileHeader) (err error) {
	//filename := file.Filename                      //获取文件名
	//indexOfDot := strings.LastIndex(filename, ".") //获取文件后缀名前的.的位置
	//if indexOfDot < 0 {
	//	return errors.New("没有获取到文件的后缀名")
	//}
	//suffix := filename[indexOfDot+1 : len(filename)] //获取后缀名
	//suffix = strings.ToLower(suffix)                 //后缀名统一小写处理
	//if !IsVideoAllowed(suffix) {
	//	return errors.New("上传的文件不符合视频的格式")
	//}
	//fmt.Println("刚才上传的文件后缀名：" + suffix)
	//id, err := idGen.NextId()
	//filename = strconv.FormatInt(id, 10)
	//filename = filename + "." + suffix
	//data, err := file.Open()
	//folderName := "video"
	//key := folderName + "/" + filename
	return err
}
func Publish(c *gin.Context) {
	file, err := c.FormFile("data")
	err = UploadVideo(file)
	if err != nil {
		c.JSON(http.StatusOK, common.Response{
			StatusCode: common.ParamInvalid,
			StatusMsg:  "Publish Parameter parsing error",
		})
		return
	}
	//var request PublishRequest
	//if err := c.ShouldBind(&request); err != nil {
	//	c.JSON(http.StatusOK, common.Response{
	//		StatusCode: common.ParamInvalid,
	//		StatusMsg:  "Publish Parameter parsing error",
	//	})
	//	return
	//}
}

func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, PublishListResponse{
		Response: common.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
