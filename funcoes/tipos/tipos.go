package tipos

import "fmt"

type Stack interface {
	push(v int)
	pop() int
	size() int
}

type StackArray struct {
	arr  []int
	topo int
}

func (p *StackArray) push(v int) {
	p.arr[p.topo] = v
	p.topo++
}

func (p *StackArray) pop() int {
	p.topo--
	return p.arr[p.topo]
}

func Main() {
	var stack = StackArray{arr: make([]int, 10), topo: 1}
	stack.push(1)
	stack.push(2)
	stack.push(3)

	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
}
