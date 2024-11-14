package main

import "fmt"

type stackA struct {
	stack []int
}

type stackB struct {
	stack []int
}

func (a *stackA) sa() {
	// swap first two elem
	stackLen := len(a.stack)
	if stackLen < 2 {
		return
	}
	a.stack[stackLen-1], a.stack[stackLen-2] = a.stack[stackLen-2], a.stack[stackLen-1]
}

func (b *stackB) sb() {
	// swap first two elem
	stackLen := len(b.stack)
	if stackLen < 2 {
		return
	}
	b.stack[stackLen-1], b.stack[stackLen-2] = b.stack[stackLen-2], b.stack[stackLen-1]
}

func ss(a *stackA, b *stackB) {
	a.sa()
	b.sb()
}

func pop(stack []int) (int, []int) {
	stackLen := len(stack)
	lastElem := stack[stackLen-1]
	stack = stack[:stackLen-1]
	return lastElem, stack
}

func (b *stackB) pa(a *stackA) {
	stackLength := len(b.stack)
	if stackLength < 1 {
		return
	}
}

func main() {
	stackA := []int{5, 8, 9, 3, 4, 8, 7}
	fmt.Println(pop(stackA))
}
