package main

import "fmt"

func main(){

	c := make(chan int)
	for i := 0; i <= 9 ; i++ {

		go func(){
			fmt.Println(i)
			c <- i
		}()
	}

	for range c {
		// fmt.Println(v)
	}
}