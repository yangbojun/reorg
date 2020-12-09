package services

import (
	"database/sql"
	"log"
)

// 定义生成碎片整理语句函数
func GenerateCompact(db *sql.DB, channel chan string, logger *log.Logger) {
	defer close(channel)
	// 生成碎片整理语句sql
	generate_sql := `select 'reorg compact '+name from sysobjects where type='U'
	UNION all
	SELECT 'reorg rebuild '+sysobjects.name+' '+sysindexes.name FROM sysobjects,sysindexes
	WHERE sysindexes.indid>0 AND sysobjects.type='U' AND sysindexes.id=sysobjects.id AND sysindexes.indid<255
	UNION all
	select 'UPDATE STATISTICS '+name from sysobjects where type='U'`

	// 生成语句
	rows, err := db.Query(generate_sql)
	if err != nil {
		logger.Printf("[ERROR]获取碎片整理语句失败:",err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var sql string
		rows.Scan(&sql)
		channel <- sql
	}
}

// 定义生成碎片整理语句函数
func GenerateRebuild(db *sql.DB, channel chan string, logger *log.Logger) {
	defer close(channel)
	// 生成碎片整理语句sql
	generate_sql := `select 'reorg rebuild '+name from sysobjects where type='U'
	UNION all
	select 'UPDATE STATISTICS '+name from sysobjects where type='U'`

	// 生成语句
	rows, err := db.Query(generate_sql)
	if err != nil {
		logger.Printf("[ERROR]获取碎片整理语句失败:",err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var sql string
		rows.Scan(&sql)
		channel <- sql
	}
}