package iter

/*
 * 班级学生的迭代器,这个迭代器可以随着业务的变化修改迭代逻辑
 * 比如初期按照学号迭代,后面需求变化需要按照名称迭代,或者动态的按照学号或者id迭代
 * 这样就可以在迭代器内部修改迭代逻辑,而不需要更改业务逻辑
 */

type ClassRoomIterator struct {
	classRoom *ClassRoom
	index     int
}

func (ci *ClassRoomIterator) HasNext() bool {
	cnt := ci.classRoom.Count()
	if cnt == 0 {
		return false
	}
	if ci.index > (cnt - 1) {
		return false
	}
	return true
}

func (ci *ClassRoomIterator) Next() {
	if ci.HasNext() {
		ci.index++
	}
}

func (ci *ClassRoomIterator) Item() any {
	id := ci.classRoom.ids[ci.index]
	return ci.classRoom.students[id]
}

func (ci *ClassRoomIterator) Reset() {
	ci.index = 0
}
