package ch6_array_test

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	arrLength := len(array)
	t.Log(arrLength)
	for i := 0; i < len(array); i++ {
		fmt.Println("Element", i, "of array is", array[i])
	}
	for i, v := range array {
		t.Log(i, v)
	}

}
func modify(array [5]int) {
	array[0] = 10
	fmt.Println("In modify array values:", array)
}
func TestModify(t *testing.T) {
	array := [5]int{1, 23, 4, 5, 6}
	modify(array) //传递给一个函数 并试图在函数体内修改这个数组内容
	t.Log("In main()array values:", array)
	//fmt.Println("In main()array values:",array)
}
