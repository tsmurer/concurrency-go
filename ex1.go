package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchingExample() {
	respch := make(chan string, 1024)
	now := time.Now()

	wg := &sync.WaitGroup{}

	go fetchUserData(respch, wg)
	wg.Add(1)
	go fetchUserRecommendations(respch, wg)
	wg.Add(1)
	go fetchUserLikes(respch, wg)
	wg.Add(1)

	wg.Wait()

	close(respch)

	for resp := range respch {
		fmt.Println(resp)
	}
	fmt.Println(time.Since(now))
}

func fetchUserData(respch chan string, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 80)
	respch <- "userData"
	wg.Done()
}

func fetchUserRecommendations(respch chan string, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 120)
	respch <- "userRecommendations"
	wg.Done()
}

func fetchUserLikes(respch chan string, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 50)
	respch <- "userLikes"
	wg.Done()
}
