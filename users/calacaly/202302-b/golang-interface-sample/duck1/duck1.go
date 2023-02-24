package duck1

type Who interface{
	Eat()
	Me()
}

//通过接口实现鸭子对象


func Eating(w Who){
	w.Eat()
}
func WhoAreYou(w Who){
	w.Me()
}
//通过接口实现鸭子对象

