package conf

import (
	"github.com/joho/godotenv"
	"hotel_cms/cache"
	"hotel_cms/entity"
	"hotel_cms/util"
	"os"
)

func init()  {
	// 从本地读取环境变量
	err := godotenv.Load(".env")
	if err != nil {
		util.Log().Panic("未读到本地配置文件.env")
	}
}

// 初始化配置项
func Init() {


	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	// 连接数据库
	entity.Database(os.Getenv("MYSQL"))

	cache.Redis()
}
