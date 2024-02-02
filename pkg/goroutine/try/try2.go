package try

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func Try2() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s, c)
	x := <-c // receive from c

	fmt.Println(x)
}
