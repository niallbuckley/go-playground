package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

// Create a custom http client with timeout and connection pooling.
// this allows us to reuse TCP connections to the same host and port server. 
// Also to hold onto DNS resolution to same server
// Very useful when we are making a lot of requests to the same server.
var httpClient = &http.Client{
	// We define timeout in the context. 
    // Why define in context and not here? that is important to know!!
	// Timeout: 5 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     90 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:   3 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	},
}

type apiResult struct {
	Body []byte
	Err  error
}

func makeAPIRequest(ctx context.Context, url string, resultCh chan<- apiResult) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		resultCh <- apiResult{nil, fmt.Errorf("request creation error: %w", err)}
		return
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		resultCh <- apiResult{nil, fmt.Errorf("http request error: %w", err)}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resultCh <- apiResult{nil, fmt.Errorf("non-200 status: %d %s", resp.StatusCode, resp.Status)}
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		resultCh <- apiResult{nil, fmt.Errorf("reading body failed: %w", err)}
		return
	}

	resultCh <- apiResult{body, nil}
}

func main() {
    // Set to Max CPUs by default. (runtime.GOMAXPROCS(0)) 
    // runtime.GOMAXPROCS(1) // force single-threaded execution
    
    urls := []string{
        "https://jsonplaceholder.typicode.com/posts/1",
        "https://jsonplaceholder.typicode.com/posts/2",
        "https://jsonplaceholder.typicode.com/posts/3",
    }
    
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resultCh := make(chan apiResult, len(urls))
    for _, url := range urls {
        go makeAPIRequest(ctx, url, resultCh)
    }

    for _, _ = range urls {
        result := <-resultCh
        if result.Err != nil {
            fmt.Println("API error:", result.Err)
            return
        }

        var data map[string]interface{}
        if err := json.Unmarshal(result.Body, &data); err != nil {
            fmt.Println("JSON error:", err)
            return
        }

        fmt.Println("userId:", data["id"])
    }
}
