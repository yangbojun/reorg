package services

import (
	"database/sql"
	"log"
)

// 定义函数执行碎片整理语句
func ExecuteSql(db *sql.DB, sql string, logger *log.Logger) {
	// 执行sql语句
	_, err := db.Exec(sql)
	if err != nil {
		logger.Printf("[ERROR]执行%s语句出错:%v", sql, err)
	} else {
		logger.Printf("[INFO]%s执行完成", sql)
	}
}