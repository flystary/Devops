package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("this is a NewTask")
		time.Sleep(time.Second)
	}
}
