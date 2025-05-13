package string

import (
	"strings"
	"testing"
	"time"
)

func TestKMPSearch(t *testing.T) {
	s := "服务市场题库软件服务爱云校好分数大数据精准教学分析系统软件运维服务阅卷系统组卷系统"
	target := "组卷系统"
	index := 0
	start := time.Now()
	for i := 0; i < 1; i++ {
		index = KMPSearch(target, s)
	}

	if index <= 0 {
		t.Logf("没有匹配到")
		return
	}

	tmp := []byte(s)[index:(index + len([]byte(target)))]
	t.Logf("在位置%d匹配到了,匹配到的字符串:%s,总耗时:%s", index, string(tmp), time.Now().Sub(start).String())

}

func TestStringContainSearch(t *testing.T) {
	s := "服务市场题库软件服务爱云校好分数大数据精准教学分析系统软件运维服务阅卷系统组卷系统"
	target := "组卷系统"
	index := 0
	start := time.Now()
	for i := 0; i < 1; i++ {
		index = strings.Index(s, target)
	}

	if index <= 0 {
		t.Logf("没有匹配到")
		return
	}

	tmp := []byte(s)[index:(index + len([]byte(target)))]
	t.Logf("在位置%d匹配到了,匹配到的字符串:%s,总耗时:%s", index, string(tmp), time.Now().Sub(start).String())

}
