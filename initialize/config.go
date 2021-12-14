package initialize


import (
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"ginjujin/config"
	"ginjujin/global"
)

func InitConfig() {
	// 实例化viper
	v := viper.New()

	// 文件的路径如何设置
	switch global.RunMode {
		case "dev":
			v.SetConfigFile("./settings-dev.yaml")
		case "sit":
			v.SetConfigFile("./settings-sit.yaml")
		case "pro":
			v.SetConfigFile("./settings-pro.yaml")
	}
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	// 给serverConfig初始值
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}

	// 传递给全局变量
	global.Settings = serverConfig
	color.Blue("开始配置初始化中", global.Settings.LogsAddress)
}