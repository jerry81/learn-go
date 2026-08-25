package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex) // also use make
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])

	// map literals
	var m2 = map[string]Vertex{
		"Bell Labs": Vertex{ // <- can be omitted
			40.68433, -74.39967,
		},
		"Google": { // <-- like this
			37.42202, -122.08408,
		},
	}
	fmt.Println(m2["Google"])
}
