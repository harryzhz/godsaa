package tests

import (
	"fmt"
	"testing"
)

// 函数
func Hello() {
	fmt.Println("Hello World!")
}

type Welcome struct {
	name string
}

// 方法
func (w Welcome) Hello() {
	fmt.Println("Hello World!")
}

// 指针方法
func (w *Welcome) SetName(name string) {
	w.name = name
}

// 值方法
func (w Welcome) Welcome() {
	fmt.Printf("Welcome %s\n", w.name)
}

func NewWelcome() Welcome {
	return Welcome{}
}

func NewWelcomePtr() *Welcome {
	return &Welcome{}
}

func TestWelcome(t *testing.T) {
	// 正确: 值类型调用值类型的方法
	NewWelcome().Hello()

	// 正确: 指针类型的变量可以调用指针方法和值方法
	NewWelcomePtr().SetName("Harry")
	NewWelcomePtr().Welcome()

	// 正常: w 为值类型的变量，但是因为 w 是可寻址的，所以编译器会自动转换
	w := NewWelcome()
	w.SetName("Harry")
	w.Welcome()

	// 报错: cannot call pointer method SetName on Welcome。
	// NewWelcome().SetName("Harry")
}
