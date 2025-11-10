package main

import "fmt"

func main() {
	const mileInKM = 1.60934
	const marathonInKM = 42.195

	marathonInMiles := marathonInKM / mileInKM
	fmt.Printf("A marathon is %.2f kilometres = %.2f miles long\n", marathonInKM, marathonInMiles)

	boilingWaterCelsius := 100.0
	boilingWaterFahrenheit := (boilingWaterCelsius * 9.0 / 5.0) + 32.0
	fmt.Printf("Water boils at %.2f°C = %.2f°F\n", boilingWaterCelsius, boilingWaterFahrenheit)
}
