package main

import (
	"bytes"
	"fmt"
	"os"
	"sync"
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

func flashChannelByClosing() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Printf("%v has begin\n", i)
		}(i)
	}

	fmt.Println("Unblocking groutines...")
	close(begin)
	wg.Wait()
}

func bufferedChannels() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 5; i++ {
			fmt.Fprintf(&stdoutBuff, "Sending %d\n", i)
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}

func main() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Recieved: %d\n", result)
	}
	fmt.Println("Done receiving!")
}
