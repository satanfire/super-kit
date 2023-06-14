package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Exec() error {
	return rootCmd.Execute()
}

func init() {
	// 添加子命令
	rootCmd.AddCommand(mysqlCmd)
}
