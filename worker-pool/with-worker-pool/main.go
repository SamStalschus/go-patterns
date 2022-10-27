package main

import (
	"fmt"
	"math/rand"
	"time"
)

type user struct {
	id    int
	age   int
	saved bool
}

func main() {
	initialTime := time.Now()

	processes := make(chan int, 5)
	results := make(chan *user, 5)

	for i := 0; i < 5; i++ {
		go worker(processes, results, 1)
	}

	for i := 0; i < 5; i++ {
		processes <- i
	}
	close(processes)

	for i := 0; i < 5; i++ {
		fmt.Printf("User processed %+v \n", <-results)
	}

	fmt.Printf("Total time %v", time.Now().Sub(initialTime))
}

func worker(processes <-chan int, results chan<- *user, workerID int) {
	for process := range processes {
		user := request(process, workerID)
		calculate(user, workerID)
		save(user, workerID)
		results <- user
	}
}

func request(id, workerID int) *user {
	fmt.Printf("Init request - User: [%v] - Worker: [%v] \n", id, workerID)
	time.Sleep(time.Second * 2)
	fmt.Printf("Finish request - User: [%v] - Worker: [%v] \n", id, workerID)

	return &user{
		id: id,
	}
}

func calculate(user *user, workerID int) {
	fmt.Printf("Init calculate - User: [%v] - Worker: [%v] \n", user.id, workerID)
	time.Sleep(time.Millisecond * 100)
	fmt.Printf("Finish calculate - User: [%v] - Worker: [%v] \n", user.id, workerID)

	user.age = rand.Intn(100)
}

func save(user *user, workerID int) {
	fmt.Printf("Init save - User: [%v] - Worker: [%v] \n", user.id, workerID)
	time.Sleep(time.Second)
	fmt.Printf("Finish save - User: [%v] - Worker: [%v] \n", user.id, workerID)

	user.saved = true
}
