package geo

import (
	"errors"
)

type Landmark struct {
	name string
	Coordinate
}

func (l *Landmark) SetName(name string) error {
	if name == "" {
		errors.New("invalid name")
	}
	l.name = name
	return nil
}

func (l *Landmark) Name() string {
	return l.name
}
