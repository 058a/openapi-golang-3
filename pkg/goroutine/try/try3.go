package try

import (
	"fmt"
)

func Try3() {
	// When
	ch := make(chan error)
	go func() {
		defer close(ch)
		ch <- Check1()
	}()

	err, ok := <-ch
	if !ok {
		// チャネルが閉じられていた場合の処理
		fmt.Println("Channel closed")
		return
	}

	// チャネルが閉じられていない場合の処理
	fmt.Println("&s", err)
}

func Check1() error {
	fmt.Println("Check1")
	return fmt.Errorf("Check1 Error")
}

func Check2() error {
	fmt.Println("Check2")
	return fmt.Errorf("Check2 Error")
}
