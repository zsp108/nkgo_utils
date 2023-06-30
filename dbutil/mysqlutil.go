package dbutil

import (
	"database/sql"
	"fmt"
	"log"
)

type dbMsg struct {
	db        *sql.DB
	ipAddress string
	userName  string
	password  string
	port      int
	dbName    string
	charset   string
}

func initDB(dbinfo *dbMsg) (err error) {
	dbConInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", dbinfo.userName, dbinfo.password, dbinfo.ipAddress, dbinfo.port, dbinfo.dbName, dbinfo.charset)
	dbinfo.db, err = sql.Open("mysql", dbConInfo)
	if err != nil {
		return err
	}
	err = dbinfo.db.Ping()
	if err != nil {
		log.Printf("校验失败！%v\n", err)
		return err
	}
	log.Println("连接数据库成功！")
	return nil
}

func ConnDB(ipaddr, username, password, dbName, charset string, port int) (dci dbMsg) {
	// var dci dbConnInfo
	dci.ipAddress = ipaddr
	dci.userName = username
	dci.password = password
	dci.dbName = dbName
	dci.charset = charset
	dci.port = port
	return

}

func ExecSQL(dbmsg dbMsg, sqlStr string) (affectrow int64, err error) {
	initDB(&dbmsg)
	defer dbmsg.db.Close()
	result, err := dbmsg.db.Exec(sqlStr)

	if err != nil {
		log.Printf("%v \n执行失败！，错误信息：%v", sqlStr, err)
	}
	// 统计影响行数
	affectrow, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return
}
