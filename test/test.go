package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func doSomething(id int, wg *sync.WaitGroup) {

	//defer wg.Done()
	log.Printf("before do job:(%d) \n", id)
	time.Sleep(3 * time.Second)
	log.Printf("after do job:(%d) \n", id)
	// fd
	// 1
	// 2
}

func main() {
	//var wg sync.WaitGroup
	//wg.Add(3)
	//
	//go doSomething(1, &wg)
	//go doSomething(2, &wg)
	//go doSomething(3, &wg)
	//
	//wg.Wait()
	//
	// wg.Done()
	//log.Printf("finish all jobs\n")




	defer fmt.Println("in main")
	if err := recover(); err != nil {
		fmt.Println(err)
	}

	panic("unknown err")
}
	//wg.Done()
	//log.Printf("finish all jobs\n")
