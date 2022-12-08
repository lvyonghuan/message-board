package service

import "message-board/dao"

func DeleteCookie(username string) (err error) {
	err = dao.DeleteCookie(username)
	return err
}
