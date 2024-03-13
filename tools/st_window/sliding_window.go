package lib

import (
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	"github.com/shopspring/decimal"

	"github.com/tobytoby/gotools/ds/queue"
)

/*
 * 这是一个滑动窗口的实现逻辑
 * 可以用滑动窗口来平滑的控制流量
 */

type CounterItem struct {
	Req uint64
	s   uint64
}

//记数模式
type cntMode uint8

const (
	CntModeBefore cntMode = iota + 1 //先记数,再判断是否允许请求,这种可能会导致请求量大的时候,qps较小,导致连续失败
	CntModeAfter                     //判断是否允许请求,之后再记数,这种会友好些,qps较小的时候,一段请求失败之后,下一秒的请求可能会成功
)

//SlidingWindow 实现平滑时间窗口的逻辑
//每个计数器是以秒为单位的
//以循环队列用于存储计时器
type SlidingWindow struct {
	Qps           int64              //允许的每秒的最大请求数
	Size          int                //累计多少个滑动时间窗口
	Counter       *queue.CircleQueue //计数器
	cm            cntMode
	LastCheckTime time.Time
	isDebug       bool
}

func NewSlidingWindow(qps int64, size int, cm cntMode) *SlidingWindow {
	return &SlidingWindow{
		Qps:           qps,
		Size:          size,
		Counter:       queue.NewCircleQueue(size),
		cm:            cm,
		LastCheckTime: time.Now(),
	}
}

func (sw *SlidingWindow) Debug(d bool) {
	sw.isDebug = d
}

func (sw *SlidingWindow) String() string {
	items := sw.Counter.Iter()
	reqs := make([]uint64, len(items))
	for i, v := range items {
		vs := v.(*CounterItem)
		reqs[i] = vs.Req
	}
	return strings.ReplaceAll(fmt.Sprintf("%v", reqs), " ", ",")
}

func (sw *SlidingWindow) totalRequest() (uint64, int) {
	items := sw.Counter.Iter()
	allCnt := uint64(0)
	for _, v := range items {
		vs := v.(*CounterItem)
		allCnt += vs.Req
	}
	return allCnt, len(items)
}

//calcWindow 累加时间窗口,这里需要当前请求时间,和最后一个窗口的时间做比较,识别是否要创建下一个时间窗口
func (sw *SlidingWindow) calcWindow() {
	if sw.Counter.IsEmpty() {
		sw.Counter.EnQueue(&CounterItem{
			Req: 0,
			s:   uint64(time.Now().Unix()),
		})
		return
	}

	ns := uint64(time.Now().Unix())
	tailItem := sw.Counter.GetTail()
	tail := tailItem.(*CounterItem)
	//比较当前时间和计数器窗口的尾部指针的时间
	if sw.isDebug {
		fmt.Printf("时间比较:%d->%d\n", ns, tail.s)
	}
	if ns > tail.s {
		//如果窗口已经满了,弹出第一个元素
		if sw.Counter.IsFull() {
			sw.Counter.DeQueue()
		}
		sw.Counter.EnQueue(&CounterItem{
			Req: 0,
			s:   uint64(time.Now().Unix()),
		})
		return
	}
}

//count 请求记数
func (sw *SlidingWindow) count() {
	tailItem := sw.Counter.GetTail()
	tail := tailItem.(*CounterItem)
	atomic.AddUint64(&(tail.Req), 1)
}

func (sw *SlidingWindow) AllowRequest() bool {
	//先创建窗口
	sw.calcWindow()
	//如果是先记数模式
	if sw.cm == CntModeBefore {
		sw.count()
	}

	//累计总数
	totalReq, cnt := sw.totalRequest()
	//计算当前qps的算法可以考虑继续优化,现在的算法是 请求总数 / 窗口数量 = 平均请求次数
	curQps := decimal.NewFromInt(int64(totalReq)).Div(decimal.NewFromInt(int64(cnt))).Ceil().IntPart()
	if sw.isDebug {
		println(fmt.Sprintf("当前计算细节: %d:%s / %d compare %d", totalReq, sw, cnt, sw.Qps))
	}

	allow := curQps <= sw.Qps
	//如果是后记数模式,则请求成功才记录数据
	if sw.cm == CntModeAfter && allow {
		sw.count()
	}

	return allow
}
