package set

/*
 * 简单实现set数据结构,保存原始值
 * 优点:实现简单,没有hash冲突,且节省内存
 * 缺点: 由于直接使用原始数据作为键,可能会导致比较和查找操作的效率低,且不支持自定义的等价关系
 */

type SSet struct {
	items map[any]struct{}
}

func NewSSet(size int) *SSet {
	return &SSet{
		items: make(map[any]struct{}, size),
	}
}

func (s *SSet) Copy() *SSet {
	if s.Size() == 0 {
		return NewSSet(0)
	}
	ns := NewSSet(s.Size())
	for v, _ := range s.items {
		ns.Append(v)
	}
	return ns
}

func (s *SSet) Append(item any) {
	s.items[item] = struct{}{}
}

func (s *SSet) Remove(item any) bool {
	if !s.Contains(item) {
		return false
	}
	delete(s.items, item)
	return true
}

func (s *SSet) Contains(item any) bool {
	_, ok := s.items[item]
	return ok
}

func (s *SSet) Size() int {
	return len(s.items)
}

//Inter 求交集
func (s *SSet) Inter(sp *SSet) *SSet {
	if s.Size() == 0 || sp.Size() == 0 {
		return NewSSet(0)
	}

	res := NewSSet(s.Size() + sp.Size())

	if sp.Size() == 0 {
		return res
	}
	for v, _ := range sp.items {
		if !s.Contains(v) {
			res.Append(v)
		}
	}
	return res
}

//Union 求合集
func (s *SSet) Union(sp *SSet) *SSet {
	if s.Size() == 0 {
		return sp
	}

	if sp.Size() == 0 {
		return s
	}
	res := s.Copy()
	for v, _ := range sp.items {
		if !res.Contains(v) {
			res.Append(v)
		}
	}
	return res
}

//Diff  求差集
func (s *SSet) Diff(sp *SSet) *SSet {
	if s.Size() == 0 {
		return NewSSet(0)
	}

	if sp.Size() == 0 {
		return s
	}
	res := NewSSet(s.Size())
	for v, _ := range s.items {
		if !sp.Contains(v) {
			res.Append(v)
		}
	}
	return res
}
