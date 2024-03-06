package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main2() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-c:
				fmt.Println("Получен сигнал ^C")
				fmt.Println("Выхожу из программы")
				os.Exit(0)
			default:
			
				for i := 1; ; i++ {
					fmt.Println(i * i)
				}
			}
		}
	}()

	
	select {}
}
