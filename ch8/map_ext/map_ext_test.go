package map_ext

import (
	"testing"
)

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	t.Log(m)
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](5))
	t.Log(m[2](5))
	t.Log(m[3](5))
}
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	mySet[2] = true
	mySet[3] = true
	n := 5
	if mySet[n] {
		t.Logf("%d is exsiting", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	t.Log(len(mySet))
}
