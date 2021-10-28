package main

import (
	"fmt"
	"time"
)

func main() {
// 	const layout1 = "Now, Monday Jan 02 15:04:05 JST 2006"
// 	const layout2 = "2006-01-02 15:04:05 JST"
// 	const layout3 = "2006-01-02 15:04:05"
// 	nowtime := time.Now()
// 	fmt.Println(nowtime)
// 	fmt.Println(nowtime.Format(layout1))
// 	fmt.Println(nowtime.Format(layout2))
// 	fmt.Println(nowtime.Format(layout3))
// 	fmt.Println(nowtime.Format(time.RFC3339))
	mylogger("aaaaa my log bbbbb")
	mylogger("aaaaa my log cccccc")
}

func mylogger (msg string){
	var logent string
	nowtime := time.Now()
    logent  = nowtime.Format(time.RFC3339) + " " + msg
	fmt.Println (logent)
}
