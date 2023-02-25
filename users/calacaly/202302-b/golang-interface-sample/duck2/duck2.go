package duck2

import "test/model"

//通过断言实现鸭子对象
func Eating(v any){
	switch v.(type){
	case model.Person:
		p := v.(model.Person)
		p.Eat()
	case model.Panda:
		p := v.(model.Panda)
		p.Eat()
	}
}
func WhoAreYou(v any){
	switch v.(type){
	case model.Person:
		p := v.(model.Person)
		p.Me()
	case model.Panda:
		p := v.(model.Panda)
		p.Me()
	}
}
//通过断言实现鸭子对象