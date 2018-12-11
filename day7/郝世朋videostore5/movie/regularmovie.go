package movie
type regularmovie struct {
	Title string
}
func (r regularmovie) getTitle() string{
	return r.Title
}

func (r regularmovie) GetCharge(daysRented int) float64{
	result := 0.0
	result += 2
	if daysRented > 2 {
		result += float64(daysRented-2.0) * float64(1.5)
	}
	return result
}

func (r regularmovie) GetFrequentRenterPoints(daysRented int) int{
	return 1
}
