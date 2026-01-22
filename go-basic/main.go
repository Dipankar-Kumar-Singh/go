package main

import (
	"fmt"
	"math/rand/v2"
)

type D struct {
	v map[int]int
}

func dataSend(c chan int) {
	c <- rand.Int()
}

func main() {
	fmt.Println("Yooo")

	m := make(map[int]string, 10)
	m[5] = "dipankar"

	fmt.Println("yoo", "hnn")

	mm := make(map[int]map[int]string)
	mm[5] = map[int]string{
		5: "helll",
	}

	val := mm[5][5]
	fmt.Println("value ::", val)

	data := D{
		v: map[int]int{
			5: 345,
		},
	}

	c := make(chan int)
	dataSend(c)

	value := <-c

	fmt.Println("data : ", data)

	fmt.Println("Hello .. ", value)
}
