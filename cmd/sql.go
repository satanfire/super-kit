/*
	auth: satanfire
	date: 2023/6/14
	desc: mysql 相关工具集合
*/

package cmd

import (
	"fmt"

	"github.com/satanfire/super-kit/repo"
	"github.com/spf13/cobra"
)

type mysqlIns struct {
	host             string
	username         string
	password         string
	charset          string
	dbName           string
	tableName        string
	mysqlToStructCmd *cobra.Command
}

var MysqlIns *mysqlIns

var mysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "mysql 转换",
	Long:  "mysql 转换",
	Run:   func(cmd *cobra.Command, args []string) {},
}

// 业务逻辑
func (obj *mysqlIns) mysqlToStruct(cmd *cobra.Command, args []string) {
	// 获取数据库连接
	dbInfo := &repo.MysqlDb{
		Host:     obj.host,
		UserName: obj.username,
		Password: obj.password,
		Charset:  obj.charset,
		DbName:   obj.dbName,
	}
	conn, err := repo.MysqlIns.Conn(dbInfo)
	if err != nil {
		fmt.Printf("repo.MysqlIns.Conn, failed, %s\n", err.Error())
		return
	}

	// 查询数据表信息
	columns, err := repo.MysqlIns.GetColumns(conn, obj.dbName, obj.tableName)
	if err != nil {
		fmt.Printf("repo.MysqlIns.GetColumns, failed, %s\n", err.Error())
		return
	}

	// 通过模版转换
	sc := repo.MysqlIns.TransColumns(columns)
	err = repo.MysqlIns.Generate(obj.tableName, sc)
	if err != nil {
		fmt.Printf("repo.MysqlIns.Generate, failed, %s\n", err.Error())
		return
	}
	return
}

func init() {
	MysqlIns = new(mysqlIns)
	// mysql表转struct
	MysqlIns.mysqlToStructCmd = &cobra.Command{
		Use:   "struct",
		Short: "mysql表转go struct",
		Long:  "mysql表转go struct",
		Run:   MysqlIns.mysqlToStruct,
	}

	// 注册自命令
	mysqlCmd.AddCommand(MysqlIns.mysqlToStructCmd)

	// 注册flag
	MysqlIns.mysqlToStructCmd.Flags().StringVarP(&MysqlIns.username, "username", "u", "root", "请输入数据库的账号")
	MysqlIns.mysqlToStructCmd.Flags().StringVarP(&MysqlIns.password, "password", "p", "root", "请输入数据库的密码")
	MysqlIns.mysqlToStructCmd.Flags().StringVarP(&MysqlIns.host, "host", "", "127.0.0.1:3306", "请输入数据库的HOST:PORT")
	MysqlIns.mysqlToStructCmd.Flags().StringVarP(&MysqlIns.charset, "charset", "c", "utf8mb4", "请输入数据库的编码")
	MysqlIns.mysqlToStructCmd.Flags().StringVarP(&MysqlIns.dbName, "db", "d", "information_schema", "请输入数据库名称")
	MysqlIns.mysqlToStructCmd.Flags().StringVarP(&MysqlIns.tableName, "table", "t", "", "请输入表名称")
}
