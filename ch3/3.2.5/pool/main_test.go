package main

import (
	"io/ioutil"
	"net"
	"testing"
)

func init() {
	deamonStartedSimply := startNetworkDeamonSimply()
	deamonStartedSimply.Wait()

	deamonStartedWithPool := startNetworkDeamonWithPool()
	deamonStartedWithPool.Wait()
}

func BenchmarkNetworkRequest8080(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}

func BenchmarkNetworkRequest8081(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8081")
		if err != nil {
			b.Fatalf("cannot dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}
