package dataaccess

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"

	"eroica.finance/src/eroica/finance/conf"
	"eroica.finance/src/eroica/finance/entity"
	"eroica.finance/src/eroica/finance/tushare"
	_ "github.com/go-sql-driver/mysql"
)

var dbGlobal *sql.DB
var initDbOnce sync.Once

func initDb() *sql.DB {
	initDbOnce.Do(func() {
		dbCfg := conf.GetConf().Db
		db, err := sql.Open(dbCfg.Driver, dbCfg.DataSource) //自带连接池
		if err != nil {
			panic(err)
		}
		err = db.Ping()
		if err != nil {
			panic(err)
		}
		db.SetMaxOpenConns(dbCfg.MaxOpenConns)
		db.SetMaxIdleConns(dbCfg.MaxIdleConns)
		log.Println("[INFO] Init db success.")
		dbGlobal = db
	})
	return dbGlobal
}

func CloseDb() {
	if dbGlobal != nil {
		dbGlobal.Close()
	}
}

func SelectTxless(sql string, params ...interface{}) (columns []string, rows [][]*string) {
	db := initDb()
	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	rs, err := stmt.Query(params...)
	if err != nil {
		panic(err)
	}
	defer rs.Close()
	columns, err = rs.Columns()
	if err != nil {
		panic(err)
	}
	rows = make([][]*string, 0, 1)
	for rs.Next() {
		row := make([]*string, len(columns))
		//TODO 以下待修改
		switch len(columns) {
		case 1:
			err = rs.Scan(&row[0])
		case 2:
			err = rs.Scan(&row[0], &row[1])
		default:
			panic("NOT SUPPORTED")
		}
		if err != nil {
			panic(err)
		}
		rows = append(rows, row)
	}
	return
}

func SaveTxless(em entity.EntityMapping, data tushare.TsRespData) {
	db := initDb()
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		panic(err)
	}
	keys := "(" + strings.Join(data.Fields, ", ") + ")"
	values := "(" + strings.Repeat("?, ", len(data.Fields)-1) + "?)"
	valuesAll := strings.Repeat(values+", ", len(data.Items)-1) + values
	stmt, err := tx.Prepare(fmt.Sprintf("insert into %s %s values %s", em.Table, keys, valuesAll))
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	defer stmt.Close()
	l := len(data.Items) * len(data.Items[0])
	argsAll := make([]interface{}, 0, l)
	for _, item := range data.Items {
		argsAll = append(argsAll, item...)
	}
	_, err = stmt.Exec(argsAll...)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
	log.Printf("[INFO] Save %v rows to %v success.\n", len(data.Items), em.Table)
}

func DeleteTxless(em entity.EntityMapping, whereClause string, params ...interface{}) sql.Result {
	db := initDb()
	tx, err := db.Begin()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		panic(err)
	}
	stmt, err := tx.Prepare(fmt.Sprintf("delete from %s %s", em.Table, whereClause))
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	defer stmt.Close()
	rs, err := stmt.Exec(params...)
	if err != nil {
		tx.Rollback()
		panic(err)
	}
	tx.Commit()
	rows, _ := rs.RowsAffected()
	log.Printf("[INFO] Delete %v rows from %v success.\n", rows, em.Table)
	return rs
}

func DeleteAllTxless(em entity.EntityMapping) sql.Result {
	return DeleteTxless(em, "")
}
