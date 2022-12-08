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
	message := r.Group("/message")
	{
		message.POST("/send", Send)            //发送消息
		message.GET("/check", Check)           //查看提到登陆者的消息
		message.GET("/look", LookAllMessage)   //查看所有消息
		message.POST("/delete", DeleteMessage) //删除消息
	}
	r.Run()
}
