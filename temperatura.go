package main

import (
	"fmt"
)

const ebulicaoKelvin = 373.0

func main() {
	temperaturaCelsius := ebulicaoKelvin - float64(273)

	fmt.Printf("A temperatura de ebulição da água em °C = %g .", temperaturaCelsius)
}
