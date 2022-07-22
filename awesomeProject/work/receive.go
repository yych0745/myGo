package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cold = &PriorityQueue{}
var hot = &PriorityQueue{}
var frozen = &PriorityQueue{}
var overflow = &PriorityQueue{}
var chann = make(chan *Item, 2)
var quit = make(chan int, 2)

var t int64
var r int64

//  接收订单，放入餐架，选择一个放入通道内（随机放入，间隔(2 - 6）秒）
func receive() {
	for {
		value, ok := <-channel
		if ok {
			fmt.Println("接收订单")
			ti := time.Now().Unix()
			v := Order2{value, ti, 1, false}
			//  放入餐架
			fmt.Println("订单为", v)
			printQue1(*cold, *hot, *frozen, *overflow)
			choice(v)
		}

		// 首次直接取餐，之后每休息2-6秒进行取餐
		if len(chann) < 1 && (first || t < time.Now().Unix()) {
			for {
				r = rand.Int63n(7)
				if r >= 2 {
					break
				}
			}

			first = false
			t = time.Now().Unix() + r
			delete(cold)
			delete(hot)
			delete(frozen)
			delete(overflow)
			putChan()
		}

	}

}
