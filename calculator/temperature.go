package calculator

import "drone-battery-estimator/config"

func TemperatureFactor(temp float64) float64 {

	if temp < 5 {

		return config.ColdPenalty

	}

	if temp > 35 {

		return config.HotPenalty

	}

	return 1.0
}
