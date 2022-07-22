package main

import (
	"container/heap"
	"fmt"
	"sync"
)

var first = true

type Order struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Temp      string  `json:"temp"`
	ShelfLife float64 `json:"shelfLife"`
	DecayRate float64 `json:"decayRate"`
}

type Order2 struct {
	Order
	start int64
	value float64
	o     bool // 是否是over中放入chann的
}

var channel = make(chan Order, 11)
var Overflow []Order2
var wg sync.WaitGroup

func main() {
	heap.Init(cold)
	heap.Init(hot)
	heap.Init(frozen)
	heap.Init(overflow)
	var order []Order
	////
	wg.Add(2)
	ReadJson(&order)
	////fmt.Println(order)
	//
	go send(order)
	go receive()
	go invite()

	wg.Wait()
	fmt.Println("结束")

}
