package config

import (
	"fmt"

	"gopkg.in/ini.v1"
)

// 定义变量

var (
	// 数据库相关配置
	Db, DbHost, DbUser, DbPort, DbPassword, DbName string
	// 服务器配置
	AppMode, HttpPort, JwtKey string
)

// 初始化配置
func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		// 读取配置文件出错，直接退出程序
		panic(fmt.Sprintf("Fail to read file: %v", err))
	}
	loadDbConfig(file)
	loadServerConfig(file)

}

// 获取数据库配置
func loadDbConfig(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbPassword = file.Section("database").Key("DbPassWord").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("Musci")
}

// 获取服务器配置
func loadServerConfig(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("8080")
	JwtKey = file.Section("server").Key("JwtKey").MustString("123456")
}
