package service

import (
	"message-board/dao"
	"message-board/model"
)

func SendMessage(u model.Message) (err error) {
	err = dao.SendMessage(u)
	return err
}

func CheckMessage(username string) (u []model.Message, err error) {
	u, err = dao.CheckMessage(username)
	return
}
