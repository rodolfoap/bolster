package gx
// Simple ternary operator: result:=iff(condition, result_if_true, result_if_false)
func Iff[T interface{}](condition bool, a T, b T) T {
	if condition {
		return a
	}
	return b
}

// ProgressiveMean: calculate an averge without having the elements count
type ProgressiveMean struct {
	C   int
	Avg float64
}

func (m *ProgressiveMean) UpdateMean(x float64) {
	m.C+=+1
	m.Avg=(m.Avg*(float64(m.C)-1)+x)/float64(m.C)
}
