package gx
// Simple ternary operator: result:=iff(condition, result_if_true, result_if_false)
func iff[T interface{}](condition bool, a T, b T) T {
	if condition {
		return a
	}
	return b
}
