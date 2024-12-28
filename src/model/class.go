package model

import "time"

type Credit struct {
	credit float64 // provide utility methods without exposing basic type
}

type Class struct {
	Id       string
	Name     string
	Credit   Credit
	Location Address
	Time     time.Time
}
