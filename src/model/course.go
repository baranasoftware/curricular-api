package model

type Term int

const (
	Spring Term = iota
	Summer
	Fall
)

type Course struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Term        Term    `json:"term"`
	TotalCredit Credit  `json:"totalCredit"`
	Classes     []Class `json:"classes"`
}
