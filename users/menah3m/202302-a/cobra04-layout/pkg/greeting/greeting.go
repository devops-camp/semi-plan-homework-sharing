package greeting

import (
	"fmt"
	"github.com/spf13/cobra"
)

/*
   @Auth: menah3m
   @Desc:
*/

type Person struct {
	Name string
	Age  int
}

func (p *Person) Greet(cmd *cobra.Command, args []string) {
	fmt.Printf("%s 你好，今年 %v 岁\n", p.Name, p.Age)
}
