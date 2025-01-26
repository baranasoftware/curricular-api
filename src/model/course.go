package model

type Term int

const (
	Spring Term = iota
	Summer
	Fall
)

type Course struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Term     Term      `json:"term"`
	Credit   Credit    `json:"credit"`
	Teachers []Teacher `json:"teachers"`
}
