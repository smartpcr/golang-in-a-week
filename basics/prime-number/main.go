package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func generateNumbers(ch chan int) {
	for i := 2; i < 1_000_000; i++ {
		ch <- i
	}
}

func filterNumbers(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func generatePrimeNumbers() {
	// Create a channel to receive signals
	sig := make(chan os.Signal, 1)

	// Notify the channel on SIGINT (Ctrl+C) or SIGTERM
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	input := make(chan int)
	go generateNumbers(input)

	for {
		select {
		case <-sig:
			fmt.Println("Stopping...")
			return
		default:
			prime := <-input
			println(prime)
			output := make(chan int)
			go filterNumbers(input, output, prime)
			input = output
		}
	}
}

func main() {
	start := time.Now()
	generatePrimeNumbers()
	duration := time.Since(start)
	fmt.Println("Done!")
	fmt.Printf("duration: %.3f seconds", duration.Seconds())
}
