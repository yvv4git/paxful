package helpers

const (
	percentageOneHalf = 1.5
)

// PercentageOneHalf is used for cal 1.5 procentage.
func PercentageOneHalf(sum float64) (result float64, err error) {
	if sum <= 0 {
		return 0.0, ErrIncorrectSum
	}

	result = sum + (sum / 100 * percentageOneHalf)
	return
}
