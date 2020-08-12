package main

import (
	"fmt"
	"log"

	"github.com/himanshudogra/geo"
)

func main() {
	coordinates := geo.Coordinate{}

	err := coordinates.SetLatitude(23.98)
	if err != nil {
		log.Fatal(err)
	}

	err = coordinates.SetLongitude(-121.76)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Latitude:", coordinates.Latitude())
	fmt.Println("Longitude:", coordinates.Longitude())

}
