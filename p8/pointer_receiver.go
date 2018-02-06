package main

import (
	"fmt"
)

const usixteenbitmax float64 = 65535
const kmhMultiple float64 = 1.60934

type car struct {
	gasPedal      uint16 // 0-65535
	brakePedal    uint16
	steeringWheel int16 // -32k - 32k
	topSpeedKmh   float64
}

func (c car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKmh / usixteenbitmax / kmhMultiple)
}

func (c *car) newTopSpeed(newSpeed float64) {
	c.topSpeedKmh = newSpeed
}

func main() {
	aCar := car{
		gasPedal:      22341,
		brakePedal:    0,
		steeringWheel: 12461,
		topSpeedKmh:   250.0}

	fmt.Println(aCar.gasPedal)
	fmt.Println(aCar.kmh())
	fmt.Println(aCar.mph())
	aCar.newTopSpeed(325.66)
	fmt.Println(aCar.kmh())
}
