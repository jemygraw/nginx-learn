package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
)

func doWork(tasks chan func()) {
	for {
		task := <-tasks
		task()
	}
}

func main() {
	var host string
	var limit int
	var worker int

	var port8080cnt int32 = 0
	var port8081cnt int32 = 0
	var port8082cnt int32 = 0
	var port8083cnt int32 = 0

	flag.StringVar(&host, "host", "localhost", "host to connect")
	flag.IntVar(&limit, "limit", 100, "request count")
	flag.IntVar(&worker, "worker", 1, "routine worker")
	flag.Parse()

	tasks := make(chan func(), worker)
	wg := sync.WaitGroup{}
	for i := 0; i < worker; i++ {
		go doWork(tasks)
	}

	for i := 0; i < limit; i++ {
		wg.Add(1)
		tasks <- func() {
			defer wg.Done()
			resp, err := http.Get("http://" + host)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer resp.Body.Close()

			respData, _ := ioutil.ReadAll(resp.Body)
			respStr := string(respData)

			if strings.Contains(respStr, "8080") {
				atomic.AddInt32(&port8080cnt, 1)
			} else if strings.Contains(respStr, "8081") {
				atomic.AddInt32(&port8081cnt, 1)
			} else if strings.Contains(respStr, "8082") {
				atomic.AddInt32(&port8082cnt, 1)
			} else if strings.Contains(respStr, "8083") {
				atomic.AddInt32(&port8083cnt, 1)
			} else {
				fmt.Println(respStr)
			}
		}
	}

	wg.Wait()
	fmt.Println("8080:", port8080cnt)
	fmt.Println("8081:", port8081cnt)
	fmt.Println("8082:", port8082cnt)
	fmt.Println("8083:", port8083cnt)
}
