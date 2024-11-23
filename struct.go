package main

import "fmt"

type Movie struct {
	name   string
	rating int
	comment string
	imdbLink string
}

func getMovie(name string, rating int) (movie Movie) {
	movie = Movie{
		name:   name,
		rating: rating,
		// comment: "init",
		// imdbLink: "",
	}
	return
}

func increaseRating(movie *Movie) {
	movie.rating += 1
}

func main() {
	mov := getMovie("xyz", 2)
	increaseRating(&mov)
	fmt.Printf("%+v", mov)
}
