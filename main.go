package main

import (
	"fmt"

	"github.com/chibby0ne/micro_neural_network/matrix"
	"gonum.org/v1/hdf5"
)

const (
	// filename of hdf5 file containing training set
	trainingSetFile string = "datasets/train_catvnoncat.h5"
	// filename of hdf5 file containing test set
	testSetFile string = "datasets/test_catvnoncat.h5"
	// name of the input dataset inside the hdf5 training set
	inputTest string = "test_set_x"
	// name of the output dataset inside the hdf5 training set
	outputTestSet string = "test_set_y"
	// name of the input dataset inside the hdf5 training set
	inputTrainingSet string = "train_set_x"
	// name of the output dataset inside the hdf5 training set
	outputTrainingSet string = "train_set_y"
	// Number of training instances
	m int = 209
	// pixels per dimension. Each picture is 64 x 64
	pixelsPerDimension int = 64
	// Each pixel is coded in RGB, therefore contains 3 values
	colorsChannel int = 3
)

// InputArray type is used to contain the input to the neuron.
// Bounds and types are taken by opening and examining the datasets in jupyter notebooks
type InputArray [m][pixelsPerDimension][pixelsPerDimension][colorsChannel]uint8

// OutputArray type is used to contain the input to the neuron.
// Bounds and types are taken by opening and examining the datasets in jupyter notebooks
type OutputArray [1][m]int64

func main() {
	// Create dummy matrix
	matrix, err := matrix.NewMatrix(3, 3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(matrix)

	// open hdf5 file
	f, err := hdf5.OpenFile(trainingSetFile, hdf5.F_ACC_RDONLY)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// Open the input training dataset
	dset, err := f.OpenDataset(inputTrainingSet)
	if err != nil {
		fmt.Println(err)
	}
	defer dset.Close()

	// read all the training examples
	var inputs InputArray
	dset.Read(&inputs)
	if err != nil {
		fmt.Println(err)
	}

	matrix2 := inputs.ToNumberArray()

	// print only the first pixel of the first training example
	fmt.Println(inputs[0][0][0])
	val1, _ := matrix2.GetValue(0, 0)
	val2, _ := matrix2.GetValue(0, 1)
	val3, _ := matrix2.GetValue(0, 2)
	fmt.Println(val1, val2, val3)

	// Open the output training dataset
	dset, err = f.OpenDataset(outputTrainingSet)
	if err != nil {
		fmt.Println(err)
	}
	defer dset.Close()

	// read all the training examples
	var outputs OutputArray
	dset.Read(&outputs)
	if err != nil {
		fmt.Println(err)
	}

	matrix3 := outputs.ToNumberArray()

	// print only the first pixel of the first training example
	fmt.Println(outputs[0][0], outputs[0][1], outputs[0][2])
	val1, _ = matrix3.GetValue(0, 0)
	val2, _ = matrix3.GetValue(0, 1)
	val3, _ = matrix3.GetValue(0, 2)
	fmt.Println(val1, val2, val3)

}

// ToNumberArray converts the InputArray into a NumberArray
func (inputArray *InputArray) ToNumberArray() matrix.NumberArray {
	matrix, _ := matrix.NewMatrix(m, pixelsPerDimension*pixelsPerDimension*colorsChannel)
	for i := 0; i < m; i++ {
		for j := 0; j < pixelsPerDimension; j++ {
			for k := 0; k < pixelsPerDimension; k++ {
				for l := 0; l < colorsChannel; l++ {
					matrix.SetValue(i,
						j*pixelsPerDimension*colorsChannel+k*colorsChannel+l,
						float64(inputArray[i][j][k][l]))
				}
			}
		}
	}
	return matrix
}

// ToNumberArray converts the OutputArray into a NumberArray
func (outputArray *OutputArray) ToNumberArray() matrix.NumberArray {
	matrix, _ := matrix.NewMatrix(1, m)
	for i := 0; i < m; i++ {
		matrix.SetValue(0, i, float64(outputArray[0][i]))
	}
	return matrix
}
