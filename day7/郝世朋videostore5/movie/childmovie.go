package movie
type childmovie struct {
	Title string
}
func (c childmovie) getTitle() string{
	return c.Title
}

func (c childmovie) GetCharge(daysRented int) float64{
	result := 0.0
	result += 1.5
	if daysRented > 3 {
		result += float64(daysRented-3) * float64(1.5)
	}
	return result
}

func (c childmovie) GetFrequentRenterPoints(daysRented int) int{
	return 1
}