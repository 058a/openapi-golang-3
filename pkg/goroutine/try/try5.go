package try

import "fmt"

// 関数型を引数として受け取る関数
func executeFunction(fn interface{}, args ...interface{}) {
	switch f := fn.(type) {
	case func(int, int) int:
		// int型の引数を2つ受け取り、int型を返す関数の場合
		result := f(args[0].(int), args[1].(int))
		fmt.Println("Result:", result)
	case func(string, string) string:
		// int型の引数を2つ受け取り、int型を返す関数の場合
		result := f(args[0].(string), args[1].(string))
		fmt.Println("Result:", result)
	default:
		fmt.Println("Unsupported function type")
	}
}

// 加算を行う関数
func add(x, y int) int {
	return x + y
}

// 文字列を結合する関数
func concatenate(a, b string) string {
	return a + b
}

func Try5() {
	// executeFunctionを使用して、異なる型の関数を呼び出す
	executeFunction(add, 3, 5)
	executeFunction(concatenate, "Hello, ", "World!")
}
