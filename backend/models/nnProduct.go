package models

func CreateInputProduct(products []Product) [][]float64 {
	input := make ([][]float64, 0)
	for i := 0; i < len(products) ; i++ {
		input = append(input, []float64 {
			float64(products[i].cathegory), products[i].price })
	}
	return input
}

func CreateTargetProduct(persons []Person, product Product, visitors []Visitor) [][]float64 {
	output := make ([][]float64, 0)
	for i := 0; i < len(persons); i++ {
		buf := FindProductInPerson(i, product.id, visitors)
		output = append(output, []float64 {
			convertBoolToFloat64(buf), convertBoolToFloat64(!buf)})
	}
	return output
}
