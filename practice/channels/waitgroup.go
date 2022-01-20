package main

import (
	"fmt"
	"sync"
	"time"
)

func service(wg *sync.WaitGroup, instance int) {

	time.Sleep(time.Second * 2)
	fmt.Println("Intance Called : ", instance)
	wg.Done()
}

func main() {

	fmt.Printf("main() started")

	var wg sync.WaitGroup

	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go service(&wg, i)
	}

	wg.Wait()
	fmt.Printf("main() ended")

}
