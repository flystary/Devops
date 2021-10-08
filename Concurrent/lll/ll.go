package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	t := now.Unix()
	t1 := 1633446000
	t2 := 1633456800
	defta := t2 - t1
	fmt.Printf("当前时间的时间戳%v\n %v秒\n %v小时", t, defta, defta/(60*60))
}
