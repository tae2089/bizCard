package repository

import (
	"bizCard/ent"
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
)

var Client *ent.Client
var once sync.Once

//client singletone 생성 필요
func OpenDB() *ent.Client {
	if Client == nil {
		once.Do(func() {
			db, err := sql.Open("mysql", "root:test@tcp(localhost:13306)/bizcard?parseTime=true")
			defer func(db *sql.DB) {
				err := db.Close()
				if err != nil {
					panic(err)
				}
			}(db)
			if err != nil {
				panic(err)
			}
			db.SetMaxIdleConns(10)
			db.SetMaxOpenConns(100)
			db.SetConnMaxLifetime(time.Hour)
			drv := entsql.OpenDB("mysql", db)
			Client = ent.NewClient(ent.Driver(drv))
		})
	}
	return Client
}
