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

func (p Person)Me(){
	fmt.Printf("我是人, 名字叫%s\n",p.Name)
}

type Panda struct{
	Name string
}

func (p Panda)Eat(){
	fmt.Println("我爱吃竹子")
}

func (p Panda)Me(){
	fmt.Printf("我是熊猫, 名字叫%s\n",p.Name)
}