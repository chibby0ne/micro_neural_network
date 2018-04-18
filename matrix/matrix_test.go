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
	fmt.Printf("a: %v, b: %v\n", a, b)
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

func TestNewMatrixPositiveBounds(t *testing.T) {
	// v := make([][]float64, 1)
	// v[0] = make([]float64, 1)
	// r := [][]float64{[]float64{0}}
	// m := &matrix{matrix: v, rows: 1, cols: 1}
	tables := []struct {
		rows           uint64
		cols           uint64
		expectedMatrix *matrix
		expectedError  error
	}{
		// {-1, -1, nil, fmt.Errorf("Can't create a matrix with -1 rows")},
		// {-1, 0, nil, fmt.Errorf("Can't create a matrix with -1 rows")},
		// {1, -1, nil, fmt.Errorf("Can't create a matrix with -1 cols")},
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
		m, err := newMatrix(table.rows, table.cols)
		if !equalMatrices(table.expectedMatrix, m) {
			t.Errorf("ExpectedMatrix: %v, ActualMatrix: %v\n", table.expectedMatrix, m)
		}
		if !equalErrors(table.expectedError, err) {
			t.Errorf("ExpectedError: %v, ActualError: %v", table.expectedError, err)
		}
	}
}
