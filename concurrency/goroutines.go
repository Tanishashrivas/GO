package main

import (
	"fmt"
	"sync"
	"time"
)

// GoRoutines: A GoRoutine is a lightweight thread managed by the Go runtime. It allows us to run functions concurrently without creating separate OS threads.
// It’s lightweight compared to normal threads, doesn’t require manual thread management and Go runtime handles scheduling automatically.
// GoRoutine is like a mini worker inside your program that runs in the background while your main program continues.

// WaitGroup: sync.WaitGroup is a synchronization mechanism used to wait for a group of goroutines to finish executing before proceeding. It helps manage concurrent operations by ensuring all goroutines complete before the main function exits.

var wg = sync.WaitGroup{} //explicitely initialized
var mu sync.Mutex         // auto-initialized
var db = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

// Mutex: (Mutual Exclusion Lock) in Golang is used to prevent race conditions when multiple goroutines try to access or modify a shared resource concurrently. The sync.Mutex type provides methods to lock and unlock critical sections of code.
// Lock() - Acquires the lock (blocking other goroutines).
// Unlock() - Releases the lock (allowing other goroutines to proceed).
// Always use defer with unlock or else it causes a deadlock! Using defer ensures that Unlock() always executes, even if a function returns early.

func main() {
	t0 := time.Now()
	for i := 0; i < len(db); i++ {
		wg.Add(1) // It is like a counter. Add 1 task to the WaitGroup
		go dbCall(i)
	}

	// wg.Wait() blocks the main() function until all GoRoutines finish.
	wg.Wait() // ⏳ Wait for all GoRoutines to finish
	fmt.Println("Result: ", results)
	fmt.Println("Time taken: ", time.Since(t0))
}

func dbCall(i int) {
	var delay = 2000
	// var delay = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("Result of dbcall is: ", db[i])
	results = append(results, db[i]) // multiple GoRoutines try to modify shared data at the same time. Might cause error. To avoid this, we use mutex
	defer wg.Done()                  // Mark this GoRoutine as finished (decrement the counter by 1)
}
