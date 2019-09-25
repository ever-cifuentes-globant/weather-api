package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"

	"github.com/ever-cifuentes-globant/weather-api/mediators"
	"github.com/ever-cifuentes-globant/weather-api/viewmodels"
)

// Operations about weather
type WeatherController struct {
	beego.Controller
}

// @Title GetWeather
// @Description get country and city
// @Param	idcity		path 	string	true		"The City"
// @Param	idcountry	path 	string	true		"The Country"
// @Success 200 {Object} models.Weather
// @Failure 403 :city not exist
// @router /city/:idcity/country/:idcountry [get]
func (c *WeatherController) GetWeather() {
	defer c.ServeJSON()

	//Recupero city y country
	city := c.GetString(":idcity")
	country := c.GetString(":idcountry")

	//Llamo al Mediators para pedir el clima
	wr, err := mediators.GetWeatherResponse(city, country)
	if err != nil {
		c.Data["json"] = viewmodels.NewErrorResponse(err)
		return
	}
	c.Data["json"] = wr
}

// @Title UpdateWeather
// @Description update the weather
// @Param	body		body 	viewmodels.WeatherPersist	true		"body for weather content"
// @Failure 403 :city not exist
// @router /scheduler/weather [put]
func (c *WeatherController) UpdateWeather() {
	fmt.Print("Entra")
	defer c.ServeJSON()

	//Recupero data del Request
	var JSONParsed viewmodels.WeatherPersist
	_ = json.NewDecoder(c.Ctx.Request.Body).Decode(&JSONParsed)

	//Compruebo si existe en la BD
	wr, err := mediators.IsOnBD(JSONParsed.City, JSONParsed.Country)
	if err == orm.ErrNoRows {
		fmt.Println("enntra al err")
		c.Data["json"] = viewmodels.NewErrorResponse(err)
		return
	}

	go mediators.UpdateWithTime(10, JSONParsed.City, JSONParsed.Country)

	c.Data["json"] = wr

}
