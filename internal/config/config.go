package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	MysqlConf MysqlAddr `json:"Mysql"`
}

type MysqlAddr struct {
	Username      string `json:"Username"`
	Password      string `json:"Password"`
	Host          string `json:"Host"`
	Port          string `json:"Port"`
	DatabasesName string `json:"DatabasesName"`
}
