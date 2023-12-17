package infrastructure

import (
	"database/sql"
	"fmt"
	"github.com/xvbnm48/go-clean-project/src/interfaces"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteHandler struct {
	Conn *sql.DB
}

type MysqlHandler struct {
	Conn *sql.DB
}

func (handler *MysqlHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *MysqlHandler) Query(statement string) interfaces.Row {
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		fmt.Println(err)
		return new(MysqlRow)
	}
	row := new(MysqlRow)
	row.Rows = rows
	return row

}

func (handler *SqliteHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *SqliteHandler) Query(statement string) interfaces.Row {
	//fmt.Println(statement)
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		fmt.Println(err)
		return new(SqliteRow)
	}
	row := new(SqliteRow)
	row.Rows = rows
	return row
}

type SqliteRow struct {
	Rows *sql.Rows
}

type MysqlRow struct {
	Rows *sql.Rows
}

func (r MysqlRow) Scan(dest ...interface{}) {
	r.Rows.Scan(dest...)
}

func (r MysqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqliteRow) Scan(dest ...interface{}) {
	r.Rows.Scan(dest...)
}

func (r SqliteRow) Next() bool {
	return r.Rows.Next()
}

func NewSqliteHandler(dbfileName string) *SqliteHandler {
	conn, _ := sql.Open("sqlite3", dbfileName)
	sqliteHandler := new(SqliteHandler)
	sqliteHandler.Conn = conn
	return sqliteHandler
}

func NewMysqlHandler(dbUsername, dbPassword, dbName string) (*MysqlHandler, error) {
	dataSource := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", dbUsername, dbPassword, dbName)
	conn, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}
	mysqlHandler := new(MysqlHandler)
	mysqlHandler.Conn = conn
	return mysqlHandler, nil
	//dataSource := dbUsername + ":" + dbPassword + "@" + dbName
	//conn, err := sql.Open("mysql", dataSource)
	//if err != nil {
	//	return nil, err
	//}
	//
	//err = conn.Ping()
	//if err != nil {
	//	conn.Close()
	//	return nil, err
	//}
	//
	//return conn, nil
}
