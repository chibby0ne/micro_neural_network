// Package matrix provides the matrix type and NumberArray interface that are used for creating manipulating and operating in linear algebra
package matrix

import (
	"fmt"
	"os"
)

// Interface that condenses basic linear algebra operations operations
type NumberArray interface {
	GetColumns() uint64
	GetRows() uint64
	Transpose()
	SetValue(i, j uint64, v float64) error
	GetValue(i, j uint64) (float64, error)
}

// Basic type used for all linear algebra operations
type matrix struct {
	matrix [][]float64
	rows   uint64
	cols   uint64
}

func checkPositiveBounds(rows, cols uint64) (b bool, err error) {
	if rows < 1 {
		return b, fmt.Errorf("Can't create a matrix with %v rows", rows)
	}
	if cols < 1 {
		return b, fmt.Errorf("Can't create a matrix with %v cols", cols)
	}
	return true, err
}

// Creates a matrix of the given rows and colums
func NewMatrix(rows, cols uint64) (m *matrix, err error) {
	if ok, err := checkPositiveBounds(rows, cols); !ok {
		return m, err
	}
	m = new(matrix)
	m.matrix = make([][]float64, rows)
	for i := uint64(0); i < rows; i++ {
		m.matrix[i] = make([]float64, cols)
	}
	m.rows = rows
	m.cols = cols
	return m, err
}

// Creates a new ColumnVector which is nothing more than a one-column matrix
func NewColumnVector(rows uint64) (*matrix, error) {
	return NewMatrix(rows, 1)
}

// // Gets the dimensions of the matrix as rows, columns pair
// func (m *matrix) GetDimensions() (uint64, uint64) {
// 	return m.rows, m.cols
// }

// Gets the number of rows of the matrix
func (m *matrix) GetRows() uint64 {
	return m.rows
}

// Gets the number of columns of the matrix
func (m *matrix) GetColumns() uint64 {
	return m.cols
}

// Gets the value in index i, j, or exits program if error
func (m *matrix) GetValue(i, j uint64) (f float64, err error) {
	if ok, err := m.checkBounds(i, j); !ok {
		return f, err
	}
	return m.matrix[i][j], err
}

// Sets a value v in the row i, and column j. i.e: m[i][j] = v
func (m *matrix) SetValue(i, j uint64, v float64) (err error) {
	if ok, err := m.checkBounds(i, j); !ok {
		return err
	}
	m.matrix[i][j] = v
	return err
}

// Checks that the indexes i and j are withing bounds
func (a *matrix) checkBounds(i, j uint64) (b bool, err error) {
	if i < 0 || i >= a.rows {
		return b, fmt.Errorf("Index %v is out of bounds for rows with size %v", i, a.rows)

	}
	if j < 0 || j >= a.cols {
		return b, fmt.Errorf("Index %v is out of bounds for cols with size %v", j, a.rows)
	}
	return true, err
}

// Creates a new matrix that is the transpose of the original one and assigns
// the matrix pointer to the new matrix.
func (a *matrix) Transpose() {
	// var newM *matrix = &((*a).matrix)
	// if !a.isSquareMatrix {
	// 	newM = newMatrix(a.cols, a.rows)
	// }
}

// Checks that the matrix is square
func (a *matrix) isSquareMatrix() bool {
	return a.cols == a.rows
}

// True when the arrays a, b are of equal dimension, false otherwise
func EqualDimensions(a, b NumberArray) bool {
	return a.GetRows() == b.GetRows() && a.GetColumns() == b.GetColumns()
}

// Performs an addition of a + b if the dimensions of the arrays are equal, and returns the reuslt in a new NumberArray
func Add(a, b NumberArray) NumberArray {
	if !EqualDimensions(a, b) {
		fmt.Errorf("matrices of different dimensions can't be added")
		os.Exit(1)
	}
	fmt.Println("Everything ok")
	return nil
}

// Performs a substraction of a - b if the dimensions of the arrays are equal, and returns the result in a new NumberArray
func Substract(a, b NumberArray) NumberArray {
	if !EqualDimensions(a, b) {
		fmt.Errorf("matrices of different dimensions can't be substracted")
		os.Exit(1)
	}
	fmt.Println("Everything ok")
	return nil
}

// Performs a elementwise multiplication of a * b if the dimensions of the arrays are equal, and returns the result in a new NumberArray
func MultiplyElementwise(a, b NumberArray) NumberArray {
	if !EqualDimensions(a, b) {
		fmt.Errorf("matrices of different dimensions can't be multiplied elementwise")
		os.Exit(1)
	}
	fmt.Println("Everything ok")
	return nil
}

// Returns true if the columns' size of array a matches the rows' size of array b
func canBeMultiplied(a, b NumberArray) (ok bool, err error) {
	if a.GetColumns() != b.GetRows() {
		err = fmt.Errorf(`Can't multiply matrices that don't satisfy multiplication 
		"criteria, A.columns(): %v, B.rows(): %v`, a.GetColumns(), b.GetRows())
		return ok, err
	}
	return true, err
}

// Peforms a dot product (a.k.a matrix multiplication) if the a and b can be multiplied, and returns the result in a new NumberArray
func Dot(a, b NumberArray) (resultingMatrix NumberArray, err error) {
	if ok, err := canBeMultiplied(a, b); !ok {
		return resultingMatrix, err
	}
	resultingMatrix, err = NewMatrix(a.GetRows(), b.GetColumns())
	if err != nil {
		return resultingMatrix, err
	}
	for i := uint64(0); i < a.GetRows(); i++ {
		for j := uint64(0); j < b.GetColumns(); j++ {
			sum := 0.0
			for k := uint64(0); k < a.GetColumns(); k++ {
				operandA, _ := a.GetValue(i, k)
				operandB, _ := b.GetValue(k, j)
				sum += operandA * operandB
			}
			resultingMatrix.SetValue(i, j, sum)
		}
	}
	return resultingMatrix, err
}

//
// func matrixMultiplication(a, b NumberArray) (resultingMatrix NumberArray, err error) {
// 	if a.GetColumns() != b.GetRows() {
// 		fmt.Errorf("Can't multiply matrices that don't satisfy multiplication criteria")
// 		err = errors.New("Can't multiply matrices that don't satisfy multiplication criteria")
// 		return nil, err
// 	}
// 	resultingMatrix = newMatrix(a.GetRows(), b.GetColumns())
// 	for i := uint64(0); i < a.GetRows(); i++ {
// 		for j := uint64(0); j < b.GetColumns(); j++ {
// 			sum := 0.0
// 			for k := uint64(0); k < a.GetColumns(); k++ {
// 				sum += a.GetValue(i, k) * b.GetValue(k, j)
// 			}
// 			resultingMatrix.SetValue(i, j, sum)
// 		}
// 	}
// 	return resultingMatrix, nil
// }
