package lib

/*
 * golang中没有直接可用的迭代器接口,所以要定义一个迭代器接口
 * 这样做的好处是集合和迭代器职责分离，做到职责单一
 * 对于复杂的集合结构，可以按照自己的需求实现迭代器，无需改动集合累
 * 而且方便更换跌打方案
 */

type Iterator interface {
	HasNext() bool //判断是否迭代结束
	Next()         //迭代指针向后移动
	Item() any     //获取当前元素
	Reset()        //重置指针
}
