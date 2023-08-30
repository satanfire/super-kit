/*
	auth: satanfire
	date: 2023/8/30
	desc: ip地址转换
*/

package cmd

import (
	"fmt"

	"github.com/satanfire/super-kit/repo"
	"github.com/spf13/cobra"
)

type ipTransIns struct {
	ip         string
	ipToIntCmd *cobra.Command
}

var IpTransIns *ipTransIns

var ipTransCmd = &cobra.Command{
	Use:   "ipTrans",
	Short: "ip转换",
	Long:  "ip地址转换",
}

// ip to int
func (obj *ipTransIns) toInt(cmd *cobra.Command, args []string) {
	num, err := repo.IpTransIns.ToInt(obj.ip)
	if err != nil {
		fmt.Printf("trans failed, %s\n", err.Error())
		return
	}
	fmt.Printf("%s trans to int, res: \x1b[1;42m%d\x1b[0m\n", obj.ip, num)
	return
}

func init() {
	IpTransIns = new(ipTransIns)
	IpTransIns.ipToIntCmd = &cobra.Command{
		Use:   "toInt",
		Short: "IP to int",
		Long:  "IP addr trans to int",
		Run:   IpTransIns.toInt,
	}

	// 注册命令
	ipTransCmd.AddCommand(IpTransIns.ipToIntCmd)

	// 注册flag
	IpTransIns.ipToIntCmd.Flags().StringVarP(&IpTransIns.ip, "ip", "", "", "请输入ip地址")
}
