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
	var currentJobs uint64
	var successfulQueries uint64
	var totalQueries uint64
	maxConcurrency := 1

	fmt.Println("Dropping DNS Hammer...")

	// open list file for reading
	file, err := os.Open("domains.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// resolve each line
	scanner := bufio.NewScanner(file)
	// for each line (domain) in the file
	for scanner.Scan() {
		// if there are already too many goroutines, wait
		for currentJobs > uint64(maxConcurrency) {
			time.Sleep(1)
		}
		// keep track of how many goroutines are running
		atomic.AddUint64(&currentJobs, 1)
		go func() {
			// at the end, subtract 1 from the wait group
			defer atomic.AddUint64(&currentJobs, ^uint64(0))
			err := resolve(ctx, myResolver, scanner.Text())
			if err == nil {
				atomic.AddUint64(&successfulQueries, 1)
			}
			atomic.AddUint64(&totalQueries, 1)
			fmt.Println(fmt.Sprintf("Success Rate: %.2f", float64(successfulQueries)/float64(totalQueries)))
		}()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// wait for all goroutines to finish
	for currentJobs > 0 {
		time.Sleep(1)
	}
}

func resolve(ctx context.Context, myResolver net.Resolver, host string) error{
	_, err := myResolver.LookupIPAddr(ctx, host)
	if err != nil {
		//fmt.Println(err)
		return err
	} else {
		// when verbose logging is implemented, these should be enabled only with verbose logging
		//fmt.Printf("%s:    ", host)
		//fmt.Printf("%s\n", answer[len(answer)-1])
		return nil
	}
}
