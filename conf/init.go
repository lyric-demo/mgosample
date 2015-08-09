package conf

type (
	// App 应用程序配置
	App struct {
		Name  string // 应用程序名称
		Port  uint   // 监听端口
		Debug bool   // 是否开启调试模式
	}

	// Db 数据库配置
	Db struct {
		Host     string
		Port     uint
		User     string
		Pwd      string
		Database string
	}
)

var (
	// APPCONFIG 应用程序配置
	APPCONFIG *App
	// DBCONFIG 数据库配置
	DBCONFIG *Db
)

func init() {
	APPCONFIG = &App{
		Name:  "Blog Manage",
		Port:  8006,
		Debug: true,
	}

	DBCONFIG = &Db{
		Host:     "192.168.59.103",
		Port:     27017,
		Database: "blog",
	}
}
