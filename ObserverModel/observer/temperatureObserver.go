package observer

import (
	"GoHeadFirst/ObserverModel/data"
	"fmt"
)

/**
 * Auther: 不在东墙
 * Descptrion: 温度观察者
 * Date: 2021/1/23
**/
type TemperatureObserver struct {
}

func (t TemperatureObserver) Update(weather data.Weather) {
	fmt.Println("温度是： " + weather.Temperature)
}