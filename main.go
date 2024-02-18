package main

import (
	"log"
	"palworldServer/daemon"
	"palworldServer/upgrade"
	"time"
)

const limit float32 = 0.95

func run() {
	palWorldServerPath := "PalServer.exe"
	for {
		log.Println("正在检查更新!!!")
		err := upgrade.Upgrade()
		if err != nil {
			log.Println("检查更新失败！！！！")
			panic(err)
		}

		time.Sleep(time.Second * 3)

		log.Println("开始启动帕鲁服务端!!!")

		time.Sleep(time.Second * 3)

		palworldServerPid, err := daemon.StartPalwordServer(palWorldServerPath)
		if err != nil {
			log.Println("启动帕鲁服务端异常!!")
			panic(err)
		}
		log.Println("帕鲁服务端启动成功!!!")

		time.Sleep(time.Second * 3)

		err = daemon.ListenProcess(palworldServerPid, limit)
		if err != nil {
			log.Println("重启帕鲁服务器异常！！！")
			panic(err)
		}
	}
}

func main() {

	run()
}
