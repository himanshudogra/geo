package main

import (
	"fmt"
	"log"

	"github.com/himanshudogra/geo"
)

func main() {
	location := geo.Landmark{}

	err := location.SetName("Taj Mahal")
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLatitude(23.98)
	if err != nil {
		log.Fatal(err)
	}

	err = location.SetLongitude(-121.76)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name:", location.Name())
	fmt.Println("Latitude:", location.Latitude())
	fmt.Println("Longitude:", location.Longitude())

}
