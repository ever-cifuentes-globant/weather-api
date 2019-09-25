package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/ever-cifuentes-globant/weather-api/controllers:WeatherController"] = append(beego.GlobalControllerRouter["github.com/ever-cifuentes-globant/weather-api/controllers:WeatherController"],
        beego.ControllerComments{
            Method: "GetWeather",
            Router: `/city/:idcity/country/:idcountry`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/ever-cifuentes-globant/weather-api/controllers:WeatherController"] = append(beego.GlobalControllerRouter["github.com/ever-cifuentes-globant/weather-api/controllers:WeatherController"],
        beego.ControllerComments{
            Method: "UpdateWeather",
            Router: `/scheduler/weather`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
