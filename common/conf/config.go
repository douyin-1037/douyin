package conf

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"os"
	"path/filepath"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
)

var loadRemoteConfigFlag = flag.Bool("remote", false, "是否从远程apollo读取配置")

func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=Local",
		d.UserName,
		d.Password,
		d.Host,
		d.DBName,
		d.Charset,
		d.ParseTime,
	)
}

type JWTConfig struct {
	Secret  string
	Expires time.Duration
}

func InitConfig() {
	parseParam()

	vp := viper.New()

	if !*loadRemoteConfigFlag {
		workDirectory, err := os.Getwd()
		if err != nil {
			klog.Fatal(err)
		}
		sep := string(filepath.Separator)
		vp.AddConfigPath(workDirectory + sep + "conf")
		for filepath.Base(workDirectory) != "douyin" {
			vp.AddConfigPath(workDirectory + sep + "conf")
			workDirectory = filepath.Dir(workDirectory)
		}
		vp.AddConfigPath(workDirectory + sep + "conf")
		vp.SetConfigName("conf")
		vp.SetConfigType("yaml")
		if err := vp.ReadInConfig(); err != nil {
			klog.Fatal(err)
		}
	} else {
		//Read configuration from apollo configuration center
		c := &config.AppConfig{
			AppID:          "douyin",
			Cluster:        "dev",
			IP:             "http://127.0.0.1:8080",
			NamespaceName:  "application",
			IsBackupConfig: false,
		}
		client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
			return c, nil
		})
		klog.Info("Initializing Apollo configuration successfully")
		cache := client.GetConfigCache(c.NamespaceName)
		confValue, _ := cache.Get("conf.yaml")
		confString := fmt.Sprint(confValue)

		vp.SetConfigType("yaml")
		err := vp.ReadConfig(bytes.NewBuffer([]byte(confString)))
		if err != nil {
			fmt.Println(err)
		}
	}

	vp.UnmarshalKey("Server", &Server)
	vp.UnmarshalKey("Database", &Database)
	vp.UnmarshalKey("JWT", &JWT)
	vp.UnmarshalKey("COS", &COS)
	vp.UnmarshalKey("Redis", &Redis)
	vp.UnmarshalKey("Pulsar", &Pulsar)
	vp.UnmarshalKey("MongoDB", &MongoDB)
	JWT.Expires *= time.Hour
	Redis.ExpireTime *= 3600
	Redis.MaxRandAddTime *= 3600
}

func parseParam() {
	flag.Parse()
}
