package dao

import (
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
