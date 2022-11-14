// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"


type stack struct {
    elemFreq map[int]int // not useful for now atleast
    freqElems map[int][]int
    maxFreq int
    stk []int
}

func newStack(ip []int) *stack{
    s := &stack {
        elemFreq: make(map[int]int),
        freqElems: make(map[int][]int),
        maxFreq: 0,
        stk: ip,
    }

    for _, i := range ip {
        v, ok := s.elemFreq[i]
        if !ok {
            s.elemFreq[i] = 1
            s.addElemInFreq(1, i)
            continue
        }
        
        elemFreq := v+1
        s.elemFreq[i] = elemFreq
        s.addElemInFreq(elemFreq, i)

        if s.maxFreq < elemFreq {
            s.maxFreq = elemFreq
        }
    }

    return s
}

func (s *stack) addElemInFreq(freq, elem int) {
    elems, ok := s.freqElems[freq]
    if !ok {
        s.freqElems[freq] = []int{elem}
        return
    }

    elems = append(elems, elem)
    s.freqElems[freq] = elems
}


// push
func (s *stack) push(elem int) {
    v, ok := s.elemFreq[elem]
    if !ok {
        s.elemFreq[elem] = 1

    } else {
        s.elemFreq[elem] = v+1
    }

    if s.maxFreq < v+1 {
        s.maxFreq = v+1
    }

    s.stk = append(s.stk, elem)
}

// remove last insterted element from the stack
// 1. remove the last insterted element which has highest frequency || what is the highest frequency elem
// 2. in case of matching frequency, return last inserted element
// 3. else return the last insterted element
func (s *stack) pop() int {
    if len(s.stk) == 0 {
        return 0
    }
    if len(s.stk) == 1 {
        res := s.stk[0]
        s.stk = []int{} // empty
        return res
    }
    if s.maxFreq == 1 { // o(m)
        // pop the last elem
        res := s.stk[len(s.stk)]
        s.stk = s.stk[:len(s.stk)]
        return res
    }
    // get the max freq all elems in order, pop the last insterted from the slice 
    fmt.Println(s.maxFreq) // 2 -> [1,2,3] // // [2,3,5,4,1,8,9]     5,3,5,4,6,5,3,8,9   2 -> 5,3
    elems, ok := s.freqElems[s.maxFreq]
    if !ok {
        return 0
    }
    return s.popStk(elems)
    // TODO: update the maxFreq value
}

func (s *stack) popStk(elems []int) int {
    fmt.Println(elems)
    for i:=len(s.stk)-1; i > 0; i-- {
        for _, e := range elems {
            if s.stk[i] == e {
                // TODO: new slice
                // before and after the i
                return e
            }
        }
    }
    return 0
}

func main() {
    s := newStack([]int{5,3,5,4,6,5,3,8,9})

    fmt.Println(s.pop())
}
