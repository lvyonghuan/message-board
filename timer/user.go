package timer

import (
	"log"
	"message-board/service"
	"time"
)

func DeleteCookieTimer(username string) {
	timer := time.NewTimer(time.Duration(3600) * time.Second)
	<-timer.C //3600秒之前，管道阻塞。之后执行删除程序。
	err := service.DeleteCookie(username)
	if err != nil {
		log.Printf("search timer error:%v", err)
		return
	}
}
