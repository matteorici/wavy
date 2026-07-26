package input

import (
	"bufio"
	"fmt"
	"os"

	"drone-battery-estimator/model"
)

func Read() model.Battery {

	reader := bufio.NewReader(os.Stdin)

	var b model.Battery

	fmt.Print("Battery (%): ")
	fmt.Fscan(reader, &b.Charge)

	fmt.Print("Temperature (°C): ")
	fmt.Fscan(reader, &b.Temperature)

	fmt.Print("Payload (kg): ")
	fmt.Fscan(reader, &b.Payload)

	return b

}
