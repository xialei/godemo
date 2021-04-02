package core

import (
	"database/sql"
	"fmt"
	"sync"
	"time"
)

//MysqlDb sql.DB 管理底层数据库连接的打开和关闭操作，管理数据库连接池
var (
	mysqlPool = sync.Map{}
)

const (
	UserName        = "root"
	Password        = "Ysbu2!nDiwsv"
	HOST            = "47.101.185.137" //inner ip
	PORT            = "3306"
	DATABASE        = "z_pe"
	CHARSET         = "utf8"
	DateTimeLayout  = "2006-01-02 15:04:05"
	DateTimeLayout2 = "2006-01-02 00:00:00"
	DateLayout      = "2006-01-02"
)

func GetDBConnection() (*sql.DB, error) {

	// mysql data source name
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", UserName, Password, HOST, PORT, DATABASE, CHARSET)

	client, ok := mysqlPool.Load(dbDSN)

	if !ok {
		client, err := sql.Open("mysql", dbDSN)

		if err != nil {
			fmt.Printf("error connection to db: " + err.Error())
			return nil, err
		}
		mysqlPool.Store(dbDSN, client)
		return client, nil
	}
	return client.(*sql.DB), nil
}

func getDateTime(day int) string {
	return time.Unix(time.Now().AddDate(0, 0, day).Unix(), 0).Format(DateTimeLayout2)
}
