package main

import (
	"fmt"
	D "go_api/lib/database"
	G "go_api/lib/global"
	"go_api/lib/response"
	"go_api/module/bihua"
	"go_api/module/chengyu"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	D.Init()

	r.Use(auth)
	r.GET("api/bihua", bihua.Get)
	r.GET("api/chengyu", chengyu.Get)

	r.Run(G.Config("server", "port"))
	fmt.Println("开启服务:" + G.Config("server", "port"))
}

func auth(c *gin.Context) {
	token := c.GetHeader("token")
	if len(token) <= 0 || token != "hanzi2020" {
		c.JSON(response.TokenInvalid.Code, G.Json(response.TokenInvalid.Message, nil))
		c.Abort()
		return
	}
}
