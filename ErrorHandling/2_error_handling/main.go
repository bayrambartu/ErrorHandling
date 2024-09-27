// What is log in short?
/* package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Program başladı.")

	// Bir hata durumu simüle edelim
	err := doSomething()
	if err != nil {
		log.Println("Hata: ", err)
	}

	log.Println("Program bitti.")
}

func doSomething() error {
	return fmt.Errorf("beklenmeyen hata olustu")
}
*/
// AI -> What  is the errorhandling ??

package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// Custom error types for different scenarios
var (
	ErrCritical = errors.New("critical error")
	ErrNonFatal = errors.New("non-fatal error")
)

// Simulates a function that can fail
func performOperation() error {
	// Simulating a failure
	return ErrNonFatal
}

// Handles errors based on their type and severity
func handleError(err error) {
	if errors.Is(err, ErrCritical) {
		log.Fatalf("Critical error encountered: %v. Terminating program.", err)
	} else if errors.Is(err, ErrNonFatal) {
		log.Printf("Non-fatal error occurred: %v. Continuing operation.", err)
	} else {
		log.Printf("Unknown error: %v", err)
	}
}

func main() {
	retries := 3
	var err error

	// Retry mechanism with error handling
	for i := 0; i < retries; i++ {
		if err = performOperation(); err == nil {
			fmt.Println("Operation succeeded.")
			break
		}
		handleError(err)
		time.Sleep(2 * time.Second) // wait before retry
	}

	// Final error check after retries
	if err != nil {
		log.Fatalf("Failed after %d attempts: %v", retries, err)
	}
}
