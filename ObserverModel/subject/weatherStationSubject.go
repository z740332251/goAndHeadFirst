package subject

import (
	"GoHeadFirst/ObserverModel/data"
	"GoHeadFirst/ObserverModel/observer"
	"math/rand"
	"strconv"
)

/**
 * Auther: 不在东墙
 * Descptrion: 气象站主题
 * Date: 2021/1/23
**/


type WeatherStationSubject struct {
	Weather *data.Weather
	observers []observer.Observer
}
func NewWeatherStationSubject() *WeatherStationSubject{
	wss := new(WeatherStationSubject)
	wss.Weather = new(data.Weather)
	return wss
}

func getWeather(weather data.Weather) data.Weather{
	/**
	 * TODO(不在东墙): 待删除
	**/
	weather.Humidity = strconv.FormatFloat(float64(rand.Float32()),'E',-1,32)
	weather.Temperature = strconv.FormatFloat(float64(rand.Float32()),'E',-1,32)
	return weather
}
func (wss *WeatherStationSubject) SetCurrentWeather() {
	wss.Weather.Humidity = strconv.Itoa(rand.Intn(100)) +"%"
	wss.Weather.Temperature = strconv.Itoa(rand.Intn(40)) + "℃"
}
func (wss *WeatherStationSubject) Register(observer observer.Observer)  {
	wss.observers = append(wss.observers, observer)
}
func (wss *WeatherStationSubject) Unregister(observer observer.Observer)  {

}
func (wss *WeatherStationSubject) notify(observer observer.Observer)  {
	observer.Update(*wss.Weather)
}
func (wss WeatherStationSubject) NotifyAll()  {
	for _, o := range wss.observers {
		wss.notify(o)
	}
}