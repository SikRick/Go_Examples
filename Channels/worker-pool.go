package main

import (
	"fmt"
	"time"
)

//worker is a sample worker.
//accepts jobs from jobs channel and pushes results to results channel
//time.Sleep represents an expensive task
func worker(id int, jobs <-chan int, res chan<- int) {

	fmt.Printf("Worker %d : started\n", id)
	for job := range jobs {
		fmt.Printf("worker %d :  received job %d\n", id, job)
		time.Sleep(time.Second)
		res <- job * job
		fmt.Printf("worker %d : finished job %d", id, job)
	}
	//close(results)
	//putting close here and having a range call over the results channel below wont work becuase,
	//the first worker to finish its jobs closes the results channel, and other workers wont be able to send their results
	//sending to a closed channel will cause panic

}

func main() {

	maxJobs := 10
	maxWorkers := 3
	jobs := make(chan int, maxJobs)
	res := make(chan int, maxJobs)
	for i := 0; i < maxJobs; i++ {
		jobs <- i
		fmt.Printf("pushed %d\n", i)
	}

	close(jobs)

	for j := 0; j < maxWorkers; j++ {
		go worker(j, jobs, res)
	}

	for k := 0; k < maxJobs; k++ {
		result := <-res
		fmt.Println(result)
	}
}
