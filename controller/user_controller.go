package controller

import (
	"fmt"
	"goserver/controller/request"
	"goserver/controller/result"
	"goserver/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var request = request.NewRegisterRequest()
	c.BindJSON(request)
	service.Register(request)
	c.JSON(http.StatusOK, result.NewCommonResult(map[string]string{}))
}

func Login(c *gin.Context) {
	var request = request.NewLoginRequest()
	c.BindJSON(request)
	var res, _ = service.Login(request)
	c.JSON(http.StatusOK, result.NewCommonResult(res))
}

func DeleteUser(c *gin.Context) {

}

func UpdateUserInfo(c *gin.Context) {

}

func QueryUserInfo(c *gin.Context) {
	// 定义一个切片
	slice := []int{1, 2, 3}

	// 尝试访问切片的越界索引
	fmt.Println(slice[3])
}
