package rental

import "sunyan/videostore5/movie"

type Rental struct {
	aMovie     movie.MovieInterface
	daysRented int //租期
}

func (r *Rental) Init(movie movie.MovieInterface, daysRented int) {
	r.aMovie = movie
	r.daysRented = daysRented
}

func (r Rental) GetDaysRented() int {
	return r.daysRented
}

func (r Rental) GetMovie() movie.MovieInterface {
	return r.aMovie
}

func (r Rental) GetCharge() float64 {
	return r.GetMovie().GetCharge(r.daysRented)
}

func (r Rental) GetFrequentRenterPoints() int {
	return r.GetMovie().GetFrequentRenterPoints(r.daysRented)
}
