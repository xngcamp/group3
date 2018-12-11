package movie
type movieinterface interface {
	GetTotalFrequentRenterPoints(daysRented int) int
	GetCharge(daysRented int) float64
	GetTitle() string
}
