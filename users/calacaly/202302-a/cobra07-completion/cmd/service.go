/*
Copyright © 2023 calacaly

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var serviceAlias []string

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "service",
	Long:  `service`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("service %s created .\n", strings.Join(args, " "))
	},
}

var nodePortCmd = &cobra.Command{
	Use:   "nodePort",
	Short: "create a NodePort service.",
	Long:  `create a NodePort service.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("service %s created .\n", strings.Join(args, " "))
	},
}

var loadbalancerCmd = &cobra.Command{
	Use:   "loadbalancer",
	Short: "create a LoadBalancer service.",
	Long:  `create a LoadBalancer service.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("service %s created .\n", strings.Join(args, " "))
	},
}

var externalnameCmd = &cobra.Command{
	Use:   "externalname",
	Short: "Create an ExternalName service.",
	Long:  `Create an ExternalName service.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("service %s created .\n", strings.Join(args, " "))
	},
}

var clusteripCmd = &cobra.Command{
	Use:   "clusterip",
	Short: "Create a ClusterIP service.",
	Long:  `Create a ClusterIP service.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("service %s created .\n", strings.Join(args, " "))
	},
}

func init() {
	createCmd.AddCommand(serviceCmd)

	serviceAlias = append(serviceAlias, "svc")

	serviceCmd.Aliases = serviceAlias

	serviceCmd.AddCommand(clusteripCmd)
	serviceCmd.AddCommand(nodePortCmd)
	serviceCmd.AddCommand(externalnameCmd)
	serviceCmd.AddCommand(loadbalancerCmd)

}
