package tests

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tour/tree"
)

func TestCh(t *testing.T) {
	ch := make(chan int, 10)

	// 启动所有发送者协程
	wg := &sync.WaitGroup{}
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			ch <- v
			fmt.Println("send:", v)
		}(i)
	}

	// 启动接受者协程
	quit := make(chan struct{})
	go func() {
		for v := range ch {
			fmt.Println("receive:", v)
		}
		quit <- struct{}{}
	}()

	// 等待所有发送者协程结束，关闭通道
	wg.Wait()
	fmt.Println("send over...")
	close(ch)

	// 等待接收者协程结束
	<-quit
	fmt.Println("receive over!!!")
}

func TestTimeout(t *testing.T) {
	ch := make(chan int, 1)

	// 执行一个耗时的任务
	go func() {
		time.Sleep(2 * time.Second)
		ch <- 1
	}()

	select {
	case res := <-ch:
		// 任务执行完毕
		fmt.Println("result: ", res)
	case <-time.After(3 * time.Second):
		// 任务执行超时
		fmt.Println("timeout")
	}
}

func TestConcurrencyNumberLimit(t *testing.T) {
	const TotalNum = 10
	const ParallelNum = 2

	wg := &sync.WaitGroup{}
	// 限制并发数，缓冲区大小即为最大并发数
	ch := make(chan struct{}, ParallelNum)

	for i := 0; i < TotalNum; i++ {
		wg.Add(1)
		// 通道满时写入操作阻塞在这里，则不会继续起新的协程
		ch <- struct{}{}
		go func(idx int) {
			defer func() {
				wg.Done()
				<-ch
			}()

			fmt.Printf("[%s] process: %d/%d\n", time.Now().Format(time.DateTime), idx, TotalNum)
			time.Sleep(1 * time.Second)
		}(i + 1)
	}

	wg.Wait()
}

func producer(ch chan<- int) {
	defer close(ch)

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Printf("[%s] Produced: %d\n", time.Now().Format(time.DateTime), i)
		ch <- i
	}
}

func consumer(ch <-chan int) {
	for {
		if v, ok := <-ch; ok {
			time.Sleep(time.Second)
			fmt.Printf("[%s] Consumed: %d\n", time.Now().Format(time.DateTime), v)
		} else {
			break
		}
	}
}

func TestProdConsume(t *testing.T) {
	ch := make(chan int, 5)

	go producer(ch)
	go consumer(ch)

	time.Sleep(10 * time.Second)
}

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int, 1), make(chan int, 1)
	go func() {
		defer close(ch1)
		Walk(t1, ch1)
	}()
	go func() {
		defer close(ch2)
		Walk(t2, ch2)
	}()

	for v1 := range ch1 {
		v2 := <-ch2
		if v1 != v2 {
			return false
		}
	}
	return true
}

func TestSame(t *testing.T) {
	t1 := tree.New(1)
	t2 := tree.New(1)
	t3 := tree.New(2)

	assert.Equal(t, Same(t1, t2), true)
	assert.Equal(t, Same(t1, t3), false)
}
