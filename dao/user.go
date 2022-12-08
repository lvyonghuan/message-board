package dao

import (
	"message-board/model"
)

func InsertUser(u model.User) (err error) { //注册，将用户信息填入数据库
	_, err = DB.Exec("insert into user(username,password,question,secrecy,administrator,cookie) values (?,?,?,?,?,?)", u.UserName, u.Password, u.SecrecyQuestion, u.Secrecy, u.Administrator, 0)
	return err
}

func SearchUserByUserName(name string) (u model.User, err error) { //查找重复用户名
	row := DB.QueryRow("select * from user where username = ?", name)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.UserName, &u.Password, &u.SecrecyQuestion, &u.Secrecy, &u.Administrator, &u.Cookie)
	return
}

func ChangePassword(name string, NewPassword string) (err error) {
	_, err = DB.Exec("update user set password=? where username=?", NewPassword, name)
	return err
}

func InsertCookieToken(username string, token string) (err error) {
	_, err = DB.Exec("update user set cookie=? where username=?", token, username)
	return err
}

func DeleteCookie(username string) (err error) {
	_, err = DB.Exec("update user set  cookie=0 where username=?", username)
	return err
}

func SearchUsernameByCookie(cookie string) (username string, err error) {
	var u model.User
	row := DB.QueryRow("select * from user where cookie = ?", cookie)
	if err = row.Err(); row.Err() != nil {
		return
	}
	err = row.Scan(&u.UserName, &u.Password, &u.SecrecyQuestion, &u.Secrecy, &u.Administrator, &u.Cookie)
	return u.UserName, err
}
