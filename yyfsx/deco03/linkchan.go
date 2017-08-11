package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() //表示让cpu将控制权给其他人
		fmt.Println(s)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go say("world")
	say("hello")
}
