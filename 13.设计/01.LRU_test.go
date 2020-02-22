package design

import (
	"fmt"
	"testing"
)

func TestCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.link.Print())
	fmt.Println(cache.Get(1))
	fmt.Println(cache.link.Print())
	cache.Put(3, 3)
	fmt.Println(cache.link.Print())
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println("4", cache.link.Print())
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
	fmt.Println(cache.link.Print())
}

func TestLRUCache_Get(t *testing.T) {
	type args struct {
		key   int
		value int
	}
	type wantArgs struct {
		key  int
		want int
	}
	tests := []struct {
		name  string
		cap   int
		args  []args
		wants []wantArgs
	}{
		{name: "1", cap: 2, args: []args{{1, 1}, {2, 2}, {3, 3}}, wants: []wantArgs{{1, -1}, {2, 2}}},
		{name: "2", cap: 3, args: []args{{1, 1}, {2, 2}, {3, 3}}, wants: []wantArgs{{1, 1}, {2, 2}, {3, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.cap)
			for _, arg := range tt.args {
				this.Put(arg.key, arg.value)
			}
			fmt.Println(this.link.Print())
			for _, want := range tt.wants {
				if got := this.Get(want.key); got != want.want {
					t.Errorf("Get() = %v, want %v", got, want.want)
				}
			}

		})
	}
}

func TestLRUCache_Put(t *testing.T) {
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name string
		cap  int
		args args
		want int
	}{
		{name: "1", cap: 2, args: args{key: 1, value: 1}, want: 1},
		{name: "2", cap: 2, args: args{key: 2, value: 2}, want: 2},
		{name: "3", cap: 2, args: args{key: 3, value: 3}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := Constructor(tt.cap)
			this.Put(tt.args.key, tt.args.value)
			if got := this.Get(tt.args.key); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestAddHead(t *testing.T) {
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name string
		args []args
		want string
	}{
		{name: "1", args: []args{{1, 1}, {2, 2}, {3, 3}}, want: "k:3, v:3=>k:2, v:2=>k:1, v:1"},
		{name: "2", args: []args{{1, 1}, {2, 2}}, want: "k:2, v:2=>k:1, v:1"},
		{name: "3", args: []args{{}}, want: "k:0, v:0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewLinked()
			for _, arg := range tt.args {
				this.AddHead(arg.key, arg.value)
			}
			if got := this.Print(); got != tt.want {
				t.Errorf("AddHead() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestAppend(t *testing.T) {
	type args struct {
		key   int
		value int
	}
	tests := []struct {
		name string
		args []args
		want string
	}{
		{name: "1", args: []args{{1, 1}, {2, 2}, {3, 3}}, want: "k:1, v:1=>k:2, v:2=>k:3, v:3"},
		{name: "2", args: []args{{1, 1}, {2, 2}}, want: "k:1, v:1=>k:2, v:2"},
		{name: "3", args: []args{{}}, want: "k:0, v:0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := NewLinked()
			for _, arg := range tt.args {
				this.Append(arg.key, arg.value)
			}
			if got := this.Print(); got != tt.want {
				t.Errorf("AddHead() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestRemoveNode(t *testing.T) {
	link := NewLinked()
	head := link.Append(1, 1)
	link.Append(2, 2)
	link.Append(3, 3)
	node4 := link.Append(4, 4)
	tail := link.Append(5, 5)
	fmt.Println(link.Print())
	link.RemoveNode(head)
	if link.Print() != "k:2, v:2=>k:3, v:3=>k:4, v:4=>k:5, v:5" {
		t.Error("remove head", link.Print())
	}
	link.RemoveNode(node4)
	if link.Print() != "k:2, v:2=>k:3, v:3=>k:5, v:5" {
		t.Error("remove middle4", link.Print())
	}
	link.RemoveNode(tail)
	if link.Print() != "k:2, v:2=>k:3, v:3" {
		t.Error("remove tail", link.Print())
	}
}

func TestRemoveTail(t *testing.T) {
	link := NewLinked()
	link.Append(1, 1)
	link.RemoveTail()
	link.RemoveTail()
	link.Append(1, 1)
	link.Append(1, 1)
	link.Append(1, 1)
	link.RemoveTail()
	link.RemoveTail()
	link.RemoveTail()
	link.RemoveTail()
}
func TestReverse(t *testing.T) {
	link := NewLinked()
	link.Append(1, 1)
	link.Append(2, 2)
	link.Append(3, 3)
	link.Append(4, 4)
	link.Append(5, 5)
	fmt.Println(link.Print())
	rev := link.Reverse()
	if link.PrintLink(rev) != "k:5, v:5=>k:4, v:4=>k:3, v:3=>k:2, v:2=>k:1, v:1" {
		t.Error("reverse is err")
	}

}
