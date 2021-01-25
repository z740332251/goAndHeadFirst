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

func (wss *WeatherStationSubject) SetCurrentWeather() {
	wss.Weather.Humidity = strconv.Itoa(rand.Intn(100)) +"%"
	wss.Weather.Temperature = strconv.Itoa(rand.Intn(40)) + "℃"
	wss.NotifyAll()
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
