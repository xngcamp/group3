package movie
type newmovie struct {
	Title string
}
func (n newmovie) getTitle() string{
	return n.Title
}

func (n newmovie) GetCharge(daysRented int) float64{
	result := 0.0
	result += float64(daysRented) * float64(3)
	return result
}

func (n newmovie) GetFrequentRenterPoints(daysRented int) int{
	if daysRented > 1 {
		return 2
	}else {
		return 1
	}
}

