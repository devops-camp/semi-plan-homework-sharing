package model

import(
	"fmt"
)


type Person struct{
	Name string
}

func (p Person)Eat(){
	fmt.Println("我爱吃大米")
}

func (p Person)Read(){
	fmt.Println("我会读书")
}

func (p Person)Me(){
	fmt.Printf("我是人, 长得像%s\n",p.Name)
}

type Panda struct{
	Name string
}

func (p Panda)Eat(){
	fmt.Println("我爱吃竹子")
}

func (p Panda)Me(){
	fmt.Printf("我是熊猫, 长得像%s\n",p.Name)
}