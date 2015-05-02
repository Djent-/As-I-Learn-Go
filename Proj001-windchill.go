package main

import "math"
import "fmt"

func main() {
	fmt.Print("Temperature (degrees F): ")
	var temp float64
	_, err := fmt.Scanln(&temp)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Wind speed (MPH): ")
	var wind float64
	_, err = fmt.Scanln(&wind)
	if err != nil {
		fmt.Println(err)
	}

	wcti := 35.74 + 0.6215 * temp - 35.75 * math.Pow(wind, 0.16) + 0.4275 * temp * math.Pow(wind, 0.16)
	fmt.Printf("Wind chill is %0.2f\n", wcti)
}
