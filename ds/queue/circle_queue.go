package queue

import "fmt"

/*
 * 循环队列的实现
 * 这里空队列的时候,头尾指针都是-1
 */

type CircleQueue struct {
	items []any
	size  int
	head  int
	tail  int
}

//NewCircleQueue 如果是空队列,头尾指针都是-1
func NewCircleQueue(size int) *CircleQueue {
	return &CircleQueue{
		items: make([]any, size),
		size:  size,
		head:  -1,
		tail:  -1,
	}
}

func (c *CircleQueue) IsEmpty() bool {
	return c.head == -1
}

func (c *CircleQueue) IsFull() bool {
	return (c.tail+1)%c.size == c.head
}

func (c *CircleQueue) Clear() {
	c.items = make([]any, c.size)
	c.head = -1
	c.tail = -1
}

//EnQueue 入队列从尾部入
func (c *CircleQueue) EnQueue(item any) bool {
	if item == nil {
		return false
	}

	if c.IsFull() {
		return false
	}

	if c.IsEmpty() {
		c.head = 0
	}

	c.tail = (c.tail + 1) % c.size
	c.items[c.tail] = item
	return true
}

//DeQueue 出队列从头部出
func (c *CircleQueue) DeQueue() any {
	if c.IsEmpty() {
		return nil
	}
	item := c.items[c.head]
	//如果头尾指针是一个,说明,循环队列只有一个元素,则把队列置空
	if c.head == c.tail {
		c.head = -1
		c.tail = -1
		c.items = make([]any, c.size)
		return item
	}

	c.head = (c.head + 1) % c.size
	return item
}

func (c *CircleQueue) GetTail() any {
	if c.IsEmpty() {
		return nil
	}
	return c.items[c.tail]
}

func (c *CircleQueue) Debug() string {
	return fmt.Sprintf("%v->%v", c.head, c.tail)
}

func (c *CircleQueue) Iter() []any {
	items := make([]any, 0, c.size)
	if c.IsEmpty() {
		return items
	}
	start := c.head
	for {
		items = append(items, c.items[start])
		if start == c.tail {
			break
		}
		start = (start + 1) % c.size
	}
	return items
}
