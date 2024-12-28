package model

type Term int

const (
	Spring Term = iota
	Summer
	Fall
)

type Course struct {
	Id          string
	Name        string
	Term        Term
	TotalCredit Credit
	Classes     []Class
}
