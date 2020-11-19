package main

import (
	"gin-template/router"
	_ "gorm.io/driver/mysql"
)

func main() {
	// 启动服务
	router.Server()
}
