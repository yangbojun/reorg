package config

import "flag"

func GetServerPort() int {
	var port int
	flag.IntVar(&port, "p", 5849, "服务监听端口")
	flag.Parse()
	return port
}
