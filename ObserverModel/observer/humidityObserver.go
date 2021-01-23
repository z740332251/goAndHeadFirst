package observer

import (
	"GoHeadFirst/ObserverModel/data"
	"fmt"
)

/**
 * Auther: 不在东墙
 * Descptrion: 湿度观察者
 * Date: 2021/1/23
**/
type HumidityObserver struct {
}

func (h HumidityObserver) Update(weather data.Weather) {
	fmt.Println("湿度是： " + weather.Humidity)
}
