package repository

import (
	"bizCard/ent"
	"context"
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
	"time"
)

var BizCardRepositoryBean BizCardRepository
var UserRepositoryBean UserRepository
var Client *ent.Client
var onceEnt sync.Once

//client singletone 생성 필요
func OpenDB() *ent.Client {
	if Client == nil {
		onceEnt.Do(func() {
			db, err := sql.Open("mysql", "root:secret@tcp(localhost:13306)/bizcard?parseTime=true")
			if err != nil {
				log.Println(err)
				Client.Close()
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

func RegisterRepositoryBeans() {
	OpenDB()
	CreateSchema()
	SetupUserRepository()
	SetupBizCardRepository()
}

func CreateSchema() {
	// Run the auto migration tool.
	if err := Client.Schema.Create(context.Background()); err != nil {
		Client.Close()
		panic(err)
	}
}
