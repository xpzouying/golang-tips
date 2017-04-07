package main

import (
	"fmt"
	"time"
)

var (
	svrs []FakeService
)

type Result string
type FakeService func() Result

// handle request with *timeUsed* second time
func handleRequest(timeUsed int) {
	time.Sleep(time.Second * time.Duration(timeUsed))
}

// Fake service node, name and time used
// timeUsed is server node need time consume
func fakeServiceNode(name string, timeUsed int) FakeService {

	return func() Result {

		handleRequest(timeUsed)
		return Result(fmt.Sprintf("get result from server node: %s\n", name))
	}
}

func getAllResult() (results []Result) {

	c := make(chan Result)

	for _, s := range svrs {
		go func() {
			c <- s()
		}()
	}

	for _ = range svrs {
		results = append(results, <-c)
	}

	return
}

func getFastestResult(cluster []FakeService) (result Result) {

	c := make(chan Result)

	// --- way 1 ---
	// makeService := func(i int) {
	// 	c <- cluster[i]()
	// }
	//
	// for i := range cluster {
	// 	go makeService(i)
	// }

	// --- way 2 ---
	for i := range cluster {
		go func(i int) {
			c <- cluster[i]()
		}(i)
	}

	return <-c
}

func testGetAllResult() {
	fmt.Println("testGetAllResult")
	begin := time.Now()
	results := getAllResult()
	fmt.Println(results)
	fmt.Println("Time used: ", time.Since(begin))
}

func testGetFirstResult() {
	fmt.Println("testGetFirstResult")

	begin := time.Now()
	result := getFastestResult(svrs)
	fmt.Println(result)
	fmt.Println("Time used: ", time.Since(begin))
}

func main() {
	testGetFirstResult()
}

func init() {
	// create fake service server cluster
	for i := 1; i < 5; i++ {
		svrs = append(svrs, fakeServiceNode(fmt.Sprintf("node-%d", i), i))
	}
}
