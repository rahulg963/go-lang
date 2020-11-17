package concurrent

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// go run --race .
var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

// func ConcurrentLearningWithoutGoRoutine() {
// 	for i := 0; i < 10; i++ {
// 		// random numbers from 0 to 9
// 		id := rnd.Intn(10) + 1
// 		if b, ok := queryCache(id); ok {
// 			fmt.Println("from cache")
// 			fmt.Println(b)
// 			continue
// 		}
// 		if b, ok := queryDatabase(id); ok {
// 			fmt.Println("from database")
// 			fmt.Println(b)
// 			continue
// 		}

// 		fmt.Printf("Book not found with id : '%v'", id)
// 		time.Sleep(150 * time.Microsecond)
// 	}
// }

func ConcurrentLearningWithGoRoutine() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}

	for i := 0; i < 10; i++ {
		// random numbers from 0 to 9
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryCache(id, m); ok {
				fmt.Println("from cache")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryDatabase(id, m); ok {
				fmt.Println("from database")
				fmt.Println(b)
			}
			wg.Done()
		}(id, wg, m)

		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()
}

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	m.RLock()
	// shared memory
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	time.Sleep(100 * time.Microsecond)
	for _, b := range books {
		if id == b.ID {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}
	return Book{}, false
}
