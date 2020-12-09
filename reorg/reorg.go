package main

import (
	"fmt"
	"net/http"
	"reorg/config"
	"reorg/routers"
	"reorg/mylog"
)

func init() {
	// 创建日志路径
	mylog.MkLogDir()
}

func main() {
	// 获取指定端口,默认5849
	addr := fmt.Sprintf(":%v", config.GetServerPort())

	// 注册服务地址
	routers.Register()

	// 监听服务
	http.ListenAndServe(addr, nil)
}
