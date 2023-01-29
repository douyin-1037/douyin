package main

import (
	messageproto "douyin/code_gen/kitex_gen/messageproto/messageservice"
	"log"
)

func main() {
	svr := messageproto.NewServer(new(MessageServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
