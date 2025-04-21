package main
import (
    "fmt"
    "time"
    )

//Implementing worker pool

//Worker function to process jobs
func worker(id int, jobs <-chan int, results chan<- int){
    for job := range jobs{
        fmt.Println("Worker ",id," started job ",job)
        time.Sleep(time.Second)
        results<-job*2
        fmt.Println("Worker ",id," finished job ",job)
    }
}

func mainGoroutine() {
  //Total number of jobs
  const numJob = 6
  //Task Queue: channel for jobs
  jobs:=make(chan int,numJob)
  //Result Channel: channel for processed jobs
  results:=make(chan int,numJob)
  
  //starting 3 worker goroutines
  for i:=1; i<=3; i++{
      go worker(i,jobs,results)
  }
  
  //sending jobs into the task queue
  for i:=1; i<=numJob; i++{
      jobs<-i
  }
  close(jobs)
  
  //fetching processed jobs from the results channel
  for i:=1; i<=numJob; i++{
      <-results
  }
  close(results)
}