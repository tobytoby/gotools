package lib

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//打印debug日志
func doLog(fmtTpl string, prefix string, args ...any) {
	tpl := ""
	if prefix != "" {
		tpl = fmt.Sprintf("[%s]%s%s", time.Now().Format("2006-01-02 15:04:05.000"), prefix, fmtTpl)
	} else {
		tpl = fmt.Sprintf("[%s]%s", time.Now().String(), fmtTpl)
	}
	println(fmt.Sprintf(tpl, args...))
}

//模拟每秒的请求次数
func simReq(sw *SlidingWindow, index int, n int64) {
	rd := rand.Intn(int(n)-1) + 1
	prefix := fmt.Sprintf("轮次[%d]", index)
	doLog("请求开始>>>>", prefix)
	doLog("请求 %d 次", prefix, rd)
	for i := 0; i < rd; i++ {
		if sw.AllowRequest() {
			doLog("第 %d 次请求, 成功", prefix, i+1)
		} else {
			doLog("第 %d 次请求, 失败", prefix, i+1)
		}
	}
	doLog("请求结束<<<<\n", prefix)
}

func TestFlowControl(t *testing.T) {
	qps := int64(6)  //每秒的请求次数
	counterSize := 5 //连续累积多少个时间窗口
	sw := NewSlidingWindow(qps, counterSize, CntModeAfter)
	sw.Debug(true)
	rand.Seed(time.Now().Unix())

	for i := 0; i < 9; i++ {
		go simReq(sw, i, qps*2)
		time.Sleep(1 * time.Second)
	}
}
