package main

import "fmt"

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		fmt.Println(i)
		ch <- i
	}
}

func Filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src {
		if i%prime != 0 {
			dst <- i
		}
	}
}

func main() {

	src := make(chan int)
	go Generate(src)

	for i := 0; i < 100; i++ {
		prime := <-src
		fmt.Println(prime)
		dst := make(chan int)
		go Filter(src, dst, prime)
		src = dst
	}
}
