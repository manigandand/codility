// To execute Go code, please declare a func main() in a package "main"

package main

import (
	"fmt"
	"sync"
)

// StackWithTwo interface
type StackWithTwo interface {
	pushFirst(i int)
	popFirst() (last int)
	pushSecond(i int)
	popSecond() (last int)
}

func main() {
	s := NewStackWithTwo(10000000)
	s.pushFirst(1)
	s.pushFirst(2)

	s.pushSecond(3)

	fmt.Println(s.popFirst())
	fmt.Println(s.popSecond())
	fmt.Println(s.popFirst())
	fmt.Println(s.popFirst())
}

// Stack struct holds the stack properties
type Stack struct {
	stacks map[int][]int
	wg     sync.RWMutex
}

// NewStackWithTwo returns StackWithTwo interface
func NewStackWithTwo(capacity int) StackWithTwo {
	var s Stack
	s.stacks = make(map[int][]int)
	s.stacks[0] = make([]int, 0, capacity)
	s.stacks[1] = make([]int, 0, capacity)

	// fmt.Println("Size of s.stacks[0]:", unsafe.Sizeof(s.stacks[0]))
	// fmt.Printf("len=%d cap=%d %v\n",
	// 	len(s.stacks[0]), cap(s.stacks[0]), s.stacks[0])
	// fmt.Printf("len=%d cap=%d %v\n",
	// 	len(s.stacks[1]), cap(s.stacks[1]), s.stacks[1])

	return &s
}

func (s *Stack) pushFirst(i int) {
	s.wg.Lock()
	s.stacks[0] = append(s.stacks[0], i)
	s.wg.Unlock()
	// fmt.Printf("len=%d cap=%d %v\n",
	// 	len(s.stacks[0]), cap(s.stacks[0]), s.stacks[0])
	// fmt.Println("Size of s.stacks[0]:", unsafe.Sizeof(s.stacks[0]))
}

func (s *Stack) popFirst() (last int) {
	s.wg.RLock()
	v := s.stacks[0]
	s.wg.RUnlock()
	if len(v) == 0 {
		return
	}
	last = v[len(v)-1]

	s.wg.Lock()
	s.stacks[0] = v[:len(v)-1]
	s.wg.Unlock()

	// fmt.Printf("len=%d cap=%d %v\n",
	// 	len(s.stacks[0]), cap(s.stacks[0]), s.stacks[0])

	return
}

func (s *Stack) pushSecond(i int) {
	s.wg.Lock()
	s.stacks[1] = append(s.stacks[1], i)
	s.wg.Unlock()
	// fmt.Printf("len=%d cap=%d %v\n",
	// 	len(s.stacks[1]), cap(s.stacks[1]), s.stacks[1])
}

func (s *Stack) popSecond() (last int) {
	s.wg.RLock()
	v := s.stacks[1]
	s.wg.RUnlock()
	if len(v) == 0 {
		return
	}
	last = v[len(v)-1]

	s.wg.Lock()
	s.stacks[1] = v[:len(v)-1]
	s.wg.Unlock()

	// fmt.Printf("len=%d cap=%d %v\n",
	// 	len(s.stacks[0]), cap(s.stacks[0]), s.stacks[0])

	return
}

/*
Your previous Plain Text content is preserved below:

class StackWithTwo(int CAPACITY) {

    int [] arr;

    void pushFirst(int arg) {

    }

    int popFirst() {

    }

    void pushSecond(int arg) {

    }

    int popSecond() {

    }
}

main() {

    StackWithTwo s = new StackWithTwo(10);

    s.pushFirst(1); []1
    s.pushFirst(2); []1,2
    s.pushSecond(3); []3

    s.popFirst(); //2 []1
    s.popSecond(); //3 []
    s.popFirst(); //1
}
*/
