/*
	auth: satanfire
	date: 2023/6/14
	desc: mysql table to struct
*/

package repo

import (
	"database/sql"
	"fmt"
	"os"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/satanfire/super-kit/utils"
)

type mysqlIns struct {
	MysqlTypeToStructType map[string]string
	uriSchema             string

	structTpl string
}

type MysqlDb struct {
	Host     string
	UserName string
	Password string
	Charset  string
	DbName   string
}

// 列信息
type MysqlTableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

// 结构体的列信息
type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type MysqlTpl struct {
	TableName string
	Columns   []*StructColumn
}

var MysqlIns *mysqlIns

// 连接数据库
func (obj *mysqlIns) Conn(dbInfo *MysqlDb) (*sql.DB, error) {
	// 拼接访问地址
	uri := fmt.Sprintf(obj.uriSchema, dbInfo.UserName, dbInfo.Password,
		dbInfo.Host, dbInfo.Charset)
	return sql.Open("mysql", uri)
}

// 获取表的列信息
func (obj *mysqlIns) GetColumns(conn *sql.DB, dbName, tableName string) ([]*MysqlTableColumn, error) {
	query := `
		SELECT 
			COLUMN_NAME, DATA_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT
		FROM 
			information_schema.COLUMNS 
		WHERE 
			TABLE_SCHEMA = ? AND TABLE_NAME = ?
		`
	rows, err := conn.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []*MysqlTableColumn
	for rows.Next() {
		var column MysqlTableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnKey,
			&column.IsNullable, &column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}

		columns = append(columns, &column)
	}

	return columns, nil
}

// 将mysql数据表的列信息转换成struct结构
func (obj *mysqlIns) TransColumns(tbColumns []*MysqlTableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	// 遍历所有的列
	for _, column := range tbColumns {
		// 生成标签
		tag := fmt.Sprintf("`db:\"%s\"`", column.ColumnName)
		// 构造struct
		sc := &StructColumn{
			Name:    column.ColumnName,
			Type:    obj.MysqlTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		}
		tplColumns = append(tplColumns, sc)
	}
	return tplColumns
}

// 生成struct
func (obj *mysqlIns) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("mysqlToStruct").Funcs(template.FuncMap{
		"ToCamel": utils.UnderLineToCamel,
	}).Parse(obj.structTpl))

	tplDB := MysqlTpl{
		TableName: tableName,
		Columns:   tplColumns,
	}
	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	MysqlIns = new(mysqlIns)
	MysqlIns.uriSchema = "%s:%s@tcp(%s)/?charset=%s&parseTime=true"

	// 结构体模版
	MysqlIns.structTpl = `type {{.TableName | ToCamel}} struct {
{{range .Columns}}   {{$length := len .Comment}} {{if gt $length 0}}// {{.Comment}}{{else}}// {{.Name}}{{end}}
    {{$typeLen := len .Type}}{{if gt $typeLen 0}}{{.Name | ToCamel}} {{.Type}} {{.Tag}} {{else}}{{.Name}} {{end}}
{{end}}}
`

	// 数据类型映射表
	MysqlIns.MysqlTypeToStructType = map[string]string{
		"int":        "int",
		"tinyint":    "int",
		"smallint":   "int8",
		"mediumint":  "int64",
		"bigint":     "int64",
		"bit":        "int",
		"bool":       "bool",
		"enum":       "string",
		"set":        "string",
		"varchar":    "string",
		"char":       "string",
		"tinytext":   "string",
		"mediumtext": "string",
		"text":       "string",
		"longtext":   "string",
		"blob":       "string",
		"tinyblob":   "string",
		"mediumblob": "string",
		"longblob":   "string",
		"date":       "time.Time",
		"datetime":   "time.Time",
		"timestamp":  "time.Time",
		"time":       "time.Time",
		"float":      "float64",
		"double":     "float64",
		"decimal":    "float64",
	}
}
