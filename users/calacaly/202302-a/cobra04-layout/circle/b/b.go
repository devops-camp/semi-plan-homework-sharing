package b

import "cobra04/circle/a"

func B() string {
	return "b"
}

func BA() string {
	return B() + a.A()
}
