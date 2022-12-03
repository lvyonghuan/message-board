// Package api 路由注册
package api

import "github.com/gin-gonic/gin"

func InitRouter() {
	r := gin.Default()
	use := r.Group("/user")
	{
		use.POST("/register", Register) //注册
		use.POST("/login", Login)       //登录
		use.PUT("/change", Password)    //修改密码
	}
	mesg := r.Group("/message")
	{
		mesg.POST("/send", Send)
		mesg.GET("/check", Check)
	}
	r.Run()
}
