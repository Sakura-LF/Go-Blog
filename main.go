package main

import (
	config "Go-Blog/init"
)

func main() {
	// 1.初始化配置
	config.Init()
	config.ZapInit()
	config.InitMysql()

	// 2. 迁移表结构
	config.Migrate()

	// 3.
}
