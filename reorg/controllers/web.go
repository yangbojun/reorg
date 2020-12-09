package controllers

import (
	"fmt"
	"net/http"
	"html/template"
	"reorg/mylog"
	"reorg/model"
	"reorg/services"
	"strconv"
	"time"
)

func WebServer(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("template/index.html", "template/success.html")
	if err != nil {
		fmt.Println("加载模板错误:", err)
	}
	if r.Method != "POST" {
		web := model.NewWeb()
		tpl.ExecuteTemplate(w, "index.html", web)
	} else {
		// 记录页面提交数据
		web := model.NewWeb()
		web.IP = r.FormValue("ip")
		web.Port = r.FormValue("port")
		web.Dbname = r.FormValue("dbname")
		web.User = r.FormValue("user")
		web.Password = r.FormValue("passwd")
		if r.FormValue("mod") == "1" {
			web.IsCompact = "checked"
			web.IsRebuild = ""
		} else {
			web.IsCompact = ""
			web.IsRebuild = "checked"
		}
		web.Parrel = r.FormValue("parrel")

		// 测试数据库配置是否正确
		connerr := web.TestConn()
		if connerr != nil {
			web.Error = fmt.Sprintf("数据库连接错误%v", connerr)
			tpl.ExecuteTemplate(w, "index.html", web)
		} else {
			// 获取并发数
			parrel, err := strconv.Atoi(web.Parrel)
			if err != nil {
				web.ParrelInfo = fmt.Sprintf("不合规并发数配置")
				tpl.ExecuteTemplate(w, "index.html", web)
			}

			// 新建日志输出
			logfile, logger := mylog.NewLogger(web.IP + "_" + web.Port + "_" + web.Dbname)

			// 开启工作池
			channel := make(chan string, parrel)

			// 获取数据库连接
			db := web.Conn()

			// 开启单独协程写入待矫正sql
			if web.IsCompact == "checked" {
				go services.GenerateCompact(db, channel, logger)
			} else {
				go services.GenerateRebuild(db, channel, logger)
			}

			exitCount := 0
			// 执行矫正语句
			for i := 1 ; i <= parrel; i++ {
				go func(th int, exitCount *int) {
					for {
						sql, notfinish := <- channel
						if !notfinish {
							*exitCount ++
							break
						}
						logger.Printf("[INFO]正在使用协程-%v执行:%v", th, sql)
						services.ExecuteSql(db, sql, logger)
					}
				} (i, &exitCount)
			}

			// 完成后关闭日志文件句柄和数据库句柄
			go func(exitCount *int) {
				for *exitCount < parrel {
					time.Sleep(time.Second)
				}
				logger.Print("本次任务完成！")
				logfile.Close()
				db.Close()
			} (&exitCount)

			tpl.ExecuteTemplate(w, "success.html", web)
		}
	}
}