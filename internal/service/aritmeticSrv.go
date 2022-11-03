package service

import (
	"strings"
)

type ArithmeticSrv interface {
	Compute(operation string, x, y int) (ops string, result int)
}

type arthSrv struct{}

func (a *arthSrv) Compute(operation string, x, y int) (ops string, result int) {

	split := strings.Split(operation, " ")
	for _, word := range split {
		if word == "add" || word == "sum" || word == "total" {
			operation = "addition"
			break
		} else if word == "minus" || word == "subtraction" || word == "decrease" {
			operation = "subtraction"
			break
		} else if word == "multiplication" || word == "subtraction" || word == "decrease" {
			operation = "multiplication"
			break
		}

	}

	switch operation {
	case "addition":
		return "addition", x + y
	case "subtraction":
		return "subtraction", x - y
	case "multiplication":
		return "multiplication", x * y
	default:
		return "", 0
	}
}

func NewArthSrv() ArithmeticSrv {
	return &arthSrv{}
}
