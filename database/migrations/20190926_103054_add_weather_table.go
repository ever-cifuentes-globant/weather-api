package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddWeatherTable_20190926_103054 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddWeatherTable_20190926_103054{}
	m.Created = "20190926_103054"

	migration.Register("AddWeatherTable_20190926_103054", m)
}

// Run the migrations
func (m *AddWeatherTable_20190926_103054) Up() {
	m.SQL(`CREATE TABLE weather (
		id SERIAL PRIMARY KEY,
		location_name varchar(255) NULL,
		temperature varchar(255) NULL,
		wind varchar(255) NULL,
		cloudines varchar(255) NULL,
		pressure varchar(255) NULL
	);`)
}

// Reverse the migrations
func (m *AddWeatherTable_20190926_103054) Down() {
	m.SQL("DROP TABLE weather")
}
