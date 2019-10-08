package main

import (
	"bufio"
	"fmt"
	"context"
	"net"
	"log"
	"os"
	"sync/atomic"
	"time"
)

func main() {
	ctx := context.Background()
	myResolver := net.Resolver{}
	var gos uint64
	CONCURRENCYCOUNT := 1

	fmt.Println("Dropping DNS Hammer...")

	// open list file for reading
	file, err := os.Open("/domains.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// resolve each line
	scanner := bufio.NewScanner(file)
	// for each line in the file
	for scanner.Scan() {
		for gos > uint64(CONCURRENCYCOUNT) {
			time.Sleep(1)
		}
		// add 1 to our wait group
		atomic.AddUint64(&gos, 1)
		go func() {
			defer atomic.AddUint64(&gos, ^uint64(0))
			resolve(ctx, myResolver, scanner.Text())
		}()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// wait for all goroutines to finish
	for gos > 0 {
		time.Sleep(1)
	}
}

func resolve(ctx context.Context, myResolver net.Resolver, host string) {
	answer, err := myResolver.LookupIPAddr(ctx, host)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s:    ", host)
		fmt.Printf("%s\n", answer[len(answer)-1])
	}
}
