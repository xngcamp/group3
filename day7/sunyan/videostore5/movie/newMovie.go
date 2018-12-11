package movie

type NewMovie struct {
	Title     string //片名
}

func (m NewMovie) GetTitle() string {
	return m.Title
}
func (m NewMovie) GetCharge(daysRented int) float64 {
	result := 0.0
	result += float64(daysRented) * float64(3)
	return result
}
func (m NewMovie) GetFrequentRenterPoints(daysRented int) int {
	return 2
}