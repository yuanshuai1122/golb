package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golb/golb-server/utils"
)

func main() {
	fmt.Println("hello world")
	// 初始化数据库
	utils.InitDB()
	r := gin.Default()

	r.Run(":8088")
}
