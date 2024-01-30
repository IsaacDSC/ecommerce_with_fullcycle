package config

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var db *sql.DB
var once sync.Once

func DbAnaLyticsConn() *sql.DB {
	once.Do(func() {
		var err error
		if db, err = sql.Open("postgres", "postgresql://root:root@localhost:5432/fullcycle_ecommerce?sslmode=disable"); err != nil {
			log.Panic(err)
		}
		db.SetMaxOpenConns(100)
		db.SetMaxIdleConns(20)
	})
	return db
}

func DbAnaLyticsClose() {
	db.Close()
}
