package api

import (
	"Gin/learnGin/golangDemo/models"
	"Gin/learnGin/golangDemo/pkg/e"
	"Gin/learnGin/golangDemo/pkg/logging"
	"Gin/learnGin/golangDemo/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `json:"Required; MaxSize(50)"`
	Password string `json:"Required; MaxSize(50)"`
}

// @Summary 生成token
// @Tags token
// @Produce  json
// @Param username query string true "Username"
// @Param password query string true "Password"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth [get]
func GetAuth(c *gin.Context){
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username:username, Password:password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]string)
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username,password)
		if isExist {
			token , err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
				logging.Info("Fail to Generate Token: %v",err)
				//log.Printf("Fail to Generate Token: %v",err)
			} else {
				data["token"] = token

				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
			//log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}