package observer

import "GoHeadFirst/ObserverModel/data"

/**
 * Auther: 不在东墙
 * Descptrion: 观察者接口
 * Date: 2021/1/23
**/
type Observer interface {
	Update(weather data.Weather)
}