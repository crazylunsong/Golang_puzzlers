package article8

import (
	"container/list"
	"container/ring"
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	var container1 list.List

	t.Log(container1.Len())
	back := container1.PushBack(1)
	container1.PushBack(2)
	container1.PushBack("xia")

	t.Log(container1.Len())
	container1.MoveToBack(back)
	front := container1.Front()
	t.Log(front.Value)
	for i := container1.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

func TestRing(t *testing.T) {
	r := ring.New(10)
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}
	for i := 0; i < r.Len(); i++ {
		r = r.Prev()
		t.Log(r.Value)
	}
	//r.Do(func(i interface{}) {
	//	t.Log(i)
	//})
}

func TestRing_Move(t *testing.T) {
	r := ring.New(5)
	n := r.Len()
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	r = r.Move(3)
	r.Do(func(i interface{}) {
		t.Log(i)
	})
}
