package model

type Result struct {
	Type   string
	Length int
	Value  string
}

func NewResult(typ string, length int, value string) *Result {
	return &Result{typ, length, value}
}
