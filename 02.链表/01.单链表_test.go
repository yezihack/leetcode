package linked

import (
	"fmt"
	"testing"
)

func TestSLinked_AddHead(t *testing.T) {
	link := NewSLinked()
	link.AddHead(1)
	link.AddHead(2)
	link.AddHead(3)
	var want string
	want = "3=>2=>1"
	if link.Print() != want {
		t.Errorf("Append()=%s, want:%s", link.Print(), want)
	}
}

func TestSLinked_Append(t *testing.T) {
	link := NewSLinked()
	link.Append(1)
	link.Append(2)
	link.Append(3)
	var want string
	want = "1=>2=>3"
	if link.Print() != want {
		t.Errorf("Append()=%s, want:%s", link.Print(), want)
	}
}

func TestSLinked_RemoveHead(t *testing.T) {
	link := NewSLinked()
	link.Append(1)
	link.Append(2)
	link.Append(3)
	link.RemoveHead()
	var want string
	want = "2=>3"
	if link.Print() != want {
		t.Errorf("Append()=%s, want:%s", link.Print(), want)
	}
	link.RemoveHead()
	want = "3"
	if link.Print() != want {
		t.Errorf("Append()=%s, want:%s", link.Print(), want)
	}
	link.RemoveHead()
	want = ""
	if link.Print() != want {
		t.Errorf("Append()=%s, want:%s", link.Print(), want)
	}
}

func TestSLinked_RemoveTail(t *testing.T) {
	link := NewSLinked()
	link.Append(1)
	link.Append(2)
	link.Append(3)
	link.RemoveTail()
	var want string
	want = "1=>2"
	if link.Print() != want {
		t.Errorf("Append()=%s, want:%s", link.Print(), want)
	}
	link.RemoveTail()
	want = "1"
	if link.Print() != want {
		t.Errorf("Append()=%s, want:%s", link.Print(), want)
	}
	link.RemoveTail()
	want = ""
	if link.Print() != want {
		t.Errorf("Append()=%s, want:%s", link.Print(), want)
	}
}
func TestSLinked_RemoveNode(t *testing.T) {
	link := NewSLinked()
	s1 := link.Append(1)
	s2 := link.Append(2)
	s3 := link.Append(3)
	_ = s1 //2=>3
	_ = s2 //1=>3
	_ = s3 //1=>2
	link.RemoveNode(s3)
	//link.RemoveNode(s1)
	//link.RemoveNode(s2)
	fmt.Println(link.Print())
}
