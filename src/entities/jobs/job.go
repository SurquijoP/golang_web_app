package entities

import "time"

type Job struct {
	Nombre string
	Delay  time.Duration
	Number int
}
