package model

import (
	"database/sql"
	"fmt"
	_ "github.com/thda/tds"
)

// 定义数据连接信息结构体
type Db struct {
	IP string
	Port string
	Dbname string
	User string
	Password string
}

// 定义测试连接信息
func (this *Db) TestConn() error {
	ctx := fmt.Sprintf("tds://%s:%s@%s:%s/%s?charset=utf8",
		this.User, this.Password, this.IP, this.Port, this.Dbname)
	db, err := sql.Open("tds", ctx)
	if err != nil {
		return err
	}
	defer db.Close()
	//通过db.Ping检查数据库连接地址
	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

// 定义获取数据库连接
func (this *Db) Conn() *sql.DB {
	ctx := fmt.Sprintf("tds://%s:%s@%s:%s/%s?charset=utf8",
		this.User, this.Password, this.IP, this.Port, this.Dbname)
	db, _ := sql.Open("tds", ctx)
	return db
}