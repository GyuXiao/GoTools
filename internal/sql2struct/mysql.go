package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// 数据库连接的核心

type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

// 存储连接 MySQL 的一些基本信息

type DBInfo struct {
	DBType   string
	Host     string
	UserName string
	Password string
	Charset  string
}

// 存储 COLUMNS 表种我们需要的一些字段

type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

// 连接 MySQL 数据库

func (m *DBModel) Connect() error {
	var err error
	s := "%s:%s@tcp(%s)/information_schema?" +
		"charset=%s&parseTime=True&loc=Local"
	dataSourceName := fmt.Sprintf(
		s,
		m.DBInfo.UserName,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
	)
	// 这里的 DBType，默认是 mysql
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dataSourceName)
	if err != nil {
		return err
	}
	return nil
}

// 针对 COLUMNS 表进行查询和数据组装

func (m *DBModel) GetColumns(dbName, tableName string) ([]*TableColumn, error) {
	query := `SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY,
		IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT 
		FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ?`

	rows, err := m.DBEngine.Query(query, dbName, tableName)

	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("没有数据")
	}
	defer rows.Close()

	var columns []*TableColumn
	for rows.Next() {
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType, &column.ColumnKey, &column.IsNullable,
			&column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}
		columns = append(columns, &column)
	}
	//fmt.Println(columns)
	return columns, nil
}

// 表字段类型映射

var DBTypeToStructType = map[string]string{
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
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
}
