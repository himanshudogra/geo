package main

import (
	"fmt"
	"log"

	"github.com/himanshudogra/geo"
)

func main() {
	coordinates := geo.coordinate{}

	err := coordinates.SetLatitude(23.98)
	if err != nil {
		log.Fatal(err)
	}

	err = coordinates.SetLongitude(-121.76)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Latitude:", geo.Latitude())
	fmt.Println("Longitude:", geo.Longitude())

}
