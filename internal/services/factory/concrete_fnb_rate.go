package factory

var constantRate = 0.05

type FnBRateFactory struct{}

func (fnbRateFactory *FnBRateFactory) CalculateRate() float64 {

	return 0.07 * constantRate
}
