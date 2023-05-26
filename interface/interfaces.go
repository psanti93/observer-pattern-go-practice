package main

type Observer interface {
	Update()
}

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}

type Display interface {
	Display()
}
