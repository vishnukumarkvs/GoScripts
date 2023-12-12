package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/joho/godotenv"
)

func BenchmarkGetBooks(b *testing.B) {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal(err)
    }

	app := setupApp()

	// Create a new http request with the desired method, path, and optional body

	// Run the benchmark
	for i := 0; i < b.N; i++ {
     	req, _ := http.NewRequest("GET", "/api/books", nil)
    	resp, err := app.Test(req, -1) // The -1 here disables request timeout
		if err != nil {
			b.Fatalf("Error on test request: %v", err)
		}

		// Read the response body. Working as expected
        // body, err := io.ReadAll(resp.Body)
        // if err != nil {
        //     b.Fatalf("Error reading response body: %v", err)
        // }

		// fmt.Printf("Status Code: %d, Body: %s\n", resp.StatusCode, string(body))


		// Check if the response status code is 200 OK
        if resp.StatusCode != http.StatusOK {
            b.Fatalf("Expected status code 200, got %d", resp.StatusCode)
        }
	}
}


// Test1 Results
/* Running tool: C:\Program Files\Go\bin\go.exe test -benchmem -run=^$ -bench ^BenchmarkGetBooks$ go-postgres-fiber

goos: windows
goarch: amd64
pkg: go-postgres-fiber
cpu: AMD Ryzen 5 4600H with Radeon Graphics         
BenchmarkGetBooks-12    	    1638	    744336 ns/op	   11415 B/op	     114 allocs/op
PASS
ok  	go-postgres-fiber	2.703s
*/

// Insights
/*
1. BenchmarkGetBooks-12  = using 12 threads
2. 1638                  = executed 1638 times
3. 744336 ns/op          = took 744336 nanoseconds / 0.7 ms for each operation
4. 11415 B/op            = memory utilized
5. 114 allocs/op         = 114 seperate allocations of memory made per each operation
6. 2.730s                = took 2 seconds
*/