package main

import (
	"fmt"
	"time"
)

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}
	or = func(channeles ...<-chan interface{}) <-chan interface{} {
		switch len(channeles) {
		case 0:
			return nil
		case 1:
			return channeles[0]
		}

		orDone := make(chan interface{})
		go func() {
			defer close(orDone)
			switch len(channeles) {
			case 2:
				select {
				case <-channeles[0]:
				case <-channeles[1]:
				}
			default:
				select {
				case <-channeles[0]:
				case <-channeles[1]:
				case <-channeles[2]:
				case <-or(append(channeles[3:], orDone)...):
				}
			}
		}()
		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
	)
	fmt.Printf("done after %v", time.Since(start))
}
