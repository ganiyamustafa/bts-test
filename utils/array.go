package utils

// Map applies the provided function to each element of the slice and returns a new slice with the results.
func Map[T any, U any](slice []T, fn func(int, T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(i, v)
	}

	return result
}
