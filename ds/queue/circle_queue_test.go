package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCircleQueue(t *testing.T) {
	q := NewCircleQueue(4)
	t.Logf("指针:%s", q.Debug())
	t.Logf("items:%v", q.Iter())
	assert.Equal(t, true, q.EnQueue(1))
	assert.Equal(t, true, q.EnQueue(2))
	assert.Equal(t, true, q.EnQueue(3))
	t.Logf("items:%v", q.Iter())
	assert.Equal(t, true, q.EnQueue(4))
	assert.Equal(t, false, q.EnQueue(5))
	t.Logf("指针:%s", q.Debug())
	t.Logf("items:%v", q.Iter())
	assert.Equal(t, 1, q.DeQueue())
	assert.Equal(t, 2, q.DeQueue())
	t.Logf("items:%v", q.Iter())
	assert.Equal(t, 3, q.DeQueue())
	assert.Equal(t, true, q.EnQueue(6))
	t.Logf("指针:%s", q.Debug())
	t.Logf("items:%v", q.Iter())
}
