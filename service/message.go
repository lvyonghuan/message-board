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

func LookAllMessage() (u []model.Message, err error) {
	u, err = dao.LookAllMessage()
	return
}

func DeleteMessage(username string, id int) (err error) {
	err = dao.DeleteMessage(username, id)
	return
}

func SearchUsernameByCookie(cookie string) (username string, err error) {
	username, err = dao.SearchUsernameByCookie(cookie)
	return username, err
}
