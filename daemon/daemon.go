package daemon

import (
	"log"
	"os/exec"
	"time"

	"github.com/shirou/gopsutil/process"
)

func StartPalwordServer(path string) (int, error) {

	cmd := exec.Command(path)
	err := cmd.Start()
	if err != nil {
		return 0, err
	}
	return cmd.Process.Pid, err
}

func ListenProcess(palWorldPid int, limit float32) error {
	log.Printf("获取到帕鲁服务端pid为: %d", palWorldPid)
	pros, err := process.Processes()
	if err != nil {
		return err
	}
	for _, p := range pros {
		if p.Pid == int32(palWorldPid) {
			for {
				palWorldMem, err := p.MemoryPercent()
				if err != nil {
					return err
				}
				if palWorldMem >= limit {
					log.Printf("内存使用率为:%0.2f,准备重启帕鲁服务端!!!!", palWorldMem)
					err := p.Kill()
					if err != nil {
						return err
					}
					return nil
				} else {
					log.Printf("当前帕鲁服务端内存使用率为：%0.2f", palWorldMem)
					time.Sleep(time.Second * 60)
				}
			}
		}
	}
	return nil
}
