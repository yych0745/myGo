package main

import (
	"fmt"
	"time"
)

//  取餐
func invite() {
	for {
		var item *Item

		item, ok := <-chann
		if ok && time.Now().Unix() >= item.priority {
			fmt.Println(time.Now().Unix())
			fmt.Println("过期", item.value)
			continue
		} else if ok {
			fmt.Println("+++++++++未过期,取餐+++++++++++++\n", item.value.Order)
			fmt.Println("value:", getValue(item.value))
			printQue1(*cold, *hot, *frozen, *overflow)
		}
		if !ok && cold.Len() == 0 && hot.Len() == 0 && frozen.Len() == 0 && overflow.Len() == 0 {
			quit <- 1
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++接收结束+++++++++++++++++++++++++++++++++++++++++++++++++++++")
			wg.Done()
			return
		}
	}
}
