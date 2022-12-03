package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"message-board/model"
	"message-board/service"
	"message-board/util"
)

func Send(c *gin.Context) {
	cookie, err := c.Cookie("LoginState")
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	message := c.PostForm("massage")
	username := c.PostForm("username")              //这个username是指给谁发送信息
	_, err = service.SearchUserByUserName(username) //查找发送对象是否存在于数据库中
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	err = service.SendMessage(model.Message{ //发送消息
		Sendbywho: cookie,
		Message:   message,
		Username:  username,
	})
	if err != nil {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}

func Check(c *gin.Context) { //查看关于用户的消息
	cookie, err := c.Cookie("LoginState")
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	u, err := service.CheckMessage(cookie)
	if err != nil {
		util.RsepInternalErr(c)
	}
	c.JSON(200, u)
	util.RespOK(c)
}
