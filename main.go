package main

import (
	"IM_QQ/chat"
	mysql "IM_QQ/dao"
	"IM_QQ/logger"
	"IM_QQ/router"
	"IM_QQ/settings"
	"fmt"
)

func main() {
	if err := settings.Init(); err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := logger.InitLogger(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := mysql.InitMysql(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("load mysql failed, err:%v\n", err)
		return
	}
	if err := chat.InitProducer(settings.Conf.MsgChannelType); err != nil {
		fmt.Printf("load producer failed, err:%v\n", err)
		return
	}
	defer chat.Close()
	if err := chat.InitConsumer(settings.Conf.MsgChannelType); err != nil {
		fmt.Printf("load consumer failed, err:%v\n", err)
		return
	}
	defer chat.CloseConsumer()
	go chat.MyServer.Start()

	r := router.InitRouter()
	r.Run(":9090")
}
