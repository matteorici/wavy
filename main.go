package main

import (
	"drone-battery-estimator/calculator"
	"drone-battery-estimator/input"
	"drone-battery-estimator/output"
)

func main() {

	data := input.Read()

	time := calculator.Estimate(

		data,

	)

	output.Print(

		data,

		time,

	)

	output.Save(

		time,

	)

}
