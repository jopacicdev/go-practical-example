package main

import (
	"fmt"
)


type car struct {
	gasPedal uint16 // 0-65535
	brakePedal uint16
	steeringWheel int16 // -32k - 32k
	topSpeedKmh float64
}

func main() {
	aCar := car {
		gasPedal: 22341,
		brakePedal: 0,
		steeringWheel: 12461,
		topSpeedKmh: 250.0}
	
	fmt.Println(aCar.gasPedal)
}