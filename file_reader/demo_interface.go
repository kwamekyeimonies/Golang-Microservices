package filereader

import (
	"math"
)

type Polynomial interface {
	Roots_Determinant() float64
}

type Roots struct {
	A float64
	B float64
	C float64
}

type Root_Value struct {
	Y string
}

func (determinant Roots) Roots_Determinant() float64 {
	x := (math.Pow(determinant.B, 2) - (4 * (determinant.A * determinant.C)))
	return x
}

func Roots_Checker(root_var float64) string {

	var string_output string

	if root_var < 0 {
		string_output = "The Equation Has Complex Roots"
		return string_output
	} else if root_var == 0 {
		string_output = "Equation has equal Roots"
		return string_output
	} else {
		string_output = "Equation has Real Roots"
		return string_output
	}
}
