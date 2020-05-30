package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Println("Please Press A for CPU count, any other key to exit")

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()

	if err != nil {
		fmt.Println(err)
	}

	switch char {
	case 'A':
		helloPrinter(threadCounter())
	default:
		fmt.Println("an unexpected key Pressed")
	}
}

func threadCounter() int {
	threadCount := runtime.NumCPU()
	return threadCount
}

func helloPrinter(threadCount int) {
	if threadCount >= 10 {
		fmt.Fprintf(os.Stderr, "Too many CPUs, I can't believe this! %v\n", threadCount)
		os.Exit(1)
	}

	for i := 1; i <= threadCount; i++ {
		fmt.Printf("Hello, World. Hello from CPU thread %d \n", i)
	}
}
