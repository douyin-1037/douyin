package dal

// @path: comment/infra/dal/init.go
// @description: initialization of gorm.DB
// @auth: wan-nan <wan_nan@foxmail.com>
import (
	config "douyin/common/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(config.Database.DSN()),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		klog.Fatal(err)
	}
}
