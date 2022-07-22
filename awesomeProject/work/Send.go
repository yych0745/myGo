package main

import (
	"fmt"
	"time"
)

//  发送信息每两秒一次
func send(order []Order) {
	var l = len(order)
	println(l)
	for i := 0; i < l; i++ {
		channel <- order[i]
		channel <- order[i+1]
		i = i + 1
		time.Sleep(2 * time.Second)
		fmt.Println("多少个_______________________________________________________________________________________________________________________________", i)
	}
	defer close(channel)
	fmt.Println("_______________________________________发送结束_______________________________________________________")
	wg.Done()
	return
}
