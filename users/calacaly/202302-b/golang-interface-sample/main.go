package main

import (
	"fmt"
	"test/duck1"
	"test/duck2"
	"test/model"
)


func main(){
	a := model.Person{Name: "鸭子"}
	b := model.Panda{Name: "鸭子"}

	fmt.Println("=====interface=====")

	duck1.WhoAreYou(a)
	duck1.Eating(a)

	
	duck1.WhoAreYou(b)
	duck1.Eating(b)

	fmt.Println("=====assert=====")
	

	duck2.WhoAreYou(a)
	duck2.Eating(a)

	duck2.WhoAreYou(b)
	duck2.Eating(b)

	fmt.Println("=====self=====")
	a.Me()
	a.Eat()
	a.Read()

	b.Me()
	b.Eat()
}