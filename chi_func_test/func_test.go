package chi_func_test

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	a = 1 << iota
	b
	c
	d
)

func TestIotaDemo(t *testing.T) {
	t.Log(a)
	t.Log(b)
	t.Log(c)
	t.Log(d)
	t.Log(7 ^ 13) //0111 1011 1100 结果是10
	t.Log("A")
	t.Log("C")
	t.Log("A" < "C")
}

func Add(a int, b int) (ret int, err error) {
	if a < 0 || b < 0 { //假设这个函数只支持两个非负数字的加法
		err = errors.New("Should be non-negative numbers!")
		return

	}
	return a + b, nil
}

func TestMainLoop(t *testing.T) {
	v, err := Add(1, 2)
	if err == nil {
		t.Log(v)
	}
	//t.Log(Add(1,2))
}
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Print("time spenr:", time.Since(start).Seconds())
		return ret
	}
}
func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}
func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSFF := timeSpent(slowFun)
	t.Log(tsSFF(10))
}
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}
func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 5, 4))
}
func Clear() {
	fmt.Println("Clear resource")
}
func TestDefer(t *testing.T) {
	defer Clear() //函数返回前会执行这个函数 类似于finally 常用与释放锁 关闭资源

	t.Log("Start")
	panic("Fatal error") //defer仍然会执行
	fmt.Println("123")
}
