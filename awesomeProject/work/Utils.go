package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var ran = 0

func ReadJson(order *[]Order) {
	jsonFile, err := os.Open("orders.json")
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}
	defer jsonFile.Close()

	var s string
	inputReader := bufio.NewReader(jsonFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			break
		}
		s = s + inputString
	}

	if err := json.Unmarshal([]byte(s), &order); err == nil {
	} else {
		fmt.Println("转换失败")
	}
}

func putChan() {

	if ran == 0 && delete(cold) {
		put(cold)
	} else if ran == 1 && delete(hot) {
		put(hot)
	} else if ran == 2 && delete(frozen) {
		put(frozen)
	} else if ran == 3 && delete(overflow) {
		putO(overflow)
	} else if delete(cold) {
		put(cold)
	} else if delete(hot) {
		put(hot)
	} else if delete(frozen) {
		put(frozen)
	} else if overflow.Len() > 0 {

		putO(overflow)
	} else {
		fmt.Println("cold Len:", cold.Len(), "hot Len:", hot.Len(), "frozen Len", frozen.Len(), "overflow Len:", overflow.Len())
		fmt.Println("全部餐架为空")
		close(chann)
	}
	ran = (ran + 1) % 3

}

func delete(que *PriorityQueue) bool {
	for que.Len() > 0 {
		item := que.Peek()
		itemi := item.(*Item)
		if itemi.priority <= time.Now().Unix() {
			fmt.Println("订单价值小于等于0，抛弃")
			fmt.Println(itemi.value)
			que.Pop()
		} else if itemi.priority >= time.Now().Unix() {
			return true
		}
	}
	return false
}

func put(que *PriorityQueue) {

	item := getItemFQue(que)
	chann <- item

}

func putO(que *PriorityQueue) {

	item := getItemFQue(que)
	item.value.o = true
	chann <- item

}

func choice(order2 Order2) {
	switch order2.Temp {
	case "cold":
		insertCold(order2)
	case "hot":
		insertHot(order2)
	case "frozen":
		insertFrozen(order2)
	}
}

func insertCold(order2 Order2) {
	second := getSeconds(&order2, 1.0)
	item := Item{order2, second, 0}

	if cold.Len() >= 10 {
		insertOverflow(order2)
	} else {
		cold.Push(&item)
	}
}

func insertHot(order2 Order2) {

	second := getSeconds(&order2, 1.0)
	item := Item{order2, second, 0}

	if hot.Len() >= 10 {
		insertOverflow(order2)
	} else {
		hot.Push(&item)
	}
}

func insertFrozen(order2 Order2) {

	second := getSeconds(&order2, 1.0)
	item := Item{order2, second, 0}

	if frozen.Len() >= 10 {
		insertOverflow(order2)
	} else {
		frozen.Push(&item)
	}
}
func insertOverflow(order2 Order2) {
	second := getSeconds(&order2, 2.0)

	//TODO  继续删除的逻辑
	if overflow.Len() >= 15 {
		item1 := overflow.Pop()
		orderR := item1.(*Item)
		orderI := orderR.value
		insertOrabandon(orderI)
	}

	item := Item{order2, second, 0}
	overflow.Push(&item)
}

func insertOrabandon(order2 Order2) {

	// 刷新订单开始时间和价值
	value := getOverValue(order2)
	order2.start = time.Now().Unix()
	order2.value = value
	second := getOverSecond(order2, value)

	if order2.Temp == "cold" {
		if cold.Len() >= 10 {
			fmt.Println("cold抛弃", order2.Order)
			printQue1(*cold, *hot, *frozen, *overflow)
		} else {
			// 重新计算其过期时间

			item := Item{order2, second + order2.start, cold.Len()}
			cold.Push(&item)
		}
	} else if order2.Temp == "hot" {
		if hot.Len() >= 10 {
			fmt.Println("hot抛弃", order2.Order)
			printQue1(*cold, *hot, *frozen, *overflow)
		} else {
			item := Item{order2, second + order2.start, hot.Len()}
			hot.Push(&item)
		}
	} else if order2.Temp == "frozen" {
		if frozen.Len() >= 10 {
			fmt.Println("frpozen抛弃", order2.Order)
			printQue1(*cold, *hot, *frozen, *overflow)
		} else {
			item := Item{order2, second + order2.start, frozen.Len()}
			frozen.Push(&item)
		}
	}
}

func getSeconds(order2 *Order2, shelfDecayModifier float64) int64 {
	res := int64(order2.ShelfLife/(1+order2.DecayRate*shelfDecayModifier)) + order2.start
	return res
}

func getItemFQue(que *PriorityQueue) *Item {
	item := que.Pop()
	itemi := item.(*Item)
	return itemi
}

func getOverValue(order2 Order2) float64 {

	//  计算价值
	orderAge1 := time.Now().Unix() - order2.start
	orderAge := float64(orderAge1)
	res := (order2.ShelfLife - orderAge - order2.DecayRate*orderAge*2) / order2.ShelfLife

	return res
}

func getOverSecond(order2 Order2, value float64) int64 {
	// 求出value为0时的时间
	res := int64(order2.ShelfLife * value / (1 + order2.DecayRate))
	return res
}

// 获得当前餐架的全部信息
func printQue1(que1 PriorityQueue, que2 PriorityQueue, que3 PriorityQueue, que4 PriorityQueue) {
	fmt.Println("cold餐架,订单数:", que1.Len())
	for que1.Len() > 0 {
		item := getItemFQue(&que1)
		fmt.Println(item, item.index)
	}

	fmt.Println("hot餐架,订单数", que2.Len())
	for que2.Len() > 0 {
		item := getItemFQue(&que2)
		fmt.Println(item, item.index)
	}

	fmt.Println("frozen餐架,订单数", que3.Len())
	for que3.Len() > 0 {
		item := getItemFQue(&que3)
		fmt.Println(item, item.index)
	}

	fmt.Println("overflow餐架,订单数", que4.Len())
	for que4.Len() > 0 {
		item := getItemFQue(&que4)
		fmt.Println(item, item.index)
	}
}

func getValue(order2 Order2) float64 {

	//  计算价值
	orderAge1 := time.Now().Unix() - order2.start
	orderAge := float64(orderAge1)
	var res float64
	if order2.o == false {
		res = order2.value - orderAge*(1+order2.DecayRate)/order2.ShelfLife
	} else {
		res = order2.value - orderAge*(1+order2.DecayRate*2)/order2.ShelfLife
	}

	return res
}

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func write2file(param string) {
	f, err := os.OpenFile("test", os.O_RDONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("open file error :", err)
		return
	}
	// 关闭文件
	defer f.Close()
	// 字节方式写入
	_, err = f.Write([]byte("write : " + param))
	if err != nil {
		log.Println(err)
		return
	}
	// 字符串写入
	_, err = f.WriteString("writeString : " + param)
	if err != nil {
		log.Println(err)
		return
	}
}
