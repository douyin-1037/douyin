package main

import (
	videoproto "douyin/code_gen/kitex_gen/videoproto/videoservice"
	"log"
)

func main() {
	svr := videoproto.NewServer(new(VideoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
