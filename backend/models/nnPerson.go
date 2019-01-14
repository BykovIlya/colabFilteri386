package models

import (
	"github.com/fxsjy/gonn/gonn"
	"time"
	"fmt"
	"os"
)

func convertBoolToFloat64 (in bool) float64 {
	out := 0.0
	if (in) {
		out = 1.0
	}
	return out
}

var input, target [][]float64

func CreateInputPerson(persons []Person, product Product, visitors  [] Visitor) ([][]float64, [][]float64) {
	input := make ([][]float64, 0)
	output := make ([][]float64, 0)
	for i := 0; i < len(persons) ; i++ {
			input = append(input, []float64 {
				convertBoolToFloat64(persons[i].gender), float64(persons[i].age), float64(product.cathegory), product.price})
		buf := FindProductInPerson(i, product.id, visitors)
		output = append(output, []float64 {
			convertBoolToFloat64(buf), convertBoolToFloat64(!buf)})
	}
	return input, output
}

func quickAppendIn(el []float64) {
	input = append(input,el)
}

func quickAppendOut(el []float64) {
	target = append(target, el)
}

func quickAppend(j int, in, o [][]float64) {
		 quickAppendIn(in[j])
		 quickAppendOut(o[j])
}

func CreateNeuralNetworkPerson(persons []Person, product []Product, visitors []Visitor, need bool) {
	if (need) {
		nn := gonn.DefaultNetwork(4, 10, 2, false)
		input = make([][]float64, 0)
		target = make([][] float64, 0)
		if (IsEmptyInputNN() && IsEmptyTargetNN()) {
			for i := 0; i < 1000 /*len(product)*/ ; i++ {
				start := time.Now()
				in, o := CreateInputPerson(persons, product[i], visitors)
				for j := 0; j < /*len(persons)*/ 1000; j++ {
					quickAppend(j, in, o)
				}
				t := time.Now()
				fmt.Println(i, ":", t.Sub(start))
			}
			ImportInputNNToDB(input)
			ImportTargetNNToDB(target)
			nn.Train(input, target, 100) //minimum
			gonn.DumpNN("gonnPerson", nn)
		} else {
			input = GetInputNNFromDB()
			target = GetTargetNNFromDB()
			nn.Train(input, target, 100) //minimum
			gonn.DumpNN("gonnPerson", nn)
		}
	} else {
		fmt.Println("nothing to do in NN")
	}
}

func GetResult(output [] float64) int64{
	max := -99999.0
	pos := -1

	for i, val := range output {
		if (val > max) {
			max = val
			pos = i
		}
	}

	switch pos {
	case 0: return 1
	case 1: return 0

	}
	return -1
}

func deleteFile(file string) {
	var err = os.Remove(file)
	if isError(err) { return }

	fmt.Println("old gonn fille deleted")
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}