package chi_for_test

import "testing"

func TestFor1(t *testing.T) {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	for {
		sum++
		if sum > 100 {
			break
		}
	}
	a := []int{1, 2, 3, 4, 5, 6}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	t.Log("sum: ", sum)
	t.Log("a: ", a)
}
