package output

import (
	"fmt"

	"drone-battery-estimator/model"
	"drone-battery-estimator/utils"
)

func Print(data model.Battery, time float64) {

	fmt.Println()

	fmt.Println("Flight Summary")

	fmt.Println("----------------")

	fmt.Printf("Charge      : %.0f%%\n", data.Charge)

	fmt.Printf("Temperature : %.1f °C\n", data.Temperature)

	fmt.Printf("Payload     : %.2f kg\n", data.Payload)

	fmt.Println()

	fmt.Println(

		"Estimated Time:",

		utils.Minutes(time),

	)

}
