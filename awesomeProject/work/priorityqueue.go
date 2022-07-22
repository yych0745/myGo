package main

import (
	"container/heap"
)

//  Item  是优先队列中包含的元素。
type Item struct {
	value    Order2 //  元素的值，可以是任意字符串。
	priority int64  //  元素在队列中的优先级。
	//  元素的索引可以用于更新操作，它由  heap.Interface  定义的方法维护。
	index int //  元素在堆中的索引。
}

//  一个实现了  heap.Interface  接口的优先队列，队列中包含任意多个  Item  结构。
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	//  我们希望  Pop  返回的是最大值而不是最小值，
	//  因此这里使用大于号进行对比。
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	//item.index = -1 //  为了安全性考虑而做的设置
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Peek() interface{} {
	item := pq.Pop()
	pq.Push(item)
	return item
}

//  更新函数会修改队列中指定元素的优先级以及值。
func (pq *PriorityQueue) update(item *Item, value Order2, priority int64) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

//  这个示例首先会创建一个优先队列，并在队列中包含一些元素
//  接着将一个新元素添加到队列里面，并对其进行操作
//  最后按优先级有序地移除队列中的各个元素。
