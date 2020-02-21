// To execute Go code, please declare a func main() in a package "main"

package main

import (
	"fmt"
	"sync"
)

/*
In Golang, there are array and slice
if the array with given size [10]int
```
var a [10]int
fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)
output: len=10 cap=10 [0 0 0 0 0 0 0 0 0 0]
```

So i am using slice here. Slices can be created with the built-in make function.
Slice is dynamically-sized arrays.

make takes three arguments. make(type, len, capacity).
```
var a = make([]int, 10) // length and capacity 10
fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)
output: len=10 cap=10 [0 0 0 0 0 0 0 0 0 0]
```

```
var a = make([]int, 0, 10) // length is 0 and capacity is 10
fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)
output: len=0 cap=10 []
```

Since i am initializing slice of len(0), the memory of of the each stack
will remain 24bytes. This will grow only the stack grows.

*/
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

// StackWithTwo interface
type StackWithTwo interface {
	pushFirst(i int)
	popFirst() (last int)
	pushSecond(i int)
	popSecond() (last int)
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
