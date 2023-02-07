package main

//TODO: 连数据库检查一下代码能否正常操作 + 每个文件记得都写注释 + service层提交 + main.go（本文件）
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
