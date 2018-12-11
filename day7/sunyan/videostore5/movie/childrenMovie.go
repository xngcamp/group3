package movie

type ChildrenMovie struct {
	Title     string //片名
}

func (m ChildrenMovie) GetTitle() string {
	return m.Title
}
func (m ChildrenMovie) GetCharge(daysRented int) float64 {
	result := 0.0
	result += 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * float64(1.5)
	}
	return result
}
func (m ChildrenMovie) GetFrequentRenterPoints(daysRented int) int {
	return 1
}
