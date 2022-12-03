// Package service 存放中间件
package service

import (
	"message-board/dao"
	"message-board/model"
)

func CreateUser(u model.User) error {
	err := dao.InsertUser(u) //连向数据库
	return err
}

func SearchUserByUserName(name string) (user model.User, err error) { //按名字查找用户
	user, err = dao.SearchUserByUserName(name)
	return user, err
}

func ChangePassword(name string, NewPassword string) (err error) { //修改密码
	err = dao.ChangePassword(name, NewPassword)
	return err
}
