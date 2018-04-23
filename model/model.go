package model

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/chibby0ne/micro_neural_network/matrix"
)

// Learned Parameters of the Neural Network
// W1, B1 for the hidden layer
// W2, B2 for the output unit
type Parameters struct {
	W1 matrix.NumberArray
	B1 matrix.NumberArray
	W2 matrix.NumberArray
	B2 matrix.NumberArray
}

// Hyperparameters of the neural network. They control the behaviour of the
// network and ultimately the actual parameters' value
type Hyperparameters struct {
	numIterations   int
	learningRate    float64
	numHiddenLayers int
	numHiddenUnits  int
}

// Values calculated in the forward propagation step that are reused  in the
// backward propagation step's calculations
type Cache struct {
	Z1 matrix.NumberArray
	A1 matrix.NumberArray
	Z2 matrix.NumberArray
	A2 matrix.NumberArray
}

// Output of the backward propagation step. Used to update the parameters
// following gradient descent algorithm
type Gradients struct {
	dW2 matrix.NumberArray
	dB2 matrix.NumberArray
	dW1 matrix.NumberArray
	dB1 matrix.NumberArray
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Initializes the models parameters (W, B)
func initializeParameters(hyperparameters *Hyperparameters, numberTrainingExamples int, numberFeatures int) *Parameters {
	rand.Seed(1)
	param := new(Parameters)
	var err error
	param.W1, err = matrix.NewRandomMatrix(hyperparameters.numHiddenUnits, numberTrainingExamples, 0.01)
	handleError(err)
	param.B1, err = matrix.NewInitializedMatrix(hyperparameters.numHiddenUnits, numberFeatures, 0)
	handleError(err)
	param.W2, err = matrix.NewRandomMatrix(1, hyperparameters.numHiddenUnits, 0.01)
	handleError(err)
	param.B2, err = matrix.NewInitializedMatrix(1, 1, 0)
	handleError(err)
	return param
}

// One forward propragation step on the entire training set
func forwardPropagation(parameters *Parameters, X matrix.NumberArray) (A2 matrix.NumberArray, cache *Cache) {
	W1 := parameters.W1
	B1 := parameters.B1
	W2 := parameters.W2
	B2 := parameters.B2

	var err error
	W1X, err := matrix.Dot(W1, X)
	handleError(err)
	Z1, err := matrix.Add(W1X, B1)
	handleError(err)
	A1 := matrix.Tanh(Z1)

	W2A1, err := matrix.Dot(W2, A1)
	handleError(err)
	Z2, err := matrix.Add(W2A1, B2)
	handleError(err)
	A2 = matrix.Sigmoid(Z2)

	cache = new(Cache)
	cache.Z1 = Z1
	cache.A1 = A1
	cache.Z2 = Z2
	cache.A2 = A2
	return A2, cache
}

// Calculate the cross-entropy loss from the desired output and the actual
// output
// J=−1m∑i=0m(y(i)log(a[2](i))+(1−y(i))log(1−a[2](i)))
func computeCost(A2 matrix.NumberArray, Y matrix.NumberArray) float64 {
	m := Y.GetColumns()

	var err error
	OnesMatrix, err := matrix.NewInitializedMatrix(Y.GetRows(), Y.GetColumns(), 1)
	handleError(err)

	OneMinusA2, err := matrix.Substract(OnesMatrix, A2)
	handleError(err)
	OneMinusY, err := matrix.Substract(OnesMatrix, Y)
	handleError(err)
	OneMinusYTimesLogOneMinusA2, err := matrix.MultiplyElementwise(OneMinusY, matrix.Log(OneMinusA2))
	handleError(err)

	LogA2TimesY, err := matrix.MultiplyElementwise(matrix.Log(A2), Y)
	handleError(err)

	logProbs, err := matrix.Add(LogA2TimesY, OneMinusYTimesLogOneMinusA2)
	handleError(err)
	// logProbs := matrix.Add(matrix.MultiplyElementwise(matrix.Log(A2), Y), matrix.MultiplyElementwise(matrix.Substract(OnesMatrix, Y), matrix.Log(matrix.Substract(OnesMatrix, A2))))

	cost := matrix.MultiplyScalar(matrix.SumByColumns(logProbs), -1.0/float64(m)) // along columns

	// cost is a matrix therefore we need to index the value to return it
	val, err := cost.GetValue(0, 0)
	handleError(err)
	return val
}

// One backward propagation step from output to input
func backwardPropagation(parameters *Parameters, cache *Cache, X matrix.NumberArray, Y matrix.NumberArray) *Gradients {
	m := Y.GetColumns()
	var err error

	dZ2, err := matrix.Substract(cache.A2, Y)
	handleError(err)
	dZ2DotA1T, err := matrix.Dot(dZ2, cache.A1.Transpose())
	handleError(err)
	dW2 := matrix.MultiplyScalar(dZ2DotA1T, 1.0/float64(m))
	handleError(err)
	dB2 := matrix.MultiplyScalar(matrix.SumByColumns(dZ2), 1.0/float64(m)) // along columns
	W2TDotdZ2, err := matrix.Dot(parameters.W2.Transpose(), dZ2)
	handleError(err)
	dZ1, err := matrix.MultiplyElementwise(W2TDotdZ2, matrix.DerivativeTanh(cache.Z1))
	handleError(err)
	dW1, err := matrix.Dot(dZ1, X.Transpose())
	handleError(err)
	dB1 := matrix.MultiplyScalar(matrix.SumByColumns(dZ1), 1.0/float64(m)) // along columns

	grads := new(Gradients)
	grads.dW2 = dW2
	grads.dB2 = dB2
	grads.dW1 = dW1
	grads.dB1 = dB1

	return grads

}

// Updates the parameters according to Gradient descent
func updateParameters(parameters *Parameters, grads *Gradients, learningRate float64) *Parameters {
	W1 := parameters.W1
	B1 := parameters.B1
	W2 := parameters.W2
	B2 := parameters.B2

	var err error
	W1, err = matrix.Substract(W1, matrix.MultiplyScalar(grads.dW1, learningRate))
	handleError(err)
	B1, err = matrix.Substract(B1, matrix.MultiplyScalar(grads.dB1, learningRate))
	handleError(err)
	W2, err = matrix.Substract(W2, matrix.MultiplyScalar(grads.dW2, learningRate))
	handleError(err)
	B2, err = matrix.Substract(B2, matrix.MultiplyScalar(grads.dB2, learningRate))

	parameters.W1 = W1
	parameters.B1 = B1
	parameters.W2 = W2
	parameters.B2 = B2

	return parameters
}

func Model(X, Y matrix.NumberArray, hyperparameters *Hyperparameters, numberTrainingExamples, numberFeatures int) *Parameters {
	parameters := initializeParameters(hyperparameters, numberTrainingExamples, numberFeatures)

	for i := 0; i < hyperparameters.numIterations; i++ {

		// Forward prop
		A2, cache := forwardPropagation(parameters, X)

		// calculate cost
		cost := computeCost(A2, Y)

		// backward prop
		grads := backwardPropagation(parameters, cache, X, Y)

		// update params
		parameters = updateParameters(parameters, grads, hyperparameters.learningRate)

		// print cost every 1000 iterations
		if i%1000 == 0 {
			fmt.Printf("Cost after %v iterations: %v\n", i, cost)
		}
	}
	return parameters
}

// Predicts a binary classification task from the given input
func Predict(parameters *Parameters, input matrix.NumberArray) bool {
	yHat, _ := forwardPropagation(parameters, input)
	val, _ := yHat.GetValue(0, 0)
	return val > 0.5
}
