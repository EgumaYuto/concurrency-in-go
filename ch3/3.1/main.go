package main

import (
	"fmt"
	"sync"
)

func main() {
	// var wg sync.WaitGroup
	// salutatoin := "hello"
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	salutatoin = "welcom"
	// }()

	// wg.Wait()
	// fmt.Println(salutatoin)

	////////////////////////////////

	// var wg sync.WaitGroup

	// for _, salutatoin := range []string{"hello", "greetings", "good day"} {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		fmt.Println(salutatoin)
	// 	}()
	// }
	// wg.Wait()

	////////////////////////////////

	var wg sync.WaitGroup

	for _, salutatoin := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutatoin string) {
			defer wg.Done()
			fmt.Println(salutatoin)
		}(salutatoin)
	}
	wg.Wait()
}
