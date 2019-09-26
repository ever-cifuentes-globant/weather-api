package database

import (
	"github.com/astaxie/beego/orm"

	_ "github.com/lib/pq"
)

func Init() {
	_ = orm.RegisterDriver("postgres", orm.DRPostgres)
	_ = orm.RegisterDataBase("default",
		"postgres",
		"user=postgres password=postgres host=pq port=5432 dbname=weather_api_pq sslmode=disable")

}
