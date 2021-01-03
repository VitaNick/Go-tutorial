package main

import (
	"fmt"
	"movie"
)

func main() {
	fmt.Println("My favorite movie")

	var myFav movie.Catalogable = &movie.Movie{}
	myFav.NewMovie("Top Gun", movie.PG, 32.8)

	fmt.Printf("My favorite movie is %s\n", myFav.GetTitle())
	fmt.Printf("It was rated %v\n", myFav.GetRating())
	fmt.Printf("It made %f in the box office\n", myFav.GetBoxOffice())

	myFav.SetTitle("Dumb and Dumber")

	fmt.Printf("My favorite movie is now %s\n", myFav.GetTitle())

}
