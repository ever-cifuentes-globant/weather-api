package models

import "github.com/astaxie/beego/orm"

type Weather struct {
	ID             int64  `orm:"column(id);pk;auto"`
	LocationName   string `orm:"column(location_name)"`
	Temperature    string `orm:"column(temperature)"`
	Wind           string `orm:"column(wind)"`
	Cloudines      string `orm:"column(cloudines)"`
	Pressure       string `orm:"column(pressure)"`
	Humidity       string `orm:"column(humidity)"`
	Sunrise        string `orm:"column(sunrise)"`
	Sunset         string `orm:"column(sunset)"`
	GeoCoordinates string `orm:"column(geo_coordinates)"`
	RequestedTime  string `orm:"column(requested_time)"`
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Weather))
}
