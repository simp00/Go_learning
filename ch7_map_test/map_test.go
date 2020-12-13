package ch7_map_test

import (
	"testing"
)

func TestMap(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m)
	t.Log(m["two"])
	m2 := map[int]int{}
	t.Log(len(m2))
	m3 := make(map[int]int, 10)
	t.Log("len m3=", len(m3))

}
func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])
	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Log("key 3's value is", v)
	} else {
		t.Log("key 3 is not existing")

	}
}
func TestTravelMap(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for k, v := range m {
		t.Log(k, v)
	}
}

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func TestMapPersonInfo(t *testing.T) {
	//var personDB map_ext[string]PersonInfo
	personDB := make(map[string]PersonInfo) //与上面的定义语句 语义重复
	//往map中插入几条数据
	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}
	//从map中查找键值为“1234“的信息
	person, ok := personDB["1234"]
	if ok {
		t.Log("Found person", person.Name, "with ID 1234")
	} else {
		t.Log("Did not find person with ID 1234")
	}
	delete(personDB, "1")

}
