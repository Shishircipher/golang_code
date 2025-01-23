package main 

import "fmt"
import "encoding/json"
import "log"
func main () {


	type Movie struct {

		Title string
		Year int `json:"released"`
		Color bool `json:"released"`
		Actors []string

	}

	var movies = []Movie {
		{ Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphary Boagart", "Ingrid Bergman"}}, {Title: "Cool hand Luke", Year : 1967, Color: true, Actors: []string{"Paul Newman"}},{Title : "Bullit", Year: 1967, 
		Color: true, Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON Marshaling failed: %s",err)
	}
	fmt.Printf("%s\n", data)
}
