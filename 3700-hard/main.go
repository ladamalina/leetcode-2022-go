package main

const mod int64 = 1e9 + 7

type Matrix [][]int64

func NewMatrix(n, m int) Matrix {
	mat := make(Matrix, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int64, m)
	}
	return mat
}

func Identity(n int) Matrix {
	I := NewMatrix(n, n)
	for i := 0; i < n; i++ {
		I[i][i] = 1
	}
	return I
}

func MatMul(A, B Matrix) Matrix {
	n := len(A)
	if n == 0 {
		return Matrix{}
	}
	m := len(B[0])
	p := len(B) // columns of A / rows of B
	C := NewMatrix(n, m)
	// standard triple loop, optimized by skipping zero A[i][k]
	for i := 0; i < n; i++ {
		for k := 0; k < p; k++ {
			av := A[i][k]
			if av == 0 {
				continue
			}
			for j := 0; j < m; j++ {
				C[i][j] += av * B[k][j]
				if C[i][j] >= (1<<63-1)/4 { // avoid potential overflow before mod (conservative)
					C[i][j] %= mod
				}
			}
		}
		for j := 0; j < m; j++ {
			C[i][j] %= mod
		}
	}
	return C
}

func CopyMatrix(A Matrix) Matrix {
	n := len(A)
	if n == 0 {
		return Matrix{}
	}
	m := len(A[0])
	C := NewMatrix(n, m)
	for i := 0; i < n; i++ {
		copy(C[i], A[i])
	}
	return C
}

func MatPow(A Matrix, e int) Matrix {
	n := len(A)
	if n == 0 {
		return Matrix{}
	}
	result := Identity(n)
	base := CopyMatrix(A)
	for e > 0 {
		if e&1 == 1 {
			result = MatMul(result, base)
		}
		base = MatMul(base, base)
		e >>= 1
	}
	return result
}

func zigZagArrays(n, l, r int) int {
	K := int(r - l + 1)
	if K <= 0 {
		return 0
	}

	// M1[i][j] = 1 if i < j else 0
	M1 := NewMatrix(K, K)
	for i := 0; i < K; i++ {
		for j := 0; j < K; j++ {
			if i < j {
				M1[i][j] = 1
			} else {
				M1[i][j] = 0
			}
		}
	}

	// M2[i][j] = K - 1 - max(i, j)
	M2 := NewMatrix(K, K)
	for i := 0; i < K; i++ {
		for j := 0; j < K; j++ {
			maxij := i
			if j > i {
				maxij = j
			}
			M2[i][j] = int64(K-1-maxij) % mod
		}
	}

	pairs := (n - 1) / 2
	rem := (n - 1) % 2

	A := MatPow(M2, pairs)
	if rem != 0 {
		A = MatMul(A, M1)
	}

	var total int64 = 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			total += A[i][j]
			if total >= (1<<63-1)/4 {
				total %= mod
			}
		}
	}
	total %= mod
	total = (total * 2) % mod
	if total < 0 {
		total += mod
	}
	return int(total)
}
