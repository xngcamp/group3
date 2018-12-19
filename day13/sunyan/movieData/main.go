package main

import (
	"encoding/json"
	"fmt"
	r "github.com/garyburd/redigo/redis"
	"movieData/customer"
	"movieData/movie"
	"movieData/redis"
	"time"
)

func main() {
	initMovieDb()

	conn := redis.GetRedisConn()
	defer conn.Close()

	aCustomer := new(customer.Customer)
	aCustomer.Init("Tom")
	movier := movie.RegularMovie{}
	movie1, _ := r.String(conn.Do("GET", "1"))
	json.Unmarshal([]byte(movie1), &movier)
	fmt.Println(movier)
	//rental1 := &rental.Rental{
	//	AMovie: movie.RegularMovie{
	//		movie.Param{Title: "The Godfather", PriceCode: movie.REGULAR},
	//	},
	//	DaysRented: 5,
	//}
	//aCustomer.AddRental(rental1)
	//
	//rental2 := &rental.Rental{
	//	AMovie: movie.NewIssueMovie{
	//		movie.Param{Title: "Fast & Furious", PriceCode: movie.NEW_ISSUE},
	//	},
	//	DaysRented: 1,
	//}
	//aCustomer.AddRental(rental2)
	//
	//result := aCustomer.Statement()
	//fmt.Println(result)
}
func initMovieDb() {
	conn := redis.GetRedisConn()
	defer conn.Close()

	movie1 := movie.RegularMovie{
		movie.Param{Id: "1", Title: "The Godfather", PriceCode: movie.REGULAR,
			Year: time.Now().Unix(), Author: "udaia"},
	}
	movie1Str, _ := json.Marshal(movie1)
	movie2 := movie.RegularMovie{
		movie.Param{Id: "2", Title: "Fast & Furious", PriceCode: movie.NEW_ISSUE,
			Year: time.Now().Unix(), Author: "dasdasd"},
	}
	movie2Str, _ := json.Marshal(movie2)
	conn.Do("set", movie1.Id, string(movie1Str))
	conn.Do("set", movie2.Id, string(movie2Str))

}
