package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"time"
	"web_app/global"
	"web_app/utils"
)

func PutHeaderImage(c *gin.Context) {
	file, _ := c.FormFile("file")
	fileObj, err := file.Open()
	if err != nil {
		fmt.Println(err)
		return
	}

	//判断文件格式
	extString := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg": true,
		//".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extString]; !ok {
		utils.ResponseError(c, global.ImageFormatError)
		return
	}

	// 把文件上传到minio对应的桶中
	ok := utils.UploadFile("userheader", file.Filename, fileObj, file.Size)
	if !ok {

		utils.ResponseError(c, global.ImageUpdateError)
		return
	}
	headerUrl := utils.GetFileUrl("userheader", file.Filename, time.Second*24*60*60)
	if headerUrl == "" {

		utils.ResponseError(c, global.ImageGetError)
		//Response.Success(c, 400, "获取用户头像失败", "headerMap")
		return
	}
	//TODO 把用户的头像地址存入到对应user表中head_url 中
	utils.ResponseSuccess(c, headerUrl)
}
