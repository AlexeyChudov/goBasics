package goroutine

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu      sync.RWMutex
	counter map[string]int
}

func (c *Counter) Inc(value string) {
	//Если не использовать Lock и Unlock,
	//то созданные горутины будут одновременно обращаться к мапе, что приведет к ошибкам
	c.mu.Lock()
	c.counter[value]++
	c.mu.Unlock()
}

func (c *Counter) Value(key string) int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.counter[key]
}

func MutexExample() {
	c := Counter{counter: make(map[string]int)}

	c.counter["test"] = 0
	for i := 0; i < 1000; i++ {

		c.Inc("test")

	}
	//time.Sleep(3 * time.Second) //Без задержки результат будет < 1000

	fmt.Println(c.Value("test"))
}

func WaitGroup() {
	wg := sync.WaitGroup{}
	var ct uint64
	//wg.Add(1) // Приведет к ошибке т.к. wg.Wait() не дождется закрытия всех горутин
	for i := 0; i < 10; i++ {
		// wg.Add(1)
		k := i
		if i > 0 {
			wg.Add(1) //если горутин будет меньше, чем попуток их закрыть, то будет ошибка
			// negative WaitGroup counter
		}
		go func() {
			defer wg.Done()
			fmt.Printf("%d goroutine is working\n", k)
			time.Sleep(300 * time.Millisecond)
			// for j := 0; j < 1000; j++ {
			// 	// ct++ Потоконебезопасно, значения потеряются
			// 	atomic.AddUint64(&ct, 1) //потокобезопасный счетчик(прибавления не потеряются)

			// }

		}()
	}
	wg.Wait()
	fmt.Println("all done")
	fmt.Println("ct: ", ct)
}

func OnceFunc() {

	once := sync.Once{}
	done := make(chan bool)

	for i := 0; i < 1213; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("It is printed only once")
			})
			done <- true
		}()

	}
}
