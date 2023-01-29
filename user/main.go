package main

import (
	userproto "douyin/code_gen/kitex_gen/userproto/userservice"
	"log"
)

func main() {
	svr := userproto.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
