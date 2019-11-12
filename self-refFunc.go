package main

import (
	"fmt"
)

// 策略与机制分离(seperation of mechanism and policy)


type Foo struct {
	v int
}
type option func(*Foo) option

// mechanism
func (f *Foo) Option(opts ...option) option {
	var pre option
	for _, opt := range opts {
		pre = opt(f)
	}
	return pre
}

// policy
func Verbosity(v int) option {
	return func(f *Foo) option {
		pre := f.v
		f.v = v
		return Verbosity(pre)
	}
}
func main() {
	f := Foo{
		v: 3,
	}
	pre := f.Option(Verbosity(2))
	fmt.Println(f.v)
	f.Option(pre)
	fmt.Println(f.v)
}
