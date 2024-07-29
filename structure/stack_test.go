package structure

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	a := NewStack[string]()
	a.Push("abc")
	a.Push("abt")
	c, e := a.Peek()
	fmt.Println(c, e)
	c, e = a.Pop()
	fmt.Println(c, e)

}
