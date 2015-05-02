// DRouterD 守护进程，当有配置文件更新的适合自动刷新DRouter进程

package main

import (
	"DRouter/conf" //自动检查配置文件的
	"fmt"
	"os"
	"os/exec"
	"time"
)

// Service 本程序就是来守护这个进程的
var Service *os.Process

func testTimer1() {
	if IsChanged() {
		go func() {
			KillService()
			StartService()
		}()
	}

}

func timer1() {
	timer1 := time.NewTicker(time.Duration((conf.Gconf["FlushPeriod"]).(int)) * time.Second)
	for {
		select {
		case <-timer1.C:
			testTimer1()
		}
	}
}

// StartService 开始service
func StartService() {
	cmd := exec.Command(conf.Gconf["DRouterPath"].(string), os.Args[1])
	go func() {
		err := cmd.Run()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(3)
		}
	}()
	Service = cmd.Process
}

// KillService 结束service
func KillService() {
	Service.Kill()
}

// IsChanged 如果配置文件修改了，就刷新服务
func IsChanged() bool {
	return true
}

func main() {
	fmt.Println("DRouter守护进程开始")
	go func() { StartService() }()
	fmt.Println("进程已经开始了:", " 5秒钟之后结束http service")
	time.Sleep(10 * time.Second)
	KillService()
	os.Exit(0)
	time.Sleep(time.Second * time.Duration((conf.Gconf["FlushPeriod"]).(int)))
	timer1()
}
