package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

type Job struct {
	server string
	port   int
}

var available []int
var jobs = make(chan Job, 10)

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		ip := job.server + ":" + strconv.Itoa(job.port)
		fmt.Println(ip)

		_, err := net.DialTimeout("tcp", ip, time.Duration(300)*time.Millisecond)
		if err != nil {
		} else {
			available = append(available, job.port)
		}
	}

	wg.Done()
}

func PortScan(done chan bool, server string) {
	for i := 1; i <= 65535; i++ {
		job := Job{server, i}
		jobs <- job
	}
	close(jobs)
	done <- true
}

func main() {
	fmt.Println("Checking for available ports...")
	done := make(chan bool)
	go PortScan(done, os.Args[1])
	noOfWorkers := 10000
	createWorkerPool(noOfWorkers)
	<-done

	fmt.Println("Ports available: ", available)

}
