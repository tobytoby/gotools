package iter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func createClassRoom(size int) *ClassRoom {
	name := []string{"toby", "lili", "tom", "liming", "job"}
	cr := NewClassRoom(size)
	for i := 0; i < size; i++ {
		nameRand := rand.Intn(5)
		st := &Student{
			Name: fmt.Sprintf("%s_%d", name[nameRand], 1+rand.Intn(size)),
			Age:  18 + rand.Intn(5),
		}
		st.id, _ = st.GetId()
		_, err := cr.Add(st)
		if err != nil {
			panic(fmt.Sprintf("增加新生报错:%s, err:%s", st, err.Error()))
		}
	}
	return cr
}

func TestClassRoomIterator(t *testing.T) {
	rand.Seed(time.Now().Unix())
	cr := createClassRoom(20)
	if cr.Count() < 20 {
		t.Logf(fmt.Sprintf("注册学生人数少于实际报道人数"))
	}
	//测试迭代逻辑
	iterator := cr.Iterator()
	i := 0
	for iterator.HasNext() {
		expect := cr.getStudentByIndex(i)
		st := iterator.Item()
		assert.Equal(t, expect, st)
		iterator.Next()
		i++
	}
	t.Logf("迭代了这么多学生:%d", i)
	iterator.Reset()
	i = 0
	for iterator.HasNext() {
		expect := cr.getStudentByIndex(i)
		st := iterator.Item()
		assert.Equal(t, expect, st)
		iterator.Next()
		i++
	}
	t.Logf("迭代了这么多学生:%d", i)
}
