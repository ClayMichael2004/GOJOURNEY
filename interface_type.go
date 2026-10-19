package main

import "fmt"

type vehicle interface {
	getSpeed() float64
	getDistance() float64
}

type car struct {
	brand string
	year  int
}

func (Car car) getSpeed() float64 {
	return 120
}

func (Car car) getDistance() float64 {
	return 4000
}

func main() {
	var veh vehicle
	veh = car{
		"Lambo",
		2009,
	}
	fmt.Println(veh)
	fmt.Println("Speed:", veh.getSpeed())
	fmt.Println("Distance:", veh.getDistance())
}
