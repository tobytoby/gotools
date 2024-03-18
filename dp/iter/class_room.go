package iter

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tobytoby/gotools/lib"
	"math/rand"
	"time"
)

//Student 学生
type Student struct {
	id   string `json:"id"` //学号
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//GetId 学生注册学号 这里的随机需要服务在初始化时候先设置随机种子 rand.Seed(time.Now().UnixNano())
func (s *Student) GetId() (string, error) {
	if s.Name == "" || s.Age <= 0 {
		return "", errors.New("不合法的学生")
	}
	t := time.Now().Format("2006010215")
	n1 := rand.Intn(90) + 10
	n2 := rand.Intn(90) + 10
	return fmt.Sprintf("%s%s%d%d%d", "G", t, s.Age, n1, n2), nil
}

func (s *Student) String() string {
	bt, _ := json.Marshal(s)
	return string(bt)
}

type ClassRoom struct {
	ids      []string                   //学生学号列表
	planNum  int                        //计划招生数量
	names    map[string]map[string]bool //学生名称列表: name -> id map,名称可能同名
	students map[string]*Student        //学生名单 id -> Student 便于查找学生
}

func NewClassRoom(size int) *ClassRoom {
	return &ClassRoom{
		planNum:  size,
		ids:      make([]string, 0, size),
		names:    make(map[string]map[string]bool, size),
		students: make(map[string]*Student, size),
	}
}

//BeyondPan 是否超出计划招生
func (c *ClassRoom) BeyondPan() bool {
	return c.Count() > c.planNum
}

func (c *ClassRoom) Count() int {
	return len(c.students)
}

//Has 是否有叫这个名字的学生
func (c *ClassRoom) Has(name string) []*Student {
	ids, ok := c.names[name]
	if !ok {
		return nil
	}
	students := make([]*Student, 0, 2)
	for id, _ := range ids {
		students = append(students, c.students[id])
	}
	return students
}

//Add 新学生前来报道
func (c *ClassRoom) Add(s *Student) (bool, error) {
	if s.id == "" {
		return false, errors.New("学生未注册学号")
	}

	_, ok := c.students[s.id]
	if ok {
		return true, nil
	}

	c.students[s.id] = s
	c.ids = append(c.ids, s.id)
	_, nameOk := c.names[s.Name]
	if !nameOk {
		c.names[s.Name] = make(map[string]bool, 2)
	}
	c.names[s.Name][s.id] = true

	return true, nil
}

//Remove 学生转学离开
func (c *ClassRoom) Remove(s *Student) (bool, error) {
	_, ok := c.students[s.id]
	if !ok {
		return false, errors.New("学校没有这个学生")
	}
	delete(c.students, s.id)
	ids, _ := c.names[s.Name]
	for id, _ := range ids {
		if id == s.id {
			delete(c.students, s.id)
		}
	}
	filterIds := make([]string, 0, len(c.students))
	for _, id := range c.ids {
		if id != s.id {
			filterIds = append(filterIds, id)
		}
	}
	c.ids = filterIds
	return true, nil
}

func (c *ClassRoom) Iterator() lib.Iterator {
	return &ClassRoomIterator{
		classRoom: c,
	}
}

//getStudentByIndex 本函数是为了配合单元测试实现,业务逻辑一般没有这个逻辑
func (c *ClassRoom) getStudentByIndex(idx int) *Student {
	if idx > (c.Count() - 1) {
		return nil
	}
	return c.students[c.ids[idx]]
}
