package Temperature

import "math/big"

type Temperature big.Float

func NewTemperature(x float64) *Temperature {
	z := Temperature(*big.NewFloat(x))
	return &z
}

func ParseString(t string) (*Temperature, error) {
	converted, _, err := big.ParseFloat(t, 10, 0, big.ToNearestEven)

	if err != nil {
		return nil, err
	}

	convertedTemperature := Temperature(*converted)
	return &convertedTemperature, nil
}

func (t *Temperature) Float64() (float64, big.Accuracy) {
	z := big.Float(*t)
	return z.Float64()
}
