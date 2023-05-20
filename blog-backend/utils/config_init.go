package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

/**
* 模块名: config_init
* 代码描述: 初始化配置信息读取
* 作者:lizibin
* 创建时间:2022/12/23 14:39:52
 */
func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
}
