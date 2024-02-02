package try

import (
	"fmt"
	"sync"
)

func Try1() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		Process1()
		wg.Done()
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		Process2()
		wg.Done()
	}(wg)

	wg.Wait()
	fmt.Println("End")
}

func Process1() {
	fmt.Println("Process1")
}

func Process2() {
	fmt.Println("Process2")
}
