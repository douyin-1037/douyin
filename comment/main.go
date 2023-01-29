package main

import (
	commentproto "douyin/code_gen/kitex_gen/commentproto/commentservice"
	"log"
)

func main() {
	svr := commentproto.NewServer(new(CommentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
