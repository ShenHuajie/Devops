package main

import (
	"Devops/model"
	"Devops/router"
)

func main() {

	// 初始化数据库
	model.InitDb()
	router.InitRouter()
}
