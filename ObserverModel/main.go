package main
import (
	observer "GoHeadFirst/ObserverModel/observer"
	"GoHeadFirst/ObserverModel/subject"
)
func main() {
	weatherStationSubject := subject.NewWeatherStationSubject()
	temperatureObserver := new(observer.TemperatureObserver)
	weatherStationSubject.Register(temperatureObserver)

	humidityObserver := new(observer.HumidityObserver)
	weatherStationSubject.Register(humidityObserver)

	weatherStationSubject.SetCurrentWeather()

	weatherStationSubject.SetCurrentWeather()

	weatherStationSubject.SetCurrentWeather()


}