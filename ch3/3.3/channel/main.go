package main

import (
	"fmt"
)

func simplest() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channeles!"
	}()
	fmt.Println(<-stringStream)
}

// write から読み込んだり、readにかきこんだりするとコンパイルエラーになる
// func errorSample() {
// 	writeStream := make(chan<- interface{})
// 	readStream := make(<-chan interface{})

// 	<-writeStream
// 	readStream <- struct{}{}
// }

//lint:ignore U1000 テストコードから呼び出すと、無限に待ち続けるため
func deadlock() {
	stringStream := make(chan string)
	go func() {
		if 0 != 1 { // 書き込みが絶対に発生しないようにする
			return
		}
		stringStream <- "Hello channeles!"
	}()
	fmt.Println(<-stringStream)
}

func openChannel() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channeles!"
	}()
	salutation, ok := <-stringStream
	fmt.Printf("(%v): %v\n", ok, salutation)
}

func closeChannel() {
	intStream := make(chan int)
	close(intStream)
	integer, ok := <-intStream
	fmt.Printf("(%v): %v\n", ok, integer)
}

func flashChannel() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}
}
