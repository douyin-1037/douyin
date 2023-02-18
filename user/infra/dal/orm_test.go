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
	userName := "test_user_2"
	encPassword := "123456"
	userId, err := CreateUser(context.Background(), userName, encPassword)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userId)
}

func TestIsFollowByID(t *testing.T) {
	testInit()
	var fanID int64 = 27
	var userId int64 = 26
	isfollowed, err := IsFollowByID(context.Background(), fanID, userId)
	if err != nil {
		fmt.Println("********", err)
	}
	fmt.Println(isfollowed, " ", fanID, " follow ", userId)
}

func TestFollowUser(t *testing.T) {
	testInit()
	var fanID int64 = 3
	var userId int64 = 4
	err := FollowUser(context.Background(), fanID, userId)
	if err != nil {
		fmt.Println("********", err)
	}
	fmt.Println(fanID, " follow ", userId)
}

func TestUnFollowUser(t *testing.T) {
	testInit()
	var fanID int64 = 3
	var userId int64 = 4
	err := UnFollowUser(context.Background(), fanID, userId)
	if err != nil {
		fmt.Println("********", err)
	}
	fmt.Println(fanID, " unfollow ", userId)
}

func TestGetFanList(t *testing.T) {
	testInit()
	var userId int64 = 34
	list, err := GetFanList(context.Background(), userId)
	if err != nil {
		fmt.Println("********", err)
	}
	fmt.Println(list)
}

func TestGetFollowList(t *testing.T) {
	testInit()
	var userId int64 = 34
	list, err := GetFollowList(context.Background(), userId)
	if err != nil {
		fmt.Println("********", err)
	}
	fmt.Println(list)
}

func TestGetFriendList(t *testing.T) {
	testInit()
	var userId int64 = 34
	list, err := GetFriendList(context.Background(), userId)
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
