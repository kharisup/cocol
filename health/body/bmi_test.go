package healthbody

import (
	"testing"
)

func TestBMI(t *testing.T) {

	//Table driven test
	var tests = []struct {
		height float64
		weight float64
		res    string
		exp    string
	}{
		{1.7, 30, "Underweight", "Underweight"},
		{1.7, 60, "Normal weight", "Normal weight"},
		{1.7, 80, "Overweight", "Overweight"},
		{1.7, 101.1, "Obesity", "Obesity"},
		{1.7, 101.2, "Severe obesity", "Severe obesity"},
	}

	for _, e := range tests {
		bmi, res := CalculateBMI(e.height, e.weight)
		if res != e.exp {
			t.Errorf("CalculateBMI(%v, %v) = \"%s\", bmi =  %v, expected %s", e.height, e.weight, res, bmi, e.exp)
		}
	}

}
