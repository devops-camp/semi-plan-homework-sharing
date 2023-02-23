package cmd

const (
	INFO = "Student info: "
	LINE = "========================================================================================================================================"
)

type Addr struct {
	Home   string `flag:"home" usage:"home address"`
	School string `flag:"school" usage:"school address"`
}

type Contcat struct {
	Email string `flag:"email" usage:"email address"`
	Phone string `flag:"phone" usage:"phone"`
}

type Student struct {
	Age     int     `flag:"age" usage:"student age"`
	Addr    Addr    `flag:"addr"`
	Contcat Contcat `flag:"contact"`
	Gender  bool    `flag:"gender"`
	Name    string  `flag:"name" usage:"student name"`
}

var student *Student

func GetStudent() *Student {
	if student == nil {
		student = &Student{
			Name: "zhangsan",
			Addr: Addr{
				Home: "Sichuan,China",
			},
			Contcat: Contcat{
				Phone: "138-8888-8888",
			},
		}
	}
	return student
}
