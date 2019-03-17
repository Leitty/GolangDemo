package upload

import (
	"Gin/learnGin/golangDemo/pkg/file"
	"Gin/learnGin/golangDemo/pkg/setting"
	"Gin/learnGin/golangDemo/pkg/util"
	"fmt"
	logger "github.com/gpmgo/gopm/modules/log"
	"log"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

func GetImageFullUrl(name string) string {
	return setting.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}

func GetImageName(name string) string {
	ext := path.Ext(name)
	filename := strings.TrimSuffix(name, ext)
	filename = util.EncodeMD5(filename)
	return filename + ext
}

func GetImagePath() string {
	return setting.AppSetting.ImageSavePath
}

func GetImageFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

func CheckImageExt(filename string) bool {
	ext := file.GetExt(filename)
	for _, allowExt := range setting.AppSetting.ImageAllowExts{
		if strings.ToUpper(allowExt) == strings.ToUpper(ext){
			return true
		}
	}

	return false
}

func CheckImageSize(f multipart.File) bool {
	size ,err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logger.Warn("Check image with error: %v", err)
		return false
	}
	return size <= setting.AppSetting.ImageMaxSize
}

func CheckImage(src string) error {
	dir ,err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
