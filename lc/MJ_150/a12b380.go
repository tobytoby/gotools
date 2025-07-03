package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"math/rand"
	"time"
)

/*
 * O(1)时间插入、删除和获取随机元素，实现RandomizedSet 类
 * https://leetcode.cn/problems/insert-delete-getrandom-o1/?envType=study-plan-v2&envId=top-interview-150
 */

/*
 * 变长数组可以实现O(1)随机返回元素，但是不能O(1)判断数据是否存在，进而删除或者插入，map可以O(1)实现判断数据是否存在，而删除或者插入
 * 但是无法O(1)随机返回元素，所以需要数组和map结合,数组保存元素,map保存元素的下标
 */
type RandomizedSet struct {
	data map[string]int
	list []any
}

func NewRandomizedSet() *RandomizedSet {
	return &RandomizedSet{
		data: make(map[string]int),
		list: make([]any, 0),
	}
}

func (this *RandomizedSet) hash(item any) string {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(item)
	if err != nil {
		//退而求其次，使用fmt
		return fmt.Sprintf("%v", item)
	}
	return buf.String()
}

func (this *RandomizedSet) Insert(item any) bool {
	h := this.hash(item)
	_, ok := this.data[h]
	if ok {
		return false
	}
	this.data[h] = len(this.list)
	this.list = append(this.list, item)
	return false
}

func (this *RandomizedSet) Remove(item any) bool {
	h := this.hash(item)
	index, ok := this.data[h]
	if !ok {
		return false
	}
	//这里为了避免数据大面积迁移,可以采取要删除的元素和最后一个元素调换位置,之后删除最后一个元素的方式
	last := len(this.list) - 1
	//list列表要删除的位置替换为最后一个元素的值
	this.list[index] = this.list[last]
	//data列表原来最后一个元素值的位置修改
	nh := this.hash(this.list[last])
	this.data[nh] = index
	//list删除最后一个元素
	this.list = this.list[:last]
	//data删除要删除的元素
	delete(this.data, h)

	return true
}

func (this *RandomizedSet) GetRandom() any {
	return this.list[rand.Intn(len(this.list))]
}

func (this *RandomizedSet) Iter() {
	for _, v := range this.data {
		fmt.Printf("%d:%+v\n", v, this.list[v])
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	s := NewRandomizedSet()
	s.Insert(map[string]any{
		"name": "toby01",
		"age":  1,
	})
	s.Insert(map[string]any{
		"name": "toby02",
		"age":  2,
	})
	s.Insert(map[string]any{
		"name": "toby03",
		"age":  3,
	})
	s.Insert(map[string]any{
		"name": "toby04",
		"age":  4,
	})
	s.Insert(map[string]any{
		"name": "toby05",
		"age":  5,
	})
	s.Insert(map[string]any{
		"name": "toby06",
		"age":  6,
	})
	s.Insert(map[string]any{
		"name": "toby07",
		"age":  7,
	})
	s.Insert(map[string]any{
		"name": "toby08",
		"age":  8,
	})
	s.Insert(map[string]any{
		"name": "toby09",
		"age":  9,
	})
	s.Insert(map[string]any{
		"name": "toby10",
		"age":  10,
	})
	fmt.Printf("%+v\n", s.GetRandom())
	fmt.Printf("%+v\n", s.GetRandom())
	fmt.Printf("%+v\n", s.GetRandom())
	fmt.Printf("%+v\n", s.GetRandom())
}
