package conf

type (
	// app 配置
	App struct {
		Name  string // 应用程序名称
		Port  uint   // 监听端口
		Debug bool   // 是否开启调试模式
	}

	// db 配置
	Db struct {
		Host     string
		Port     uint
		User     string
		Pwd      string
		Database string
	}
)

var (
	APP_CONFIG *App
	DB_CONFIG  *Db
)

func init() {
	APP_CONFIG = &App{
		Name:  "Blog Manage",
		Port:  8006,
		Debug: true,
	}

	DB_CONFIG = &Db{
		Host:     "192.168.59.103",
		Port:     27017,
		Database: "blog",
	}
}
