// Package api 用户的注册、登陆和修改密码实现
package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"message-board/model"
	"message-board/service"
	"message-board/util"
)

func Register(c *gin.Context) {
	Username := c.PostForm("username")
	Password := c.PostForm("password")
	SecrecyQuestion := c.PostForm("question")
	Secrecy := c.PostForm("secrecy")
	if Username == "" || Password == "" { //当用户名或密码为空时，返回错误
		util.RespParamErr(c)
		return
	}
	u, err := service.SearchUserByUserName(Username) //查重模块
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	if u.UserName != "" {
		util.NormErr(c, 300, "用户名已注册")
		return
	}
	err = service.CreateUser(model.User{ //创建用户信息
		UserName:        Username,
		Password:        Password,
		SecrecyQuestion: SecrecyQuestion,
		Secrecy:         Secrecy,
	})
	if err != nil {
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}

func Login(c *gin.Context) {
	Username := c.PostForm("username")
	Password := c.PostForm("password")
	u, err := service.SearchUserByUserName(Username) //查找用户名是否存在于数据库中,且把密码从数据库取出来
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	if u.UserName == "" {
		util.NormErr(c, 300, "用户未注册或用户名输入错误")
		return
	}
	if u.Password != Password {
		util.NormErr(c, 20001, "密码错误")
		return
	}
	c.SetCookie("LoginState", u.UserName, 3600, "/", "localhost", false, true) //保存登陆状态
	util.RespOK(c)
}

func Password(c *gin.Context) {
	Username := c.PostForm("username")
	u, err := service.SearchUserByUserName(Username) //查找用户名是否存在于数据库中
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	if u.UserName == "" {
		util.NormErr(c, 300, "用户未注册或用户名输入错误")
		return
	}
	if u.SecrecyQuestion == "" {
		util.NormErr(c, 20011, "用户未设置密保！")
		return
	}
	c.JSON(200, u.SecrecyQuestion) //密保问题
	Answer := c.PostForm("answer")
	if u.Secrecy != Answer {
		util.NormErr(c, 20012, "密保填写错误")
		return
	}
	NewPassword := c.PostForm("NewPassword")
	if NewPassword == "" { //新密码为空时，返回错误
		util.RespParamErr(c)
		return
	}
	err = service.ChangePassword(Username, NewPassword) //修改密码
	if err != nil && err != sql.ErrNoRows {
		log.Printf("search user error:%v", err)
		util.RsepInternalErr(c)
		return
	}
	util.RespOK(c)
}
