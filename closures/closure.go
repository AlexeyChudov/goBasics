package closures

// func main() {
// 	f := outter(1)
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(f())
// 	}
// }

func Outter(initial int) func(int) int {
	current := initial
	return func(x int) int {
		current += x
		return current
	}
}
