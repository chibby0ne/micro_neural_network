package main

import (
	"fmt"
	"os"

	"github.com/chibby0ne/micro_neural_network/matrix"
	"gonum.org/v1/hdf5"
)

const (
	// filename of hdf5 file containing training set
	trainingSetFile string = "datasets/train_catvnoncat.h5"
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

func main() {
	// Create dummy matrix
	matrix, err := matrix.NewMatrix(3, 3)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println(matrix)

	// open hdf5 file
	f, err := hdf5.OpenFile(trainingSetFile, hdf5.F_ACC_RDONLY)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// Open the input training dataset
	dset, err := f.OpenDataset(inputTrainingSet)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	defer dset.Close()
	// read all the training examples
	var inputs [m][pixelsPerDimension][pixelsPerDimension][colorsChannel]uint8
	dset.Read(&inputs)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	// print only the first pixel of the first training example
	fmt.Println(inputs[0][0][0])
}
