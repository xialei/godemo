package demo

import (
	"database/sql"
	"fmt"
	"godemo/core"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//MysqlDb sql.DB 管理底层数据库连接的打开和关闭操作，管理数据库连接池
var MysqlDb *sql.DB
var MysqlDbErr error

const (
	UserName = "root"
	Password = "xxxx"
	HOST     = "xxxx" //inner ip
	PORT     = "3306"
	DATABASE = "xxx"
	CHARSET  = "utf8"
)

func processData() {
	var wg = sync.WaitGroup{}

	stmt, _ := MysqlDb.Prepare("select entity_id, entity_name, legal_rep_name, legal_rep_type, reg_province, rec_cap, found_date, entity_kind, logo from base_entity_basic_info where id BETWEEN ? AND ? and z_update_time >= ?")
	defer stmt.Close()

	each := 700000
	total := 70000000 / each
	g := core.New(20)
	updateTime := time.Unix(time.Now().AddDate(0, 0, -1).Unix(), 0).Format("2006-01-02 00:00:00")
	fmt.Printf("update time: %s\n", updateTime)

	for i := 0; i < total; i++ {
		wg.Add(1)
		index := i
		goFunc := func() {
			// do business
			fmt.Printf("go func: %d\n", index)
			startID := each*index + 1
			endID := each * (index + 1)
			rows, err := stmt.Query(startID, endID, updateTime)
			if err != nil {
				fmt.Printf("query data err: %d\n", index)
			}
			comp := new(EntityBasic)
			for rows.Next() {
				err = rows.Scan(&comp.EntityID, &comp.Name, &comp.LegalRep, &comp.LegalType, &comp.Province, &comp.Cap, &comp.Found, &comp.Type, &comp.Logo)
				if err != nil {
					fmt.Printf("get data error %s \n", err.Error())
				}
			}
			err = rows.Err()
			if err != nil {
				fmt.Printf(err.Error())
			}
			wg.Done()
		}
		g.Run(goFunc)
	}
	wg.Wait()
}

//DBReadPerformance test performance of operating db
func DBReadPerformance() {

	// mysql data source name
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", UserName, Password, HOST, PORT, DATABASE, CHARSET)

	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)
	defer MysqlDb.Close()

	if MysqlDbErr != nil {
		panic("error connection to db: " + MysqlDbErr.Error())
	}

	fmt.Println("start time: ", time.Now().Format("2006-01-02 15:04:05"))
	processData()
	fmt.Println("finish time: ", time.Now().Format("2006-01-02 15:04:05"))

}
