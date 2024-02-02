package try

import (
	"fmt"
	"sync"
)

func Try4() {
	// チェックの結果を受け取るチャネル
	resultCh := make(chan error, 10) // バッファリングされたチャネルを使用

	wg := &sync.WaitGroup{}

	a := 2
	b := 4
	// 例として3つのゴルーチンを起動
	wg.Add(1)
	go func() {
		defer wg.Done()
		resultCh <- performCheck1(a)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		resultCh <- performCheck2(b)
	}()

	// 全てのゴルーチンの終了を待つ
	wg.Wait()

	// チャネルを閉じる（これが終了条件となる）
	close(resultCh)

	// エラーの数をカウント
	errorCount := 0
	for err := range resultCh {
		if err == nil {
			continue
		}
		fmt.Println("Error:", err)
		errorCount++
	}

	fmt.Println("Total Errors:", errorCount)
}

// チェックを模倣する関数
func performCheck1(id int) error {
	if id%2 == 0 {
		// エラーメッセージを小文字で始める
		return fmt.Errorf("error in check1 %d", id)
	}
	return nil
}
func performCheck2(id int) error {
	if id%2 == 0 {
		// エラーメッセージを小文字で始める
		return fmt.Errorf("error in check2 %d", id)
	}
	return nil
}
