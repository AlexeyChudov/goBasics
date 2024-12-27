package goroutine

import (
	"fmt"
	"time"
)

func MainFunc() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)

}

func Goroutines() {
	// Ключевое слово go создает горутину - отдельный легковесный поток,
	//если основной поток заканчивает выполнение, то горутина тоже заканчивает
	go sayHello() // не выполнится, т.к. основной поток выполнится быстрее
	fmt.Println("main stream")
}

func sayHello() {
	time.Sleep(5 * time.Second)
	fmt.Println("Hello")
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func ChannelsExample() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func ChannelsExample2(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch) // Закрытие Канала, если этого не сделать,
	// то в том случае если где ожидаются данные из канала, и они больше туда не запишутся, то будет ошибка
}

func SelectData(data chan int, exit chan string) {
	var x int = 0
	for {
		select {
		case data <- x:
			x += 1

		case <-exit:
			fmt.Println("exit")
			return

		default:
			fmt.Println("Put something into data or exit")

		}
	}

}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
