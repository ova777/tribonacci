package lib

// matrix33 is array to store matrix 3x3
type matrix33 [3][3]uint

// Tribonacci calculates tribonacci number using matrix exponentiation
func Tribonacci(n uint) uint {
	matrix := matrix33{{1, 1, 1}, {1, 0, 0}, {0, 1, 0}}
	power := matrix33{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}

	for e := n; e > 0; e /= 2 {
		if e%2 != 0 {
			power = multiply(power, matrix)
		}
		matrix = multiply(matrix, matrix)
	}

	return power[2][0]
}

// multiply multiplies two matrices
func multiply(a, b matrix33) (result matrix33) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return
}
