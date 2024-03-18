package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOSet_Append(t *testing.T) {
	s1 := NewOSet(5)
	s1.Append("toby")
	s1.Append("lili")
	s1.Append("lili")
	s1.Append("liming")
	s1.Append("dachui")
	t.Logf("s1 size:%d", s1.Size())
	iter := s1.Iterator()
	idx := 0
	for iter.HasNext() {
		assert.Equal(t, []string{"toby"}, iter.Item())
		iter.Next()
		idx++
	}
	iter.Reset()
}
