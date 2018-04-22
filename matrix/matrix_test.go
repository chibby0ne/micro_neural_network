package matrix

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func equalMatrices(a, b *matrix) bool {
	// check that they are null and equal
	if a == nil && b == nil {
		return true
	}
	// both of them must be NOT null at this point otherwise they are not equal
	if a == nil || b == nil {
		return false
	}
	if a.rows != b.rows || a.cols != b.cols {
		return false
	}
	return equalSlices(a.matrix, b.matrix)
}

func equalSlices(a, b [][]float64) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a[0]) != len(b[0]) {
		return false
	}
	// fmt.Printf("a: %v, b: %v\n", a, b)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			if a[i][j] != b[i][j] && (!math.IsNaN(a[i][j]) && !math.IsNaN(b[i][j])) {
				return false
			}
		}
	}
	return true
}

func equalErrors(err1, err2 error) bool {
	if err1 == nil && err2 == nil {
		return true
	}
	if err1 == nil || err2 == nil {
		return false
	}
	return err1.Error() == err2.Error()
}

func TestNewMatrix(t *testing.T) {
	tables := []struct {
		rows           int
		cols           int
		expectedMatrix *matrix
		expectedError  error
	}{
		{-1, -1, nil, fmt.Errorf("Can't create a matrix with -1 rows")},
		{-1, 0, nil, fmt.Errorf("Can't create a matrix with -1 rows")},
		{1, -1, nil, fmt.Errorf("Can't create a matrix with -1 cols")},
		{0, 0, nil, fmt.Errorf("Can't create a matrix with 0 rows")},
		{0, 1, nil, fmt.Errorf("Can't create a matrix with 0 rows")},
		{1, 0, nil, fmt.Errorf("Can't create a matrix with 0 cols")},
		{1, 1, &matrix{matrix: [][]float64{{0}}, rows: 1, cols: 1}, nil},
		{2, 2, &matrix{matrix: [][]float64{{0, 0}, {0, 0}}, rows: 2, cols: 2}, nil},
		{3, 3, &matrix{matrix: [][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, rows: 3, cols: 3}, nil},
		{3, 2, &matrix{matrix: [][]float64{{0, 0}, {0, 0}, {0, 0}}, rows: 3, cols: 2}, nil},
		{5, 1, &matrix{matrix: [][]float64{{0}, {0}, {0}, {0}, {0}}, rows: 5, cols: 1}, nil},
		{1, 5, &matrix{matrix: [][]float64{{0, 0, 0, 0, 0}}, rows: 1, cols: 5}, nil},
	}
	for _, table := range tables {
		m, err := NewMatrix(table.rows, table.cols)
		if !equalMatrices(table.expectedMatrix, m) {
			t.Errorf("ExpectedMatrix: %v, ActualMatrix: %v\n", table.expectedMatrix, m)
		}
		if !equalErrors(table.expectedError, err) {
			t.Errorf("ExpectedError: %v, ActualError: %v", table.expectedError, err)
		}
	}
}

func TestNewInitializedMatrix(t *testing.T) {
	tables := []struct {
		rows           int
		cols           int
		val            float64
		expectedMatrix *matrix
		expectedError  error
	}{
		{-1, -1, 10, nil, fmt.Errorf("Can't create a matrix with -1 rows")},
		{-1, 0, 0, nil, fmt.Errorf("Can't create a matrix with -1 rows")},
		{1, -1, 15, nil, fmt.Errorf("Can't create a matrix with -1 cols")},
		{0, 0, 10, nil, fmt.Errorf("Can't create a matrix with 0 rows")},
		{0, 1, 5, nil, fmt.Errorf("Can't create a matrix with 0 rows")},
		{1, 0, 10, nil, fmt.Errorf("Can't create a matrix with 0 cols")},
		{1, 1, 12, &matrix{matrix: [][]float64{{12}}, rows: 1, cols: 1}, nil},
		{2, 2, 10, &matrix{matrix: [][]float64{{10, 10}, {10, 10}}, rows: 2, cols: 2}, nil},
		{3, 3, 5, &matrix{matrix: [][]float64{{5, 5, 5}, {5, 5, 5}, {5, 5, 5}}, rows: 3, cols: 3}, nil},
		{3, 2, 4.3, &matrix{matrix: [][]float64{{4.3, 4.3}, {4.3, 4.3}, {4.3, 4.3}}, rows: 3, cols: 2}, nil},
		{5, 1, 0.12, &matrix{matrix: [][]float64{{0.12}, {0.12}, {0.12}, {0.12}, {0.12}}, rows: 5, cols: 1}, nil},
		{1, 5, -2.1, &matrix{matrix: [][]float64{{-2.1, -2.1, -2.1, -2.1, -2.1}}, rows: 1, cols: 5}, nil},
	}
	for _, table := range tables {
		m, err := NewInitializedMatrix(table.rows, table.cols, table.val)
		if !equalMatrices(table.expectedMatrix, m) {
			t.Errorf("ExpectedMatrix: %v, ActualMatrix: %v\n", table.expectedMatrix, m)
		}
		if !equalErrors(table.expectedError, err) {
			t.Errorf("ExpectedError: %v, ActualError: %v", table.expectedError, err)
		}
	}
}

func TestNewRandomMatrix(t *testing.T) {
	tables := []struct {
		rows           int
		cols           int
		expectedMatrix *matrix
		expectedError  error
	}{
		{-1, -1, nil, fmt.Errorf("Can't create a matrix with -1 rows")},
		{-1, 0, nil, fmt.Errorf("Can't create a matrix with -1 rows")},
		{1, -1, nil, fmt.Errorf("Can't create a matrix with -1 cols")},
		{0, 0, nil, fmt.Errorf("Can't create a matrix with 0 rows")},
		{0, 1, nil, fmt.Errorf("Can't create a matrix with 0 rows")},
		{1, 0, nil, fmt.Errorf("Can't create a matrix with 0 cols")},
		// These values are calculated using Seed(1) and in this order
		{1, 1, &matrix{matrix: [][]float64{{0.6046602879796196}}, rows: 1, cols: 1}, nil},
		{2, 2, &matrix{matrix: [][]float64{{0.9405090880450124, 0.6645600532184904}, {0.4377141871869802, 0.4246374970712657}}, rows: 2, cols: 2}, nil},
		{3, 3, &matrix{matrix: [][]float64{{0.6868230728671094, 0.06563701921747622, 0.15651925473279124}, {0.09696951891448456, 0.30091186058528707, 0.5152126285020654}, {0.8136399609900968, 0.21426387258237492, 0.380657189299686}}, rows: 3, cols: 3}, nil},
		{3, 2, &matrix{matrix: [][]float64{{0.31805817433032985, 0.4688898449024232}, {0.28303415118044517, 0.29310185733681576}, {0.6790846759202163, 0.21855305259276428}}, rows: 3, cols: 2}, nil},
		{5, 1, &matrix{matrix: [][]float64{{0.20318687664732285}, {0.360871416856906}, {0.5706732760710226}, {0.8624914374478864}, {0.29311424455385804}}, rows: 5, cols: 1}, nil},
		{1, 5, &matrix{matrix: [][]float64{{0.29708256355629153, 0.7525730355516119, 0.2065826619136986, 0.865335013001561, 0.6967191657466347}}, rows: 1, cols: 5}, nil},
	}
	// To get the same values between runs
	rand.Seed(1)
	for _, table := range tables {
		m, err := NewRandomMatrix(table.rows, table.cols, 1)
		if !equalMatrices(table.expectedMatrix, m) {
			t.Errorf("ExpectedMatrix: %v, ActualMatrix: %v\n", table.expectedMatrix, m)
		}
		if !equalErrors(table.expectedError, err) {
			t.Errorf("ExpectedError: %v, ActualError: %v", table.expectedError, err)
		}
	}
}

func TestNewColumnVector(t *testing.T) {
	tables := []struct {
		rows           int
		expectedMatrix *matrix
		expectedError  error
	}{
		{-1, nil, fmt.Errorf("Can't create a matrix with -1 rows")},
		{0, nil, fmt.Errorf("Can't create a matrix with 0 rows")},
		{1, &matrix{matrix: [][]float64{{0}}, rows: 1, cols: 1}, nil},
		{2, &matrix{matrix: [][]float64{{0}, {0}}, rows: 2, cols: 1}, nil},
		{3, &matrix{matrix: [][]float64{{0}, {0}, {0}}, rows: 3, cols: 1}, nil},
		{4, &matrix{matrix: [][]float64{{0}, {0}, {0}, {0}}, rows: 4, cols: 1}, nil},
		{5, &matrix{matrix: [][]float64{{0}, {0}, {0}, {0}, {0}}, rows: 5, cols: 1}, nil},
	}
	for _, table := range tables {
		m, err := NewColumnVector(table.rows)
		if !equalMatrices(table.expectedMatrix, m) {
			t.Errorf("ExpectedMatrix: %v, ActualMatrix: %v\n", table.expectedMatrix, m)
		}
		if !equalErrors(table.expectedError, err) {
			t.Errorf("ExpectedError: %v, ActualError: %v", table.expectedError, err)
		}
	}
}

func TestGetRows(t *testing.T) {
	tables := []struct {
		rows         int
		cols         int
		expectedRows int
	}{
		{1, 1, 1},
		{10, 1, 10},
		{1, 10, 1},
	}
	for _, table := range tables {
		m, _ := NewMatrix(table.rows, table.cols)
		if table.expectedRows != m.rows {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedRows, m.rows)
		}
	}
}

func TestGetColumns(t *testing.T) {
	tables := []struct {
		rows         int
		cols         int
		expectedCols int
	}{
		{1, 1, 1},
		{10, 1, 1},
		{1, 10, 10},
	}
	for _, table := range tables {
		m, _ := NewMatrix(table.rows, table.cols)
		if table.expectedCols != m.cols {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedCols, m.cols)
		}
	}
}

func TestGetValue(t *testing.T) {
	m := matrix{matrix: [][]float64{{1, 2}, {3, 4}, {5, 6}}, rows: 3, cols: 2}
	tables := []struct {
		i             int
		j             int
		expectedValue float64
		expectedError error
	}{
		{-1, 2, 0.0, fmt.Errorf("Index -1 is out of bounds for rows with size 3")},
		{1, -1, 0.0, fmt.Errorf("Index -1 is out of bounds for cols with size 2")},
		{0, 0, 1, nil},
		{2, 1, 6, nil},
		{1, 1, 4, nil},
		{4, 3, 0.0, fmt.Errorf("Index 4 is out of bounds for rows with size 3")},
		{2, 3, 0.0, fmt.Errorf("Index 3 is out of bounds for cols with size 2")},
	}
	for _, table := range tables {
		val, err := m.GetValue(table.i, table.j)
		if table.expectedValue != val {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedValue, val)
		}
		if !equalErrors(table.expectedError, err) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedError, err)
		}
	}
}

func TestSetValue(t *testing.T) {
	m := matrix{matrix: [][]float64{{1, 2}, {3, 4}, {5, 6}}, rows: 3, cols: 2}
	tables := []struct {
		i             int
		j             int
		value         float64
		expectedValue float64 // only used for tests with valid indexes
		expectedError error
	}{
		{-1, 2, 0.0, 0.0, fmt.Errorf("Index -1 is out of bounds for rows with size 3")},
		{1, -1, 10, -10, fmt.Errorf("Index -1 is out of bounds for cols with size 2")},
		{0, 0, 5, 5, nil},
		{2, 1, 10, 10, nil},
		{1, 1, 4.4, 4.4, nil},
		{4, 3, 32.2, 32.2, fmt.Errorf("Index 4 is out of bounds for rows with size 3")},
		{2, 3, 32.2, 32.2, fmt.Errorf("Index 3 is out of bounds for cols with size 2")},
	}
	for _, table := range tables {
		err := m.SetValue(table.i, table.j, table.value)
		// Only check for set value if the indexes given are in between bounds
		if ok, _ := m.checkBounds(table.i, table.j); ok {
			val, _ := m.GetValue(table.i, table.j)
			if table.expectedValue != val {
				t.Errorf("Expected: %v, Actual: %v\n", table.expectedValue, val)
			}
		}
		if !equalErrors(table.expectedError, err) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedError, err)
		}
	}
}

type MockNumberArray struct {
	rows, cols int
}

func newNockNumberArray(rows, cols int) *MockNumberArray {
	mock := new(MockNumberArray)
	mock.rows = rows
	mock.cols = cols
	return mock
}

func (numArray *MockNumberArray) GetColumns() int {
	return numArray.cols
}

func (numArray *MockNumberArray) GetRows() int {
	return numArray.rows
}

func (numArray *MockNumberArray) GetValue(i, j int) (float64, error) {
	return 0, nil
}

func (numArray *MockNumberArray) SetValue(i, j int, val float64) error {
	return nil
}

func (numArray *MockNumberArray) Transpose() {

}

func TestEqualDimensions(t *testing.T) {
	tables := []struct {
		a              NumberArray
		b              NumberArray
		expectedResult bool
	}{
		{
			&MockNumberArray{rows: 3, cols: 10},
			&MockNumberArray{rows: 3, cols: 10},
			true,
		},
		{
			&MockNumberArray{rows: 3, cols: 9},
			&MockNumberArray{rows: 3, cols: 10},
			false,
		},
		{
			&MockNumberArray{rows: 2, cols: 10},
			&MockNumberArray{rows: 3, cols: 10},
			false,
		},
		{
			&MockNumberArray{rows: 2, cols: 10},
			&MockNumberArray{rows: 1, cols: 3},
			false,
		},
	}
	for _, table := range tables {
		actual := EqualDimensions(table.a, table.b)
		if table.expectedResult != actual {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedResult, actual)
		}
	}
}

func TestAdd(t *testing.T) {
	tables := []struct {
		a              *matrix
		b              *matrix
		expectedMatrix *matrix
		expectedError  error
	}{
		{
			&matrix{matrix: [][]float64{{1, 2}, {3, 4}, {5, 6}}, rows: 3, cols: 2},
			&matrix{matrix: [][]float64{{1, 2}, {3, 4}, {5, 6}}, rows: 3, cols: 2},
			&matrix{matrix: [][]float64{{2, 4}, {6, 8}, {10, 12}}, rows: 3, cols: 2},
			nil,
		},
		{
			&matrix{matrix: [][]float64{{1, 2, 3}, {3, 4, 5}, {5, 6, 7}}, rows: 3, cols: 3},
			&matrix{matrix: [][]float64{{1, 2}, {3, 4}, {5, 6}}, rows: 3, cols: 2},
			nil,
			fmt.Errorf("Can't perform Add on matrices of different dimensions"),
		},
	}
	for _, table := range tables {
		resultMatrix, err := Add(table.a, table.b)
		v, _ := resultMatrix.(*matrix)
		if !equalMatrices(table.expectedMatrix, v) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedMatrix, v)
		}
		if !equalErrors(table.expectedError, err) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedError, err)
		}
	}
}

func TestSubstract(t *testing.T) {
	tables := []struct {
		a              *matrix
		b              *matrix
		expectedMatrix *matrix
		expectedError  error
	}{
		{
			&matrix{matrix: [][]float64{{2, 4}, {6, 8}, {10, 12}}, rows: 3, cols: 2},
			&matrix{matrix: [][]float64{{1, 2}, {3, 4}, {5, 6}}, rows: 3, cols: 2},
			&matrix{matrix: [][]float64{{1, 2}, {3, 4}, {5, 6}}, rows: 3, cols: 2},
			nil,
		},
		{
			&matrix{matrix: [][]float64{{1, 2, 3}, {3, 4, 5}, {5, 6, 7}}, rows: 3, cols: 3},
			&matrix{matrix: [][]float64{{1, 2}, {3, 4}, {5, 6}}, rows: 3, cols: 2},
			nil,
			fmt.Errorf("Can't perform Substract on matrices of different dimensions"),
		},
	}
	for _, table := range tables {
		resultMatrix, err := Substract(table.a, table.b)
		v, _ := resultMatrix.(*matrix)
		if !equalMatrices(table.expectedMatrix, v) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedMatrix, v)
		}
		if !equalErrors(table.expectedError, err) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedError, err)
		}
	}
}

func TestMultiplyElementwise(t *testing.T) {
	tables := []struct {
		a              *matrix
		b              *matrix
		expectedMatrix *matrix
		expectedError  error
	}{
		{
			&matrix{matrix: [][]float64{{1, 2}, {3, 4}, {5, 6}}, rows: 3, cols: 2},
			&matrix{matrix: [][]float64{{1, 2}, {3, 4}, {5, 6}}, rows: 3, cols: 2},
			&matrix{matrix: [][]float64{{1, 4}, {9, 16}, {25, 36}}, rows: 3, cols: 2},
			nil,
		},
		{
			&matrix{matrix: [][]float64{{1, 2, 3}, {3, 4, 5}, {5, 6, 7}}, rows: 3, cols: 3},
			&matrix{matrix: [][]float64{{1, 2}, {3, 4}, {5, 6}}, rows: 3, cols: 2},
			nil,
			fmt.Errorf("Can't perform MultiplyElementwise on matrices of different dimensions"),
		},
	}
	for _, table := range tables {
		resultMatrix, err := MultiplyElementwise(table.a, table.b)
		v, _ := resultMatrix.(*matrix)
		if !equalMatrices(table.expectedMatrix, v) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedMatrix, v)
		}
		if !equalErrors(table.expectedError, err) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedError, err)
		}
	}
}

func TestDot(t *testing.T) {
	tables := []struct {
		a              *matrix
		b              *matrix
		expectedMatrix *matrix
		expectedError  error
	}{
		{
			&matrix{matrix: [][]float64{{1, 2, 3}, {3, 4, 5}, {5, 6, 7}}, rows: 3, cols: 3},
			&matrix{matrix: [][]float64{{1, 2}, {3, 4}, {5, 6}}, rows: 3, cols: 2},
			&matrix{matrix: [][]float64{{22, 28}, {40, 52}, {58, 76}}, rows: 3, cols: 2},
			nil,
		},
		{
			&matrix{matrix: [][]float64{{1, 2}, {3, 4}, {5, 6}}, rows: 3, cols: 2},
			&matrix{matrix: [][]float64{{1, 2}, {3, 4}, {5, 6}}, rows: 3, cols: 2},
			nil,
			fmt.Errorf("Can't multiply matrices that don't satisfy multiplication criteria, A.columns(): 2, B.rows(): 3"),
		},
	}
	for _, table := range tables {
		resultMatrix, err := Dot(table.a, table.b)
		v, _ := resultMatrix.(*matrix)
		if !equalMatrices(table.expectedMatrix, v) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedMatrix, v)
		}
		if !equalErrors(table.expectedError, err) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedError, err)
		}
	}
}

func TestExp(t *testing.T) {
	tables := []struct {
		a              *matrix
		expectedMatrix *matrix
	}{
		{
			&matrix{matrix: [][]float64{{1, 2, 3, 4}}, rows: 1, cols: 4},
			&matrix{matrix: [][]float64{{math.Exp(1), math.Exp(2), math.Exp(3), math.Exp(4)}}, rows: 1, cols: 4},
		},
		{
			&matrix{matrix: [][]float64{{-2.3, 3.3}, {1.2, -4.0}}, rows: 2, cols: 2},
			&matrix{matrix: [][]float64{{math.Exp(-2.3), math.Exp(3.3)}, {math.Exp(1.2), math.Exp(-4.0)}}, rows: 2, cols: 2},
		},
	}
	for _, table := range tables {
		actual, _ := Exp(table.a)
		v, _ := actual.(*matrix)
		if !equalMatrices(table.expectedMatrix, v) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedMatrix, v)
		}
	}
}

func TestLog(t *testing.T) {
	tables := []struct {
		a              *matrix
		expectedMatrix *matrix
	}{
		{
			&matrix{matrix: [][]float64{{1, 2, 3, 4}}, rows: 1, cols: 4},
			&matrix{matrix: [][]float64{{math.Log(1), math.Log(2), math.Log(3), math.Log(4)}}, rows: 1, cols: 4},
		},
		{
			&matrix{matrix: [][]float64{{-2.3, 3.3}, {1.2, -4.0}}, rows: 2, cols: 2},
			&matrix{matrix: [][]float64{{math.Log(-2.3), math.Log(3.3)}, {math.Log(1.2), math.Log(-4.0)}}, rows: 2, cols: 2},
		},
	}
	for _, table := range tables {
		actual, _ := Log(table.a)
		v, _ := actual.(*matrix)
		if !equalMatrices(table.expectedMatrix, v) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedMatrix, v)
		}
	}
}

func TestTranspose(t *testing.T) {
	tables := []struct {
		a              *matrix
		expectedMatrix *matrix
	}{
		{
			&matrix{matrix: [][]float64{{1, 2, 3, 4}}, rows: 1, cols: 4},
			&matrix{matrix: [][]float64{{1}, {2}, {3}, {4}}, rows: 4, cols: 1},
		},
		{
			&matrix{matrix: [][]float64{{-2.3, 3.3}, {1.2, -4.0}}, rows: 2, cols: 2},
			&matrix{matrix: [][]float64{{-2.3, 1.2}, {3.3, -4.0}}, rows: 2, cols: 2},
		},
	}
	for _, table := range tables {
		table.a.Transpose()
		if !equalMatrices(table.expectedMatrix, table.a) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedMatrix, table.a)
		}
	}
}

func TestSigmoid(t *testing.T) {
	tables := []struct {
		a              *matrix
		expectedMatrix *matrix
	}{
		{
			&matrix{matrix: [][]float64{{1, 2, 3, 4}}, rows: 1, cols: 4},
			&matrix{matrix: [][]float64{{sigmoid(1), sigmoid(2), sigmoid(3), sigmoid(4)}}, rows: 1, cols: 4},
		},
		{
			&matrix{matrix: [][]float64{{-2.3, 3.3}, {1.2, -4.0}}, rows: 2, cols: 2},
			&matrix{matrix: [][]float64{{sigmoid(-2.3), sigmoid(3.3)}, {sigmoid(1.2), sigmoid(-4.0)}}, rows: 2, cols: 2},
		},
	}
	for _, table := range tables {
		actual, _ := Sigmoid(table.a)
		v, _ := actual.(*matrix)
		if !equalMatrices(table.expectedMatrix, v) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedMatrix, v)
		}
	}
}

func TestDerivativeSigmoid(t *testing.T) {
	tables := []struct {
		a              *matrix
		expectedMatrix *matrix
	}{
		{
			&matrix{matrix: [][]float64{{1, 2, 3, 4}}, rows: 1, cols: 4},
			&matrix{matrix: [][]float64{{derivativeSigmoid(1), derivativeSigmoid(2), derivativeSigmoid(3), derivativeSigmoid(4)}}, rows: 1, cols: 4},
		},
		{
			&matrix{matrix: [][]float64{{-2.3, 3.3}, {1.2, -4.0}}, rows: 2, cols: 2},
			&matrix{matrix: [][]float64{{derivativeSigmoid(-2.3), derivativeSigmoid(3.3)}, {derivativeSigmoid(1.2), derivativeSigmoid(-4.0)}}, rows: 2, cols: 2},
		},
	}
	for _, table := range tables {
		actual, _ := DerivativeSigmoid(table.a)
		v, _ := actual.(*matrix)
		if !equalMatrices(table.expectedMatrix, v) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedMatrix, v)
		}
	}
}

func TestTanh(t *testing.T) {
	tables := []struct {
		a              *matrix
		expectedMatrix *matrix
	}{
		{
			&matrix{matrix: [][]float64{{1, 2, 3, 4}}, rows: 1, cols: 4},
			&matrix{matrix: [][]float64{{math.Tanh(1), math.Tanh(2), math.Tanh(3), math.Tanh(4)}}, rows: 1, cols: 4},
		},
		{
			&matrix{matrix: [][]float64{{-2.3, 3.3}, {1.2, -4.0}}, rows: 2, cols: 2},
			&matrix{matrix: [][]float64{{math.Tanh(-2.3), math.Tanh(3.3)}, {math.Tanh(1.2), math.Tanh(-4.0)}}, rows: 2, cols: 2},
		},
	}
	for _, table := range tables {
		actual, _ := Tanh(table.a)
		v, _ := actual.(*matrix)
		if !equalMatrices(table.expectedMatrix, v) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedMatrix, v)
		}
	}
}

func TestDerivativeTanh(t *testing.T) {
	tables := []struct {
		a              *matrix
		expectedMatrix *matrix
	}{
		{
			&matrix{matrix: [][]float64{{1, 2, 3, 4}}, rows: 1, cols: 4},
			&matrix{matrix: [][]float64{{derivativeTanh(1), derivativeTanh(2), derivativeTanh(3), derivativeTanh(4)}}, rows: 1, cols: 4},
		},
		{
			&matrix{matrix: [][]float64{{-2.3, 3.3}, {1.2, -4.0}}, rows: 2, cols: 2},
			&matrix{matrix: [][]float64{{derivativeTanh(-2.3), derivativeTanh(3.3)}, {derivativeTanh(1.2), derivativeTanh(-4.0)}}, rows: 2, cols: 2},
		},
	}
	for _, table := range tables {
		actual, _ := DerivativeTanh(table.a)
		v, _ := actual.(*matrix)
		if !equalMatrices(table.expectedMatrix, v) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedMatrix, v)
		}
	}
}

func TestReLU(t *testing.T) {
	tables := []struct {
		a              *matrix
		expectedMatrix *matrix
	}{
		{
			&matrix{matrix: [][]float64{{1, 2, 3, 4}}, rows: 1, cols: 4},
			&matrix{matrix: [][]float64{{reLU(1), reLU(2), reLU(3), reLU(4)}}, rows: 1, cols: 4},
		},
		{
			&matrix{matrix: [][]float64{{-2.3, 3.3}, {1.2, -4.0}}, rows: 2, cols: 2},
			&matrix{matrix: [][]float64{{reLU(-2.3), reLU(3.3)}, {reLU(1.2), reLU(-4.0)}}, rows: 2, cols: 2},
		},
	}
	for _, table := range tables {
		actual, _ := ReLU(table.a)
		v, _ := actual.(*matrix)
		if !equalMatrices(table.expectedMatrix, v) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedMatrix, v)
		}
	}
}

func TestDerivativeReLU(t *testing.T) {
	tables := []struct {
		a              *matrix
		expectedMatrix *matrix
	}{
		{
			&matrix{matrix: [][]float64{{1, 2, 3, 4}}, rows: 1, cols: 4},
			&matrix{matrix: [][]float64{{derivativeReLU(1), derivativeReLU(2), derivativeReLU(3), derivativeReLU(4)}}, rows: 1, cols: 4},
		},
		{
			&matrix{matrix: [][]float64{{-2.3, 3.3}, {1.2, -4.0}}, rows: 2, cols: 2},
			&matrix{matrix: [][]float64{{derivativeReLU(-2.3), derivativeReLU(3.3)}, {derivativeReLU(1.2), derivativeReLU(-4.0)}}, rows: 2, cols: 2},
		},
	}
	for _, table := range tables {
		actual, _ := DerivativeReLU(table.a)
		v, _ := actual.(*matrix)
		if !equalMatrices(table.expectedMatrix, v) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedMatrix, v)
		}
	}
}

func TestMultiplyScalar(t *testing.T) {
	tables := []struct {
		a              *matrix
		scalar         float64
		expectedMatrix *matrix
	}{
		{
			&matrix{matrix: [][]float64{{1, 2, 3, 4}}, rows: 1, cols: 4},
			5,
			&matrix{matrix: [][]float64{{5, 10, 15, 20}}, rows: 1, cols: 4},
		},
		{
			&matrix{matrix: [][]float64{{-2.3, 3.3}, {1.2, -4.0}}, rows: 2, cols: 2},
			0.5,
			&matrix{matrix: [][]float64{{-1.15, 1.65}, {0.6, -2}}, rows: 2, cols: 2},
		},
	}
	for _, table := range tables {
		actual := MultiplyScalar(table.a, table.scalar)
		v, _ := actual.(*matrix)
		if !equalMatrices(table.expectedMatrix, v) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedMatrix, v)
		}
	}
}

func TestSumByColumns(t *testing.T) {
	tables := []struct {
		a              *matrix
		expectedMatrix *matrix
	}{
		{
			&matrix{matrix: [][]float64{{1, 2, 3, 4}}, rows: 1, cols: 4},
			&matrix{matrix: [][]float64{{10}}, rows: 1, cols: 1},
		},
		{
			&matrix{matrix: [][]float64{{-2.3, 3.3}, {1.2, -4.0}}, rows: 2, cols: 2},
			&matrix{matrix: [][]float64{{1}, {-2.8}}, rows: 2, cols: 1},
		},
	}
	for _, table := range tables {
		actual := SumByColumns(table.a)
		v, _ := actual.(*matrix)
		if !equalMatrices(table.expectedMatrix, v) {
			t.Errorf("Expected: %v, Actual: %v\n", table.expectedMatrix, v)
		}
	}
}
