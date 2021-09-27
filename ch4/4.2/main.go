package main

func main() {

	done := make(chan interface{})

	go func() {
		defer close(done)
	}()

	var stringStream chan<- string

	for _, s := range []string{"a", "b", "c"} {
		select {
		case <-done:
			return
		case stringStream <- s:
		default:
			// do nothing
		}
	}
}
