package dal

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

func TestCreateUser(t *testing.T) {
	testInit()
	userName := "lyy"
	encPassword := "123456"
	userID, err := CreateUser(context.Background(), userName, encPassword)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userID)
}
