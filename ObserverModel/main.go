package main
import (
	observer "GoHeadFirst/ObserverModel/observer"
	"GoHeadFirst/ObserverModel/subject"
)
func main() {
	weatherStationSubject := subject.NewWeatherStationSubject()
	weatherStationSubject.SetCurrentWeather()
	temperatureObserver := new(observer.TemperatureObserver)
	weatherStationSubject.Register(temperatureObserver)

	humidityObserver := new(observer.HumidityObserver)
	weatherStationSubject.Register(humidityObserver)

	weatherStationSubject.NotifyAll()
}