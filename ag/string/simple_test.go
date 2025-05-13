package string

import (
	"testing"
	"time"
)

func TestSimpleSearch(t *testing.T) {
	s := "软件测试服务市场题库软件服务爱云校好分数大数据精准教学分析系统软件运维服务阅卷系统组卷系统"
	target := "好分大数据"
	start := time.Now()
	index := 0
	for i := 0; i < 100000; i++ {
		index = SimpleSearch(target, s)
	}

	if index < 0 {
		t.Logf("没有匹配到")
		return
	}

	tmp := []byte(s)[index:(index + len([]byte(target)))]
	t.Logf("在位置%d匹配到了,匹配到的字符串:%s,总耗时:%s", index, string(tmp), time.Now().Sub(start).String())
}

func TestDoublePointerSearch(t *testing.T) {
	s := "软件测试服务市场题库软件服务爱云校好分数大数据精准教学分析系统软件运维服务阅卷系统组卷系统"
	target := "好分大数据"
	start := time.Now()
	index := 0
	for i := 0; i < 100000; i++ {
		index = DoublePointerSearch(target, s)
	}

	if index < 0 {
		t.Logf("没有匹配到")
		return
	}

	tmp := []byte(s)[index:(index + len([]byte(target)))]
	t.Logf("在位置%d匹配到了,匹配到的字符串:%s,总耗时:%s", index, string(tmp), time.Now().Sub(start).String())
}
