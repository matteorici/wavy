package calculator

func PayloadFactor(weight float64) float64 {

	factor := 1.0 - weight*0.12

	if factor < 0.35 {

		return 0.35

	}

	return factor

}
