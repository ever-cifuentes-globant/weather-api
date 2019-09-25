package database

import (
	"github.com/astaxie/beego/orm"

	_ "github.com/lib/pq"
)

func Init() {
	_ = orm.RegisterDriver("postgres", orm.DRPostgres)
	_ = orm.RegisterDataBase("default",
		"postgres",
		"user=postgres password=postgres host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")

}
