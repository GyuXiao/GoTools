package sql2struct

import (
	"fmt"
	"goTools/internal/word"
	"os"
	"text/template"
)

// 预定义模板
const structTpl = `type {{.TableName | ToCamelCase}} struct {
{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
	{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
{{end}}}

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}`

type StructTemplate struct {
	structTpl string
}

// 存储转换后的 Go 结构体中的所有字段信息

type StructColumn struct {
	Name string
	Type string
	Tag string
	Comment string
}

// 存储最终用于渲染的模板对象信息

type StructTemplateDB struct {
	TableName string
	Columns []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

// 模板渲染
// 数据库类型到 Go 结构体的转换和 JSON Tag 的处理

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tag := fmt.Sprintf("`"+"json:"+"%s"+"`", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

// 先声明一个名为 sql2struct 的新模板对象，接着自定义函数 ToCamelCase，用上 word.UnderscoreToUpperCamelCase 方法进行绑定
// 最后组装符合预定义模板的模板对象，再调用 Execute 方法进行渲染

func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.structTpl))
	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	err := tpl.Execute(os.Stdout, tplDB)
	if err != nil {
		return err
	}
	return nil
}