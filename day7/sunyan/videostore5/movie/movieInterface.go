package movie

type MovieInterface interface {
	GetTitle() string
	GetCharge(daysRented int) float64
	GetFrequentRenterPoints(daysRented int) int
}
