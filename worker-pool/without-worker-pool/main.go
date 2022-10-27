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

	for i := 0; i < 5; i++ {
		fmt.Printf("User processed %+v \n", process(i))
	}

	finalTime := time.Now()

	fmt.Printf("Total time %v", finalTime.Sub(initialTime))
}

func process(id int) (user *user) {
	user = request(id)
	calculate(user)
	save(user)

	return user
}

func request(id int) *user {
	fmt.Printf("Init request - User: [%v] \n", id)
	time.Sleep(time.Second * 2)
	fmt.Printf("Finish request - User: [%v] \n", id)

	return &user{
		id: id,
	}
}

func calculate(user *user) {
	fmt.Printf("Init calculate - User: [%v] \n", user.id)
	time.Sleep(time.Millisecond * 100)
	fmt.Printf("Finish calculate - User: [%v] \n", user.id)

	user.age = rand.Intn(100)
}

func save(user *user) {
	fmt.Printf("Init save - User: [%v] \n", user.id)
	time.Sleep(time.Second)
	fmt.Printf("Finish save - User: [%v] \n", user.id)

	user.saved = true
}
