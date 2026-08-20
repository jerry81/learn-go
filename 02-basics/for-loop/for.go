package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum2 := 1
	for sum2 < 1000 { // note start and step are optional
		sum2 += sum2
		if sum2 == 128 {
			fmt.Println("sum2 is 128")
		}
	}
	fmt.Println(sum2)
	/* this makes infinite loop
	for {}
	*/
}
