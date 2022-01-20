// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var i int

// func worker(wg *sync.WaitGroup) {
// 	i = i + 1
// 	wg.Done()
// }

// func main() {

// 	var wg sync.WaitGroup

// 	for j := 0; j < 1000; j++ {

// 		wg.Add(1)
// 		go worker(&wg)

// 	}

// 	fmt.Println("Value of i after 1000 iteration is : ", i)
// }

//Lets use Mutex

package main

import (
	"fmt"
	"sync"
)

var i int

func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	i += 1
	m.Unlock()
	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	var m sync.Mutex

	for j := 0; j < 1000; j++ {
		wg.Add(1)
		worker(&wg, &m)
	}

	fmt.Println("Value of i after 1000 iteration : ", i)
}
