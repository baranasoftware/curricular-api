package model

import "time"

type Credit struct {
	credit float64 // provide utility methods without exposing basic type
}

type Class struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Credit   Credit    `json:"credit"`
	Location Address   `json:"location"`
	Time     time.Time `json:"time"`
}
