package output

import (
	"fmt"
	"os"
)

func Save(minutes float64) {

	file, _ := os.Create(

		"flight_report.txt",

	)

	defer file.Close()

	fmt.Fprintf(

		file,

		"Estimated Flight Time: %.1f minutes\n",

		minutes,

	)

}
