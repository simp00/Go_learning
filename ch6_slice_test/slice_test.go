package ch6_slice_test

import (
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
	t.Log(len(s2), cap(s2))

}
func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknow"
	t.Log(Q2)
	t.Log(year)

}

//数组容量是否可以伸缩
//是否可以比较
func TestSliceCompare(t *testing.T) {
	//a:=[]int{}
	//b:=[]int{}
	//if a==b {
	//	t.Log("a==b")
	//
	//}
}
func TestSliceArray(t *testing.T) {
	var myArray [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//基于数组创建一个数组切片
	var mySlcie []int = myArray[:5]
	t.Log("Elements of myArray: ")
	for _, v := range myArray {
		t.Log(v)
	}
	t.Log("Elements of mySlice")
	for _, v := range mySlcie {
		t.Log(v)
	}
}
func TestSlice2(t *testing.T) {
	mySlice := []int{1, 2, 3}
	mySlice2 := []int{1, 2, 3, 4, 5}
	t.Log(len(mySlice), cap(mySlice))
	t.Log(len(mySlice2), cap(mySlice2))
	mySlice = append(mySlice, mySlice2...) //如果没有这个省略号的话 会编译出错 因为按照append的语义从第二个参数起的所有参数都是待附加的元素 加上省略号相当于把mySlice2包含的所有元素打散后传入
	//oldSlice:=[]int{1,2,3,4,5}
	//newSlice :=oldSlice[:3]//基于oldSlice的前三个元素构建新数组切片
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1) //只复制slice1的前三个元素到slice2中
	copy(slice1, slice2) //只会复制slice2的3个元素到slice1的前三个位置

}
