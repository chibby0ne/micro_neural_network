// Package matrix provides the matrix type and NumberArray interface that are
// used for creating manipulating and operating in linear algebra
package matrix

import (
	"fmt"
	"math"
	"math/rand"
)

// Interface that condenses basic linear algebra operations operations
type NumberArray interface {
	GetColumns() int
	GetRows() int
	Transpose()
	SetValue(i, j int, v float64) error
	GetValue(i, j int) (float64, error)
}

// Basic type used for all linear algebra operations
type matrix struct {
	matrix [][]float64
	rows   int
	cols   int
}

// Checks that the sizes of the given matrix are non-zero and positive
func checkPositiveBounds(rows, cols int) (b bool, err error) {
	if rows < 1 {
		return b, fmt.Errorf("Can't create a matrix with %v rows", rows)
	}
	if cols < 1 {
		return b, fmt.Errorf("Can't create a matrix with %v cols", cols)
	}
	return true, err
}

// Creates a matrix of the given rows and colums
func NewMatrix(rows, cols int) (m *matrix, err error) {
	if ok, err := checkPositiveBounds(rows, cols); !ok {
		return m, err
	}
	m = new(matrix)
	m.matrix = make([][]float64, rows)
	for i := 0; i < rows; i++ {
		m.matrix[i] = make([]float64, cols)
	}
	m.rows = rows
	m.cols = cols
	return m, err
}

// Creates a matrix of the given rows and colums with all elements initialized
// with val. Useful when substracting 1 - Matrix in a formula, considering that
// 1 can be a matrix with all its elements set to 1.
func NewInitializedMatrix(rows, cols int, val float64) (m *matrix, err error) {
	if ok, err := checkPositiveBounds(rows, cols); !ok {
		return m, err
	}
	m = new(matrix)
	m.matrix = make([][]float64, rows)
	for i := 0; i < rows; i++ {
		m.matrix[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			m.matrix[i][j] = val
		}
	}
	m.rows = rows
	m.cols = cols
	return m, err
}

// Creates a random matrix of the given dimensions and whose random value is
// scaled by the given scaler.
// Hint: Scaler is used for initilizing the weights matrix as with a very small
// value, considering that only between very small values the sigmoid function
// is non-linear and at big enough values the function turns almost completely
// flat.
func NewRandomMatrix(rows, cols int, scaler float64) (m *matrix, err error) {
	if ok, err := checkPositiveBounds(rows, cols); !ok {
		return m, err
	}
	m = new(matrix)
	m.matrix = make([][]float64, rows)
	for i := 0; i < rows; i++ {
		m.matrix[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			m.matrix[i][j] = rand.Float64() * scaler
		}
	}
	m.rows = rows
	m.cols = cols
	return m, err
}

// Creates a new ColumnVector which is nothing more than a one-column matrix
func NewColumnVector(rows int) (*matrix, error) {
	return NewMatrix(rows, 1)
}

// Gets the number of rows of the matrix
func (m *matrix) GetRows() int {
	return m.rows
}

// Gets the number of columns of the matrix
func (m *matrix) GetColumns() int {
	return m.cols
}

// Gets the value in index i, j, or exits program if error
func (m *matrix) GetValue(i, j int) (f float64, err error) {
	if ok, err := m.checkBounds(i, j); !ok {
		return f, err
	}
	return m.matrix[i][j], err
}

// Sets a value v in the row i, and column j. i.e: m[i][j] = v
func (m *matrix) SetValue(i, j int, v float64) (err error) {
	if ok, err := m.checkBounds(i, j); !ok {
		return err
	}
	m.matrix[i][j] = v
	return err
}

// Checks that the indexes i and j are withing bounds
func (a *matrix) checkBounds(i, j int) (b bool, err error) {
	if i < 0 || i >= a.rows {
		return b, fmt.Errorf("Index %v is out of bounds for rows with size %v",
			i, a.rows)
	}
	if j < 0 || j >= a.cols {
		return b, fmt.Errorf("Index %v is out of bounds for cols with size %v",
			j, a.cols)
	}
	return true, err
}

// Helper function for Transposing a square matrix
func swap(val, val2 *float64) {
	temp := *val
	*val = *val2
	*val2 = temp
}

// Creates a new matrix that is the transpose of the original one and assigns
// the matrix pointer to the new matrix. In the special case that the matrix is
// square, it returns the same matrix object with the elements swapped
// corresponding to the transpose operation
func (a *matrix) Transpose() {
	var newMatrix *matrix
	if a.isSquareMatrix() {
		newMatrix = a
		for i := 0; i < a.rows; i++ {
			for j := i + 1; j < a.cols; j++ {
				swap(&(newMatrix.matrix[j][i]), &(a.matrix[i][j]))
			}
		}
	} else {
		newMatrix, _ = NewMatrix(a.cols, a.rows)
		for i := 0; i < a.rows; i++ {
			for j := 0; j < a.cols; j++ {
				newMatrix.matrix[j][i] = a.matrix[i][j]
			}
		}
	}
	*a = *newMatrix
}

// Checks that the matrix is square
func (a *matrix) isSquareMatrix() bool {
	return a.cols == a.rows
}

// True when the arrays a, b are of equal dimension, false otherwise
func EqualDimensions(a, b NumberArray) bool {
	return a.GetRows() == b.GetRows() && a.GetColumns() == b.GetColumns()
}

// Type used for handling binary math functions in binaryOperation function
type binaryMathFunc func(a, b float64) float64

// Peforms a + b and returns the result
func add(a, b float64) float64 {
	return a + b
}

// Peforms a - b and returns the result
func substract(a, b float64) float64 {
	return a - b
}

// Peforms a * b and returns the result
func multiply(a, b float64) float64 {
	return a * b
}

// Helper function that handles all the binary matrix operations, with the
// exception of Dot product operation
func binaryOperation(operation string, a, b NumberArray) (resultingMatrix NumberArray, err error) {
	var mathFunc binaryMathFunc
	switch operation {
	case "Add":
		mathFunc = add
	case "Substract":
		mathFunc = substract
	case "MultiplyElementwise":
		mathFunc = multiply
	default:
		return resultingMatrix, fmt.Errorf("Can't handle the given operation:"+
			" %v\n", operation)
	}
	if !EqualDimensions(a, b) {
		return resultingMatrix, fmt.Errorf("Can't perform " + operation + " on" +
			" matrices of different dimensions")
	}
	resultingMatrix, _ = NewMatrix(a.GetRows(), a.GetColumns())
	for i := 0; i < a.GetRows(); i++ {
		for j := 0; j < a.GetColumns(); j++ {
			operandA, _ := a.GetValue(i, j)
			operandB, _ := b.GetValue(i, j)
			resultingMatrix.SetValue(i, j, mathFunc(operandA, operandB))
		}
	}
	return resultingMatrix, err
}

// Performs an addition of a + b if the dimensions of the arrays are equal, and
// returns the reuslt in a new NumberArray
func Add(a, b NumberArray) (resultingMatrix NumberArray, err error) {
	return binaryOperation("Add", a, b)
}

// Performs a substraction of a - b if the dimensions of the arrays are equal,
// and returns the result in a new NumberArray
func Substract(a, b NumberArray) (resultingMatrix NumberArray, err error) {
	return binaryOperation("Substract", a, b)
}

// Performs a elementwise multiplication of a * b if the dimensions of the
// arrays are equal, and returns the result in a new NumberArray
func MultiplyElementwise(a, b NumberArray) (resultingMatrix NumberArray, err error) {
	return binaryOperation("MultiplyElementwise", a, b)
}

// Returns true if the columns' size of array a matches the rows' size of array b
func canBeMultiplied(a, b NumberArray) (ok bool, err error) {
	if a.GetColumns() != b.GetRows() {
		err = fmt.Errorf("Can't multiply matrices that don't satisfy "+
			"multiplication criteria, A.columns(): %v, B.rows(): %v",
			a.GetColumns(), b.GetRows())
		return ok, err
	}
	return true, err
}

// Peforms a dot product (a.k.a matrix multiplication) if the a and b can be
// multiplied, and returns the result in a new NumberArray
func Dot(a, b NumberArray) (resultingMatrix NumberArray, err error) {
	if ok, err := canBeMultiplied(a, b); !ok {
		return resultingMatrix, err
	}
	resultingMatrix, _ = NewMatrix(a.GetRows(), b.GetColumns())
	for i := 0; i < a.GetRows(); i++ {
		for j := 0; j < b.GetColumns(); j++ {
			sum := 0.0
			for k := 0; k < a.GetColumns(); k++ {
				operandA, _ := a.GetValue(i, k)
				operandB, _ := b.GetValue(k, j)
				sum += operandA * operandB
			}
			resultingMatrix.SetValue(i, j, sum)
		}
	}
	return resultingMatrix, err
}

// Type used for handling unary math functions in unaryOperation function
type unaryMathFunc func(float64) float64

// Performs a reLU function on val and returns the result
func reLU(val float64) float64 {
	return math.Max(0, val)
}

// Performs a reLU' function on val and returns the result
func derivativeReLU(val float64) float64 {
	if val >= 0 {
		return 1
	} else {
		return 0
	}
}

// Performs a tanh' function on val and returns the result
func derivativeTanh(val float64) float64 {
	return 1 - math.Pow(math.Tanh(val), 2)
}

// Performs a sigmoid function on val and returns the result
func sigmoid(val float64) float64 {
	return 1 / (1 + math.Exp(-1*val))
}

// Performs a sigmoid' function on val and returns the result
func derivativeSigmoid(val float64) float64 {
	return sigmoid(val) * (1 - sigmoid(val))
}

// helper function that handles all the unary matrix operations
func unaryOperation(operation string, a NumberArray) (resultingMatrix NumberArray, err error) {
	var mathFunc unaryMathFunc
	switch operation {
	case "Exp":
		mathFunc = math.Exp
	case "Log":
		mathFunc = math.Log
	case "Sigmoid":
		mathFunc = sigmoid
	case "DerivativeSigmoid":
		mathFunc = derivativeSigmoid
	case "Tanh":
		mathFunc = math.Tanh
	case "DerivativeTanh":
		mathFunc = derivativeTanh
	case "ReLU":
		mathFunc = reLU
	case "DerivativeReLU":
		mathFunc = derivativeReLU
	default:
		return resultingMatrix, fmt.Errorf("Can't handle the given operation:"+
			" %v\n", operation)
	}
	resultingMatrix, _ = NewMatrix(a.GetRows(), a.GetColumns())
	for i := 0; i < a.GetRows(); i++ {
		for j := 0; j < a.GetColumns(); j++ {
			operandA, _ := a.GetValue(i, j)
			resultingMatrix.SetValue(i, j, mathFunc(operandA))
		}
	}
	return resultingMatrix, err
}

// Performs a math.Exp operation on the entire NumberArray and returns the
// result in a new NumberArray
func Exp(a NumberArray) (resultingMatrix NumberArray, err error) {
	return unaryOperation("Exp", a)
}

// Performs a math.Log operation on the entire NumberArray and returns the
// result in a new NumberArray
func Log(a NumberArray) (resultingMatrix NumberArray, err error) {
	return unaryOperation("Log", a)
}

// Performs a Sigmoid function i.e: 1 / (1 + e^(-z)) on the entire NumberArray
// and returns the result in a new NumberArray
func Sigmoid(a NumberArray) (resultingMatrix NumberArray, err error) {
	return unaryOperation("Sigmoid", a)
}

// Performs the derivative of the Sigmoid function i.e: Sigmoid(x) * (1 -
// Sigmoid(x)) on the entire NumberArray and returns the result in a new
// NumberArray
func DerivativeSigmoid(a NumberArray) (resultingMatrix NumberArray, err error) {
	return unaryOperation("DerivativeSigmoid", a)
}

// Performs Tanh function i.e: (e^z - e^(-z)) / (e^z + e^(-z)) on the entire
// NumberArray and returns the result in a new NumberArray
func Tanh(a NumberArray) (resultingMatrix NumberArray, err error) {
	return unaryOperation("Tanh", a)
}

// Performs the derivative of the Tanh function i.e: 1 - (Tanh(x))^2 on the
// entire NumberArray and returns the result in a new NumberArray
func DerivativeTanh(a NumberArray) (resultingMatrix NumberArray, err error) {
	return unaryOperation("DerivativeTanh", a)
}

// Performs a ReLU function i.e: max(0, x) on the NumberArray and returns the
// result in a new NumberArray
func ReLU(a NumberArray) (resultingMatrix NumberArray, err error) {
	return unaryOperation("ReLU", a)
}

// Performs the derivative of the ReLU function on the NumberArray and returns
// the result in a new NumberArray
func DerivativeReLU(a NumberArray) (resultingMatrix NumberArray, err error) {
	return unaryOperation("DerivativeReLU", a)
}

// Multiply Matrix by scalar. Useful for operations where we divide the matrix
// elements among the number of training examples m.
func MultiplyScalar(a NumberArray, scalar float64) NumberArray {
	resultingMatrix, _ := NewMatrix(a.GetRows(), a.GetColumns())
	for i := 0; i < a.GetRows(); i++ {
		for j := 0; j < a.GetColumns(); j++ {
			operandA, _ := a.GetValue(i, j)
			resultingMatrix.SetValue(i, j, operandA*scalar)
		}
	}
	return resultingMatrix
}

// Sum among columns, i.e: returns a column vector where each row element is
// the sum of the elements of all the columns for that row.
// Emulates np.sum(X, axis=1, keepdims=True) in Python.
func SumByColumns(a NumberArray) NumberArray {
	resultingMatrix, _ := NewMatrix(a.GetRows(), 1)
	var sum float64 = 0
	for i := 0; i < a.GetRows(); i++ {
		sum = 0
		for j := 0; j < a.GetColumns(); j++ {
			operandA, _ := a.GetValue(i, j)
			sum += operandA
		}
		resultingMatrix.SetValue(i, 0, sum)
	}
	return resultingMatrix
}
