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
	var userID int64 = 4
	var videoID int64 = 7
	content1 := "test1111"
	content2 := "test2222"
	err := CreateComment(context.Background(), userID, videoID, content1)
	if err != nil {
		fmt.Println(err)
	}
	err = CreateComment(context.Background(), userID, videoID, content2)
	if err != nil {
		fmt.Println(err)
	}

	comments, err := GetCommentList(context.Background(), videoID)
	if err != nil {
		fmt.Println(err)
	}
	for _, comment := range comments {
		fmt.Println("commentID: ", comment.ID)
		fmt.Println("userID: ", comment.UserId)
		fmt.Println("videoID: ", comment.VideoId)
		fmt.Println("content: ", comment.Contents)
	}

	err = DeleteComment(context.Background(), int64(comments[0].ID))
	if err != nil {
		fmt.Println(err)
	}

	comments, err = GetCommentList(context.Background(), videoID)
	if err != nil {
		fmt.Println(err)
	}
	for _, comment := range comments {
		fmt.Println("commentID: ", comment.ID)
		fmt.Println("userID: ", comment.UserId)
		fmt.Println("videoID: ", comment.VideoId)
		fmt.Println("content: ", comment.Contents)
	}
}

//func TestCreateComment(t *testing.T) {
//	testInit()
//	var userID int64 = 1
//	var videoID int64 = 666
//	content := "test"
//	err := CreateComment(context.Background(), userID, videoID, content)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
//func TestDeleteComment(t *testing.T) {
//	testInit()
//	var commentID int64 = 1
//	err := DeleteComment(context.Background(), commentID)
//	if err != nil {
//		fmt.Println(err)
//	}
//}
//func TestGetCommentList(t *testing.T) {
//	testInit()
//	var videoID int64 = 666
//	comments, err := GetCommentList(context.Background(), videoID)
//	if err != nil {
//		fmt.Println(err)
//	}
//	for _, comment := range comments {
//		fmt.Println("commentID: ", comment.ID)
//		fmt.Println("userID: ", comment.UserId)
//		fmt.Println("videoID: ", comment.VideoId)
//		fmt.Println("content: ", comment.Contents)
//	}
//}