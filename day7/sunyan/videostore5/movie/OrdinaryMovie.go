package movie

type OrdinaryMovie struct {
	Title string //片名
}

func (m OrdinaryMovie) GetTitle() string {
	return m.Title
}
func (m OrdinaryMovie) GetCharge(daysRented int) float64 {
	result := 0.0
	result += 2
	if daysRented > 2 {
		result += float64(daysRented-2.0) * float64(1.5)
	}
	return result
}
func (m OrdinaryMovie) GetFrequentRenterPoints(daysRented int) int {
	return 1
}
