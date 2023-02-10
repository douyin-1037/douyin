package dal

import (
	"context"
	config "douyin/common/conf"
	"fmt"

	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func TestCreateUser(t *testing.T) {
	testInit()
	userName := "cdx9"
	encPassword := "123456"
	userID, err := CreateUser(context.Background(), userName, encPassword)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userID)
}

func TestIsFollowByID(t *testing.T) {
	testInit()
	var fanID int64 = 27
	var userID int64 = 26
	isfollowed, err := IsFollowByID(context.Background(), fanID, userID)
	if err != nil {
		fmt.Println("********", err)
	}
	fmt.Println(isfollowed, " ", fanID, " follow ", userID)
}

func TestFollowUser(t *testing.T) {
	testInit()
	var fanID int64 = 34
	var userID int64 = 36
	err := FollowUser(context.Background(), fanID, userID)
	if err != nil {
		fmt.Println("********", err)
	}
	fmt.Println(fanID, " follow ", userID)
}

func TestUnFollowUser(t *testing.T) {
	testInit()
	var fanID int64 = 27
	var userID int64 = 26
	err := UnFollowUser(context.Background(), fanID, userID)
	if err != nil {
		fmt.Println("********", err)
	}
	fmt.Println(fanID, " unfollow ", userID)
}

func TestGetFanList(t *testing.T) {
	testInit()
	var userID int64 = 34
	list, err := GetFanList(context.Background(), userID)
	if err != nil {
		fmt.Println("********", err)
	}
	fmt.Println(list)
}

func TestGetFollowList(t *testing.T) {
	testInit()
	var userID int64 = 34
	list, err := GetFollowList(context.Background(), userID)
	if err != nil {
		fmt.Println("********", err)
	}
	fmt.Println(list)
}

func TestGetFriendList(t *testing.T) {
	testInit()
	var userID int64 = 34
	list, err := GetFriendList(context.Background(), userID)
	if err != nil {
		fmt.Println("********", err)
	}
	fmt.Println(list)
}

/*
func TestIdxInout(t *testing.T) {
	// 生成Context
	ctx, err := GetContext()
	if err != nil {
		t.Errorf("get context error: %s\n", err.Error())
	} else {
		t.Logf("%+v", ctx)
	}
}
*/
