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
	Name    string
	Age     int
	Gender  string
	Account int
	Email   string
	Address string
}

func (p *Person) SetAccountTo0(cmd *cobra.Command, args []string) {
	fmt.Printf("烂赌鬼 %s 账户余额为 %v\n", p.Name, p.Account)
}
