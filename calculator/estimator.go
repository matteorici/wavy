package calculator

import (
	"drone-battery-estimator/config"
	"drone-battery-estimator/model"
)

func Estimate(data model.Battery) float64 {

	time := config.BaseFlightTime

	time *= data.Charge / 100

	time *= TemperatureFactor(data.Temperature)

	time *= PayloadFactor(data.Payload)

	return time

}
