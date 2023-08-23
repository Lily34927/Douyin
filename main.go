package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"groupwork/module"
	"groupwork/repository/mysql"
	"groupwork/router"
)

func main() {
	// 配置mysql连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("connect failed, err: %v\n", err)
		return
	}
	defer mysql.Close()

	// 初始化snowflake算法
	if err := module.Init("2023-08-07", 1); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	//go service.RunMessageServer()

	r := gin.Default()

	router.InitRouter(r)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
