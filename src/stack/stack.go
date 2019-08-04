package stack

import (
	"errors"
	"strconv"
)

const LIMIT  = 10

type obj interface {}

type Stack struct {
	idx int
	data []obj
}

/*
newStack 初始化stack，大小为size，blablabla。。。。
 */
func NewStack(size int) (stack *Stack, err error){
	if (stack != nil) {
		err = errors.New("stack已经存在，不需要再初始化！")
	}
	if size < 0 {
		stack = nil
		err = errors.New("size不能小于0")
	}
	if size > LIMIT {
		stack = nil
		errors.New("size不能大于" + strconv.Itoa(LIMIT))
	}
	stack = &Stack{-1, make([]obj, size)}
	err = nil
	return
}

func (s *Stack) Push(element obj) (idx int, err error) {
	if s == nil {
		errors.New("需要先调用newStack方法构建stack！")
	}
	if s.idx + 1  == len(s.data) {
		errors.New("stack已满！")
	}
	s.idx ++
	s.data[s.idx] = element
	err = nil
	return
}

func (s * Stack) Pop() (data obj, err error) {
	if s == nil {
		errors.New("需要先调用newStack方法构建stack！")
	}
	if s.idx == -1 {
		errors.New("stack为空！")
	}
	data = s.data[s.idx]
	err = nil
	return
}


