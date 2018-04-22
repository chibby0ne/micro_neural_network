package model

import (
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

// Initializes the models parameters (W, B)
func initializeParameters(hyperparameters *Hyperparameters) {
	rand.Seed(1)
	param := new(Parameters)
	param.W1 = matrix.NewRandomMatrix(hyperparameters.numHiddenUnits, numberTrainingExamples, 0.01)
	param.B1 = matrix.NewInitializedMatrix(hyperparameters.numHiddenUnits, numberFeatures, 0)
	param.W2 = matrix.NewRandomMatrix(1, hyperparameters.numHiddenUnits, 0.01)
	param.B2 = matrix.NewInitializedMatrix(1, 1, 0)
	return param
}

// One forward propragation step on the entire training set
func forwardPropagation(parameters *Parameters, X matrix.NumberArray) (A2 matrix.NumberArray, cache *Cache) {
	W1 = parameters.W1
	B1 = parameters.B1
	W2 = parameters.W2
	B2 = parameters.B2

	Z1 := Add(Dot(W1, X), B1)
	A1 := Sigmoid(Z1)
	Z2 := Add(Dot(W2, A1), B2)
	A2 = Sigmoid(Z2)

	cache := new(Cache)
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
	m := Y.getColumns()
	OnesMatrix := NewInitializedMatrix(Y.getRows(), Y.getColumns(), 1)
	logProbs := Add(MultiplyElementwise(Log(A2), Y), MultiplyElementwise(Substract(OnesMatrix, Y), Log(Substract(OnesMatrix, A2))))
	cost = -1 * Add(logProbs) / m  // along columns
	return cost
}


// One backward propagation step from output to input
func backwardPropagation(parameters *Parameters, cache *Cache, X matrix.NumberArray, Y matrix.NumberArray) *Gradients {
	m := Y.getColumns()
	OnesMatrix := NewInitializedMatrix(cache.A1.GetRows(), cache.A1.GetColumns())

	dZ2 = cache.A2 - Y
	dW2 = (Dot(dZ2, Cache.A1.Transpose()) / m
	dB2 = Add(dZ2) / m // along columns
	temp := OnesMatrix - Power(A1, 2)
	dZ1 = MultiplyElementwise(Dot(parameters.W2.Transpose(), dZ2), temp)
	dW1 = Dot(dZ1, X.Tranpose())
	dB1 = np.sum(dZ1, axis=1, keepdims=True) / m // along columns
}

// Updates the parameters according to Gradient descent
func updateParameters(parameters *Parameters, grads *Gradients, learningRate float64) *Parameters {
	W1 := parameters.W1
	B1 := parameters.B1
	W2 := parameters.W2
	B2 := parameters.B2

	W1 = W1 - learningRate * grads.dW1
	B1 = B1 - learningRate * grads.dB1
	W2 = W2 - learningRate * grads.dW2
	B2 = B2 - learningRate * grads.dB2

	parameters.W1 = W1
	parameters.B1 = B1
	parameters.W2 = W2
	parameters.B2 = B2

	return parameters
}

func Model(X, Y matrix.NumberArray, hyperparameters *Hyperparameters) *Parameters {
	parameters := initializeParameters()


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
		if i % 1000 {
			fmt.Println(cost)
		}
	}
	return parameters
}

// Predicts a binary classification task from the given input
func Predict(parameters Parameters, input matrix.NumberArray) bool {
	yHat, cache := forwardPropagation(parameters, input)
	return yHat.Get(0,0) > 0.5
}
