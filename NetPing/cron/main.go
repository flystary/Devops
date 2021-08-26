package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"os/exec"
)

func CreateCron() {
	c := cron.New()
	spec := "*/5 * * * * ?"

	if err := c.AddFunc(spec, func() {
		err := NetWorkStatus()
		log.Println(err)
	}); err != nil {
		log.Println(err)
	}
	c.Start()
}

func NetWorkStatus() bool {
	cmd := exec.Command("ping", "www.baidu.com", "-n", "5")
	//fmt.Println("PING start:",time.Now().Format("2006/01/02 15:04:05"))
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

func main() {
	CreateCron()
	for {
	}
}
