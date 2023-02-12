package dal

// @path: video/infra/dal/init.go
// @description: initialization of gorm.DB
// @author: Chongzhi <dczdcz2001@aliyun.com>
import (
	"douyin/common/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init DB
func Init() {
	conf.InitConfig()
	var err error
	DB, err = gorm.Open(mysql.Open(conf.Database.DSN()),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		klog.Fatal(err)
	}
}
