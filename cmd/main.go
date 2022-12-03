package main

import (
	"message-board/api"
	"message-board/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
