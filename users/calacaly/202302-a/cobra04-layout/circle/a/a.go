package a

import "cobra04/circle/b"

func A() string {
	return "a"
}

func AB() string {
	return A() + b.B()
}
