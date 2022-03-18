package model

// DbInfo 数据库基础信息
type DbInfo struct {
	Version   string
	Charset   string
	Collation string
	DbName    string
}

// DbConfig 数据库配置
type DbConfig struct {
	DbType    int // 1. mysql  2. oracle 3. mssql
	DocType   int // 1. online 3. offline
	Host      string
	Port      int
	User      string
	Password  string
	Database  string
	Sid       string
	TableName string //表名通配符，如果为空默认全部表
}

// Column info
type Column struct {
	ColName    string
	ColType    string
	ColKey     string
	IsNullable string
	ColComment string
	ColDefault string
}

// Table info
type Table struct {
	TableName    string
	TableComment string
	ColList      []Column
}
