package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(worker int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range jobs {
		fmt.Printf("[Worker %d] Uploading image on url: %s\n", worker, url)
		time.Sleep(time.Second * 1)
		results <- fmt.Sprintf("[Worker %d] Successfully uploaded image over url: %s", worker, url)
	}
}

func main() {
	var wg sync.WaitGroup
	noOfJobs := 100
	noOfWorkers := 10
	jobs := make(chan string, 100)
	results := make(chan string, 10)

	urls := [100]string{}
	for i := range 100 {
		urls[i] = fmt.Sprintf("www.example%d.com", i)
	}

	for i := range noOfWorkers {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	for i := range noOfJobs {
		jobs <- urls[i]
	}

	close(jobs)
	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}

}
