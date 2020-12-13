package ch9_string

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	//s[2]='2'//编译错误
	t.Log(len(s))
	s = "\xE4\xB8\xA5\x2B" //可以存储任意二进制数据
	t.Log(s)
	t.Log(len(s))
	s = "中"
	c := []rune(s) //rune是能够取出字符串里的Unicode  字符串转化为rune的切片
	t.Log(len(c))
	t.Logf("中 Unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

/*
Unicode 是一种字符集
UTF8 是unicode的存储实现

*/
func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}
