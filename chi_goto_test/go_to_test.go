package chi_goto_test

import "testing"

func TestGoto(t *testing.T) {
	i := 0
HERE:
	t.Log(i)
	i++
	if i < 10 {
		goto HERE
	}
}
