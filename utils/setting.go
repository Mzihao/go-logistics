package utils

import (
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")       // 要读取的文件名
	viper.SetConfigType("yml")               // 要读取的文件类型
	viper.AddConfigPath(workDir + "/config") // 要读取的文件路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
