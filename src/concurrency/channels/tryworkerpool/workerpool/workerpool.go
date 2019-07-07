package workerpool

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var jobs chan Job
var results chan Result

const poolDepth = 10

func init() {
	fmt.Println("workerpool package init()")
	jobs = make(chan Job, poolDepth)
	results = make(chan Result, poolDepth)
}

//Job struct represents incoming data
type Job struct {
	id       int
	randomno int
}

//Result struct represents outcoming data
type Result struct {
	job         Job
	sumofdigits int
}

//getDigitsFromNumber decomposes "number" to digits and puts them to "output" channel
func getDigitsFromNumber(number int, output chan int) {
	if output != nil {
		for number != 0 {
			digit := number % 10
			output <- digit
			number /= 10
		}
		close(output)
	}
}

//summator summarises all digits of number and puts summary to output channel
func summator(number int, output chan int) {
	if output != nil {
		sum := 0
		digit := make(chan int)
		go getDigitsFromNumber(number, digit)

		for v := range digit {
			sum += v
		}
		//time.Sleep(2 * time.Second)
		output <- sum
	}
}

//worker runs summator function for each jobs in buffered channel "jobs" and puts summaries to "results" buffered channel
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		sumchan := make(chan int)
		go summator(job.randomno, sumchan)
		outputValue := Result{job, <-sumchan}
		results <- outputValue
	}
	wg.Done()
}

//createWorkerPool runs workers number of "population"
//"population" should be positive digit
func createWorkerPool(population byte) {
	var wg sync.WaitGroup
	var i byte
	for ; i < population; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

//allocate creates pool of jobs by size "jobsHeapSize" and puts random numbers to jobs data
func allocate(jobsHeapSize int) {
	for i := 0; i < jobsHeapSize; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}
	close(jobs)
}

//result func prints "results" buffered channel data
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random %d, sum of digits %d", result.job.id, result.job.randomno, result.sumofdigits)
		fmt.Println()
	}
	done <- true
}

//RunWorkerPool runs worker pool
func RunWorkerPool(numOfJobs int, numOfWorkers byte) {
	if (numOfJobs * int(numOfWorkers)) > 0 {
		startTime := time.Now()
		go allocate(numOfJobs)
		done := make(chan bool)
		go result(done)
		createWorkerPool(numOfWorkers)
		<-done
		endTime := time.Now()
		diffTime := endTime.Sub(startTime)
		fmt.Println("total time taken ", diffTime.Seconds(), " seconds")
	}
}
