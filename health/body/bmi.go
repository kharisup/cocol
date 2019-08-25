package healthbody

//CalculateBMI is moethod to calcule Body Mass Index
// BMI = Weight * Height * Height
func CalculateBMI(height float64, weight float64) (float64, string) {

	bmi := weight / (height * height)
	status := ""

	switch {
	case bmi < 18.5:
		status = "Underweight"
	case bmi >= 18.5 && bmi < 24.9:
		status = "Normal weight"
	case bmi >= 25 && bmi < 29.9:
		status = "Overweight"
	case bmi >= 30 && bmi < 35:
		status = "Obesity"
	case bmi >= 35:
		status = "Severe obesity"
	default:
		status = ""
	}

	return bmi, status
}
