package correlation

import "math"

func Correlation(y1, y2 []float64, phase float64) float64 {
	acc := 0.0

	len := min(len(y1), len(y2)-int(phase))

	for i := 0; i < len; i++ {
		acc += y1[i] * y2[(i+int(phase))%len]
	}

	return acc / float64(len)
}

func AutoCorrelation(y []float64, phase float64) float64 {
	return Correlation(y, y, phase)
}

func NormalizedCorrelation(y1, y2 []float64, phase float64) float64 {
	return Correlation(y1, y2, phase) / math.Sqrt(AutoCorrelation(y1, 0)*AutoCorrelation(y2, 0))
}
