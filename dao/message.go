package dao

import (
	"errors"
	"message-board/model"
)

func SendMessage(u model.Message) (err error) {
	_, err = DB.Exec("insert into message(username,message,sendbywho) values (?,?,?)", u.Username, u.Message, u.Sendbywho)
	return err
}

func CheckMessage(username string) (U []model.Message, err error) {
	row, err := DB.Query("select * from message where username=?", username)
	if err != nil {
		return
	}
	for row.Next() {
		var u model.Message
		err = row.Scan(&u.Username, &u.Message, &u.Sendbywho, &u.ID)
		if err != nil {
			return
		}
		U = append(U, u)
	}
	return
}

func LookAllMessage() (U []model.Message, err error) {
	row, err := DB.Query("select * from message")
	if err != nil {
		return
	}
	for row.Next() {
		var u model.Message
		err = row.Scan(&u.Username, &u.Message, &u.Sendbywho, &u.ID)
		if err != nil {
			return
		}
		U = append(U, u)
	}
	return
}

func DeleteMessage(username string, id int) (err error) {
	var u model.Message
	row := DB.QueryRow("select * from message where id=?", id)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.Username, &u.Message, &u.Sendbywho, &u.ID)
	if err != nil {
		return err
	}
	if u.Sendbywho != username {
		return errors.New("无操作权限")
	}
	_, err = DB.Exec("delete from message where id=?", id) //删除信息
	if err != nil {
		return err
	}
	return err
}
