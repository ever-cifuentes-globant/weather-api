package mediators

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"

	"github.com/ever-cifuentes-globant/weather-api/models"
	"github.com/ever-cifuentes-globant/weather-api/viewmodels"
)

//const Continue bool = true

var (
	ErrInvalidCountryOrCity = errors.New("Country or City is not valid")
	ErrCannotUnmarshalBody  = errors.New("Cannot Unmarshal Body")
	ErrHttpGETWeather       = errors.New("Cannot GET Weather from External API")
)

type WeatherResponse struct {
	LocationName   string
	Temperature    string
	Wind           string
	Cloudines      string
	Pressure       string
	Humidity       string
	Sunrise        string
	Sunset         string
	GeoCoordinates string
	RequestedTime  string
}

type WindInterval struct {
	Tipo string
	Min  int
	Max  int
}

func GetWind(dep int) string {
	intervalos := []WindInterval{
		WindInterval{Tipo: "North", Min: 0, Max: 0},
		WindInterval{Tipo: "North-NorthEast", Min: 1, Max: 44},
		WindInterval{Tipo: "WeatherResponseNorthEast", Min: 45, Max: 45},
		WindInterval{Tipo: "East-NorthEast", Min: 46, Max: 89},
		WindInterval{Tipo: "East", Min: 90, Max: 90},
		WindInterval{Tipo: "East-SouthEast", Min: 91, Max: 134},
		WindInterval{Tipo: "SouthEast", Min: 135, Max: 135},
		WindInterval{Tipo: "South-SouthEast", Min: 136, Max: 179},
		WindInterval{Tipo: "South", Min: 180, Max: 180},
		WindInterval{Tipo: "South-SouthWest", Min: 181, Max: 224},
		WindInterval{Tipo: "South-West", Min: 225, Max: 225},
		WindInterval{Tipo: "West-SouthWest", Min: 226, Max: 269},
		WindInterval{Tipo: "West", Min: 270, Max: 270},
		WindInterval{Tipo: "west-NorthWest", Min: 271, Max: 314},
		WindInterval{Tipo: "NorthWest", Min: 315, Max: 315},
		WindInterval{Tipo: "North-NorthWest", Min: 316, Max: 359},
		WindInterval{Tipo: "North", Min: 360, Max: 360},
	}

	for _, value := range intervalos {
		if dep == value.Min || dep > value.Min && dep < value.Max {
			return value.Tipo
		}
	}
	return ""
}

func getSunrise(JSONParsed viewmodels.WeatherJSON) string {
	i, err := strconv.ParseInt(strconv.Itoa((JSONParsed.APISys.Sunrise)), 10, 64)
	if err != nil {
		panic(err)
	}
	tm1 := time.Unix(i, 0)
	sunrise := strconv.Itoa(tm1.Hour()) + ":" + strconv.Itoa(tm1.Minute())
	return sunrise
}

func getCloudines(JSONParsed viewmodels.WeatherJSON) string {
	var cloudines string
	for _, v := range JSONParsed.APIWeather {
		cloudines = v.Description
	}
	return cloudines
}

func getSunset(JSONParsed viewmodels.WeatherJSON) string {
	j, err := strconv.ParseInt(strconv.Itoa((JSONParsed.APISys.Sunset)), 10, 64)
	if err != nil {
		panic(err)
	}
	tm2 := time.Unix(j, 0)
	sunset := strconv.Itoa(tm2.Hour()) + ":" + strconv.Itoa(tm2.Minute())
	return sunset
}

func getRequestedTime() string {
	t := time.Now()
	RequestedTime := (t.Format(time.RFC3339))
	return RequestedTime
}

func init() {
	_ = orm.RegisterDriver("postgres", orm.DRPostgres)

	_ = orm.RegisterDataBase("default",
		"postgres",
		"user=postgres password=postgres host=pq port=5432 dbname=postgres sslmode=disable")
}

func getLocation(city, country string) string {
	//Genero la locacion para buscar en la tabla
	Location := strings.Title(city) + "," + strings.ToUpper(country)
	return Location
}

func getWeather(city, country string) (*models.Weather, error) {
	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "," + country + "&appid=1508a9a4840a5574c822d70ca2132032")
	if err != nil {
		return nil, ErrHttpGETWeather
	}

	JSONParsed, err := viewmodels.NewWeatherResponse(response.Body)
	if err != nil {
		return nil, ErrCannotUnmarshalBody
	}

	if JSONParsed.City != city {
		return nil, ErrInvalidCountryOrCity
	}
	var wr models.Weather
	wr.LocationName = JSONParsed.City + "," + JSONParsed.APISys.Country
	wr.Pressure = strconv.Itoa(int(JSONParsed.APIMain.Pressure)) + " hpa"
	wr.Humidity = strconv.Itoa((JSONParsed.APIMain.Humidity)) + "%"
	wr.Sunrise = getSunrise(*JSONParsed)
	wr.Sunset = getSunset(*JSONParsed)
	wr.Temperature = strconv.Itoa(int(JSONParsed.APIMain.Temp-273.15)) + " °C"
	wr.Wind = strconv.FormatFloat((JSONParsed.APIWind.Speed), 'f', 2, 64) + " m/s" + ", " + GetWind(int(JSONParsed.APIWind.Deg))
	wr.GeoCoordinates = "[" + strconv.FormatFloat((JSONParsed.APICoords.Lon), 'f', 2, 64) + "," + strconv.FormatFloat((JSONParsed.APICoords.Lat), 'f', 2, 64) + "]"
	wr.Cloudines = getCloudines(*JSONParsed)
	wr.RequestedTime = getRequestedTime()
	return &wr, nil
}

func IsDataUpdated(wt models.Weather) bool {
	var result int

	str := wt.RequestedTime
	t1, _ := time.Parse(time.RFC3339, str)

	aux := time.Now()
	RequestedTime := (aux.Format(time.RFC3339))
	t2, _ := time.Parse(time.RFC3339, RequestedTime)

	result = int(t2.Sub(t1).Seconds())

	return (result < 10)
}

func Update(wt, new *models.Weather) {
	wt.Temperature = new.Temperature
	wt.Wind = new.Wind
	wt.Cloudines = new.Cloudines
	wt.Pressure = new.Pressure
	wt.Humidity = new.Humidity
	wt.Sunrise = new.Sunrise
	wt.Sunset = new.Sunset
	wt.GeoCoordinates = new.GeoCoordinates
	wt.RequestedTime = new.RequestedTime
}

func IsOnBD(city, country string) (models.Weather, error) {

	//Inicializo DB
	db := orm.NewOrm()
	_ = db.Using("default")

	//Busco en BD si existe
	wt := models.Weather{LocationName: getLocation(city, country)}
	err := db.Read(&wt, "LocationName")
	return wt, err
}

func InsertOnBD(city, country string) (*models.Weather, error) {

	//Inicializo DB
	db := orm.NewOrm()
	_ = db.Using("default")

	//Genero la informacion del clima
	wr, err := getWeather(city, country)

	if err != nil {
		return nil, err
	}

	//Inserto en BD - Nuevo campo
	_, _ = db.Insert(wr)
	return wr, nil
}

func ApiAndUpdate(city, country string, wt models.Weather) (*models.Weather, error) {

	//Inicializo DB
	db := orm.NewOrm()
	_ = db.Using("default")

	new, err := getWeather(city, country)
	if err != nil {
		return nil, err
	}
	//Actualizo Valores en wt
	Update(&wt, new)
	//Actualizo DB_, _
	_, _ = db.Update(&wt)
	return &wt, nil
}

func GetWeatherResponse(city, country string) (*models.Weather, error) {

	wt, err := IsOnBD(city, country)

	if err == orm.ErrNoRows {
		//Nuevo Registro
		return InsertOnBD(city, country)
	}

	if IsDataUpdated(wt) {
		//Devuelvo valor de la BD
		return &wt, nil
	}

	//External API + UPDATE
	return ApiAndUpdate(city, country, wt)

}

func UpdateWithTime(numberOfUpdates int, city, country string) {

	for i := 0; i < numberOfUpdates; i++ {
		time.Sleep(2 * time.Second)
		//Actualizar y Actualizar
		//External API + UPDATE
		fmt.Println("UPDATE")
		wt, _ := IsOnBD(city, country)
		_, _ = ApiAndUpdate(city, country, wt)
	}
}
