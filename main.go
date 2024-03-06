package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := make(chan int)
	squares := make(chan int)
	products := make(chan int)
	stop := make(chan struct{})
	wg := sync.WaitGroup{}

	
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			var input int
			fmt.Scan(&input)
			numbers <- input
			if input == 0 {
				close(stop)
				return
			}
		}
	}()

	
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range numbers {
			square := num * num
			squares <- square
		}
		close(squares)
	}()

	
	wg.Add(1)
	go func() {
		defer wg.Done()
		for square := range squares {
			product := square * 2
			products <- product
		}
		close(products)
	}()

	
	wg.Add(1)
	go func() {
		defer wg.Done()
		for product := range products {
			fmt.Println("Ввод:", product/2)
			fmt.Println("Квадрат:", product/2*product/2)
			fmt.Println("Произведение:", product)
		}
	}()

	
	wg.Wait()
}
