package dal

// @path: comment/infra/dal/orm_test.go
// @description: (immature)test of orm in comment
// @auth: wan-nan <wan_nan@foxmail.com>
import (
	"context"
	config "douyin/common/conf"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

func testInit() {
	config.InitConfig()
	DSN := config.Database.DSN()
	var err error
	DB, err = gorm.Open(mysql.Open(DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	DB = DB.Debug()
	if err != nil {
		klog.Fatal(err)
	}
}

func TestCommentGorm(t *testing.T) {
	testInit()
	var userId int64 = 22
	var videoID int64 = 7
	content1 := "2023年2月13日13:59:41"
	//content2 := "test2.10.2"
	var commentUUID int64 = 61
	createTime := time.Now().Unix()
	_, err := CreateComment(context.Background(), userId, videoID, content1, commentUUID, createTime)
	if err != nil {
		fmt.Println(err)
	}
	//_, err = CreateComment(context.Background(), userId, videoID, content2)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//comments, err := GetCommentList(context.Background(), videoID)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, comment := range comments {
	//	fmt.Println("commentID: ", comment.ID)
	//	fmt.Println("userId: ", comment.UserId)
	//	fmt.Println("videoID: ", comment.VideoId)
	//	fmt.Println("content: ", comment.Contents)
	//	fmt.Println("publish time: ", time.Unix(comment.CreatedAt.Unix(), 0).Format("2006-01-02 15:04:05")+"\n")
	//}
	//
	//err = DeleteComment(context.Background(), int64(comments[4].ID), comments[4].VideoId)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//comments, err = GetCommentList(context.Background(), videoID)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, comment := range comments {
	//	fmt.Println("commentID: ", comment.ID)
	//	fmt.Println("userId: ", comment.UserId)
	//	fmt.Println("videoID: ", comment.VideoId)
	//	fmt.Println("content: ", comment.Contents)
	//	fmt.Println("publish time: ", time.Unix(comment.CreatedAt.Unix(), 0).Format("2006-01-02 15:04:05")+"\n")
	//}
}

//	func TestCreateComment(t *testing.T) {
//		testInit()
//		var userId int64 = 1
//		var videoID int64 = 666
//		content := "test"
//		err := CreateComment(context.Background(), userId, videoID, content)
//		if err != nil {
//			fmt.Println(err)
//		}
//	}
func TestDeleteComment(t *testing.T) {
	testInit()
	var commentID int64 = 1243
	err := DeleteComment(context.Background(), commentID, 23)
	if err != nil {
		fmt.Println(err)
	}
}
func TestGetCommentList(t *testing.T) {
	testInit()
	var videoID int64 = 3
	comments, err := GetCommentList(context.Background(), videoID)
	if err != nil {
		fmt.Println(err)
	}
	for _, comment := range comments {
		fmt.Println("commentID: ", comment.ID)
		fmt.Println("userId: ", comment.UserId)
		fmt.Println("videoID: ", comment.VideoId)
		fmt.Println("content: ", comment.Contents)
		fmt.Println("publish time: ", time.Unix(comment.CreatedAt.Unix(), 0).Format("2006-01-02 15:04:05")+"\n")
	}
}
