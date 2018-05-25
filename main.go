package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

//Work body of request
type Work struct {
	Data      string `json:"data"`
	TimeSleep int    `json:"time"`
}

var workers = 100

func main() {
	fmt.Println("Starting server on Port :8888")

	//handle for word extractor
	http.HandleFunc("/", getData)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}

}

//handler to handle incoming request
func getData(w http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	defer req.Body.Close()
	var work []Work
	err := decoder.Decode(&work)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s", "Invalid request!")
		return
	}
	//channel through which works are sent unit work
	jobs := make(chan Work, 100)

	//for each unit of work, the result is sent back from minion through out channel
	out := make(chan int64)

	//channel used sync go func ranging out channel and main go routine
	wai := make(chan struct{})

	//wait group variable to make sure that all minion exit before request in served and out channel is closed
	var waiter sync.WaitGroup
	waiter.Add(workers)

	//initialising the given number of workers
	for i := 0; i < workers; i++ {
		go func() {
			for job := range jobs {
				sum := 0
				for c := range job.Data {
					sum = sum + c
				}
				//sleeps for the specified amount of time
				time.Sleep(time.Duration(job.TimeSleep) * time.Millisecond)
				out <- int64(sum)
			}
			waiter.Done()
		}()
	}

	var sum int64
	//ranging over the out channel to aggregate the sum
	go func() {
		for r := range out {
			sum = sum + r
		}
		wai <- struct{}{}
	}()

	//sending work to the channel
	for _, val := range work {
		jobs <- val
	}
	close(jobs)
	waiter.Wait()
	close(out)
	<-wai
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fmt.Sprintf(`{"result": %d}`, sum))

}
