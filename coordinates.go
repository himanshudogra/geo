package geo

import (
	"errors"
)

type coordinate struct {
	latitude  float64
	longitude float64
}

func (c *coordinate) SetLatitude(latitude float64) error {
	if latitude < -90 || latitude > 90 {
		return errors.New("invalid latitude.")
	}
	c.latitude = latitude
	return nil
}

func (c *coordinate) SetLongitude(longitude float64) error {
	if longitude < -180 || longitude > 180 {
		return errors.New("invalid longitude.")
	}
	c.longitude = longitude
	return nil
}

func (c *coordinate) Latitude() float64 {
	return c.latitude
}

func (c *coordinate) Longitude() float64 {
	return c.longitude
}
