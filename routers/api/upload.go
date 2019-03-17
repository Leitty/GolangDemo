package api

import (
	"Gin/learnGin/golangDemo/pkg/e"
	"Gin/learnGin/golangDemo/pkg/upload"
	"github.com/gin-gonic/gin"
	"github.com/gpmgo/gopm/modules/log"
	"net/http"
)

func UploadImage(c *gin.Context){
	code := e.SUCCESS
	data := make(map[string]string)

	//发送报文的图片key要为image
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		log.Warn("Request from file err: %v", err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg": e.GetMsg(code),
			"data": data,
		})
	}

	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()

		src := fullPath+imageName
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file){
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				log.Warn("%v", err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				log.Warn("Fail to save image: %v", err)
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath+imageName
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}