package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径：", err)
	}

	LoanServer(file)
	LoanDatabase(file)
}

func LoanServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
}

func LoanDatabase(file *ini.File) {
	DbHost = file.Section("databese").Key("DbHost").MustString("localhost")
	DbPort = file.Section("databese").Key("DbPort").MustString("3306")
	DbUser = file.Section("databese").Key("DbUser").MustString("root")
	DbPassWord = file.Section("databese").Key("DbPassWord").MustString("123456")
	DbName = file.Section("databese").Key("DbName").MustString("ginblog")
}
