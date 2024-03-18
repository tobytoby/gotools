package set

type OSetIterator struct {
	s         *OSet
	index     int
	sortIndex int
}

func (oi *OSetIterator) HasNext() bool {
	if oi.s.Size() == 0 {
		return false
	}
	k := oi.s.index[oi.index]
	if oi.index > (len(oi.s.index)-1) || (oi.index == (len(oi.s.index)-1) && oi.sortIndex > (len(oi.s.items[k])-1)) {
		return false
	}
	return true
}

func (oi *OSetIterator) Next() {
	if oi.HasNext() {
		//如果还没访问到最后一个链
		if oi.index < (len(oi.s.index) - 1) {
			//没访问到最后一个元素
			if oi.sortIndex < (len(oi.s.items) - 1) {
				oi.sortIndex++
			} else {
				oi.index++
				oi.sortIndex = 0
			}
		} else if oi.index == (len(oi.s.index) - 1) {
			oi.sortIndex++
		}
	}
}

func (oi *OSetIterator) Item() any {
	k := oi.s.index[oi.index]
	for item, idx := range oi.s.items[k] {
		if idx == oi.sortIndex {
			return item
		}
	}
	return nil
}

func (oi *OSetIterator) Reset() {
	oi.index = 0
}
