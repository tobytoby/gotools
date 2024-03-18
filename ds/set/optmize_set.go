package set

import (
	"fmt"
	"hash/fnv"
	"reflect"
	"time"

	"encoding/json"

	"github.com/tobytoby/gotools/lib"
	"github.com/tobytoby/gotools/utils"
)

/*
 * 优化版本实现set数据结构,数据保存会经过散列函数处理,支持散列冲突,支持顺序迭代
 * 解决hash冲突的方式是数据具体元素使用二维map存储,map存储的好处主要是为了解决集合操作对比的性能问题
 * 这里为了支持顺序迭代消耗了一定的存储空间,且有性能损耗,通常集合是不要求顺序迭代的,所以如果考虑这个方面,可以不支持顺序迭代,也可以考虑将
 * 讲数据结构修改为支持动态指定是有序或者乱序
 * 优点: 查找性能较好,等价关系的比较比较简单
 * 缺点: 比较浪费内存,且散列函数的性能要较高,且冲突要低
 */

type osItem struct {
	item any
	sort int
}

//OSet 此处的items参数可以继续优化,修改为链表存储
type OSet struct {
	index   []uint32               //这个主要是为了实现顺序迭代
	indexMp map[uint32]int         //为了记录散列值在index中的位置,便于删除的时候,快速查找
	size    int                    //集合里总共数据的大小
	items   map[uint32]map[any]int //二维的map 键是原始数据 值是数据存入的顺序,这里使用map,主要是考虑查找的性能
}

func NewOSet(size int) *OSet {
	return &OSet{
		index: make([]uint32, 0, size),
		items: make(map[uint32]map[any]int, size),
	}
}

func (s *OSet) hash(data any) uint32 {
	typ := reflect.TypeOf(data)
	value := reflect.ValueOf(data)
	vStr := ""
	switch typ.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		vStr = fmt.Sprintf("%d", value.Uint())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		vStr = fmt.Sprintf("%d", value.Int())
	case reflect.Float32, reflect.Float64:
		vStr = fmt.Sprintf("%0.8f", value.Float())
	case reflect.Ptr:
		//如果是指针,则获取指针地址,且用一个特殊字符作为前缀拼接
		pAddr := value.Pointer()
		vStr = fmt.Sprintf("##_P_##%v", pAddr)
	default:
		tt := value.Type().String()
		if tt == "time.Time" {
			vTime := value.Interface().(time.Time)
			vStr = vTime.Format("2006-01-02 15:04:05.000")
		} else if tt == "string" {
			vStr = value.String()
		} else {
			bts, _ := json.Marshal(value.Interface())
			vStr = string(bts)
		}
	}
	h := fnv.New32a()
	h.Write([]byte(vStr))
	return h.Sum32()
}

func (s *OSet) isTail(idx int) bool {
	return idx == (len(s.index) - 1)
}

func (s *OSet) isHead(idx int) bool {
	return idx == 0
}

func (s *OSet) removeIndex(idx int) {
	isHead := s.isHead(idx)
	if isHead {
		s.index = s.index[1:]
		return
	}

	isTail := s.isTail(idx)
	if isTail {
		s.index = s.index[0 : len(s.index)-1]
		return
	}

	tmp := s.index[0:idx]
	tmp = append(tmp, s.index[idx+1:]...)
	s.index = tmp
}

func (s *OSet) Copy() *OSet {
	if s.Size() == 0 {
		return NewOSet(0)
	}
	ns := NewOSet(s.size)
	for i, v := range s.index {
		ns.index[i] = v
	}
	for k, v := range s.indexMp {
		ns.indexMp[k] = v
	}
	for k, items := range s.items {
		tmpItems := make(map[any]struct{}, 5)
		for kk, _ := range items {
			tmpItems[kk] = struct{}{}
		}
		ns.items[k] = items
	}
	return ns
}

func (s *OSet) Append(item any) {
	k := s.hash(item)
	vs, ok := s.items[k]
	if !ok {
		vs = make(map[any]int, 2)
		s.index = append(s.index, k)
		s.indexMp[k] = len(s.index) - 1
	}

	vs[item] = len(vs)
	s.items[k] = vs
	s.size++
}

func (s *OSet) doDelete(k uint32) {
	delete(s.items, k)
	s.removeIndex(s.indexMp[k])
	delete(s.indexMp, k)
}

//Remove 可能会删除多个元素
func (s *OSet) Remove(item any) {
	k := s.hash(item)
	vs, ok := s.items[k]
	if !ok {
		return
	}

	deleteObj := 0
	if len(vs) > 1 {
		filterVs := make(map[any]int, 2)
		for v, weight := range vs {
			if item != v {
				filterVs[v] = weight
			} else {
				deleteObj++
			}
		}
		if len(filterVs) > 0 {
			//对过滤数据做排序值重置
			filterVs = utils.ResetWeight(filterVs)
			s.items[k] = filterVs
		} else {
			s.doDelete(k)
		}
	} else {
		deleteObj++
		s.doDelete(k)
	}
	s.size = s.size - deleteObj
}

func (s *OSet) Contains(item any) bool {
	k := s.hash(item)
	vs, ok := s.items[k]
	if !ok {
		return false
	}
	for _, v := range vs {
		if v == item {
			return true
		}
	}
	return false
}

func (s *OSet) Size() int {
	return s.size
}

//Inter 求交集
func (s *OSet) Inter(sp *OSet) *OSet {
	if sp.Size() == 0 || sp.size == 0 {
		return NewOSet(0)
	}
	res := NewOSet(s.Size() + sp.size)
	for k, items := range s.items {
		targetItems, ok := sp.items[k]
		if !ok {
			continue
		}
		interItems := make(map[any]int, 5)
		for sv, weight := range items {
			if _, ok := targetItems[sv]; ok {
				interItems[sv] = weight
			}
		}
		//这里通过手动添加的方式,而不是在175行通过append方式添加,主要是考虑性能
		if len(interItems) > 0 {
			res.items[k] = utils.ResetWeight(interItems)
			res.index = append(res.index, k)
			res.indexMp[k] = len(res.index) - 1
			res.size++
		}
	}
	return res
}

//Union 求合集
func (s *OSet) Union(sp *OSet) *OSet {
	if s.Size() == 0 {
		return sp
	}
	if sp.Size() == 0 {
		return s
	}
	res := s.Copy()
	for k, items := range sp.items {
		originItems, ok := res.items[k]
		if !ok {
			res.items[k] = items
			res.index = append(res.index, k)
			res.indexMp[k] = len(res.index) - 1
			res.size += len(items)
			continue
		}
		addCnt := 0
		for item, _ := range items {
			if _, ok := originItems[item]; !ok {
				originItems[item] = len(originItems)
				addCnt++
			}
		}
		res.items[k] = originItems
		res.size += addCnt
	}
	return res
}

//Diff  求差集
func (s *OSet) Diff(sp *OSet) *OSet {
	if s.Size() == 0 {
		return NewOSet(0)
	}
	if sp.Size() == 0 {
		return s
	}
	res := NewOSet(s.Size())
	for k, items := range s.items {
		targetItems, ok := sp.items[k]
		if !ok {
			res.items[k] = items
			res.index = append(res.index, k)
			res.indexMp[k] = len(res.index) - 1
			res.size += len(items)
			continue
		}
		addCnt := 0
		filterItems := make(map[any]int, 5)
		for item, weight := range items {
			if _, ok := targetItems[item]; !ok {
				filterItems[item] = weight
				addCnt++
			}
		}
		res.items[k] = utils.ResetWeight(filterItems)
		res.index = append(res.index, k)
		res.indexMp[k] = len(res.index) - 1
		res.size += len(items)
	}
	return res
}

func (s *OSet) Iterator() lib.Iterator {
	return &OSetIterator{
		s: s,
	}
}

//getItemsByIndex 本函数是为了配合单元测试实现,业务逻辑一般没有这个逻辑
func (s *OSet) getItemsByIndex(idx int) []any {
	if idx > (len(s.index) - 1) {
		return nil
	}

	k := s.index[idx]
	items := make([]any, 0, len(s.items[k]))
	for v := range s.items[k] {
		items = append(items, v)
	}
	return items
}
