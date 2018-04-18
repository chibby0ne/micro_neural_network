package matrix

import (
	"fmt"
	"testing"
)

func equalMatrices(a, b *matrix) bool {
	// check that they are null and equal
	if a == nil && b == nil {
		return true
	}
	// both of them must NOT null at this point otherwise they are not equal
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
			if a[i][j] != b[i][j] {
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
			fmt.Errorf("matrices of different dimensions can't be added"),
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
			fmt.Errorf("matrices of different dimensions can't be substracted"),
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
			fmt.Errorf("matrices of different dimensions can't be multiplied elementwise"),
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
