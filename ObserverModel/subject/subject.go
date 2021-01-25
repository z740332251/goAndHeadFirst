package subject

import "GoHeadFirst/ObserverModel/observer"

type Subject interface {
	Register(observer observer.Observer)
	Unregister(observer observer.Observer)
	NotifyAll()
}
