package other

import (
	"errors"
)

/*
最简单的数据结构就是stack了, 就是先进后出了, 用slice就可以实现了;
 */
type Stack []interface{}

func (stack *Stack) Len() int {
	return len(*stack)
}

// 这个就是去看一下栈顶, 有点像Peek
func (stack *Stack) Top() (interface{}, error) {
	if length := (*stack).Len(); length == 0 {
		return nil, errors.New("The stack is empty!")
	} else {
		return (*stack)[length-1], nil
	}
}

// 入栈，无非就是append一下就好了;
func (stack *Stack) Push(v interface{}) {
	*stack = append(*stack, v)
}

// 出栈，也就是pop一下就好了;
func (stack *Stack) Pop() (interface{}, error) {
	if length := (*stack).Len(); length == 0 {
		return nil, errors.New("This stack is empty!")
	} else {
		value := (*stack)[length-1]
		*stack = (*stack)[:(length-1)]
		return value, nil
	}
}
