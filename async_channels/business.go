// business.go
package main

import (
	"errors"
	"fmt"
	"time"
)

// StatusUpdate represents the status of an IP.
type StatusUpdate struct {
	IP    string
	State string
}

// BusinessArgs holds the arguments for the business logic.
type BusinessArgs struct {
	IPs []string
}

// BusinessResult holds the results from the business logic.
type BusinessResult struct {
	Status chan []StatusUpdate
	Error  chan error
}

// fetchEtagCtx simulates a business logic context.
type fetchEtagCtx struct{}

// Spawn starts the business logic asynchronously.
func (f *fetchEtagCtx) Spawn(args *BusinessArgs) (*BusinessResult, error) {
	if args == nil || len(args.IPs) == 0 {
		return nil, errors.New("invalid arguments")
	}

	result := &BusinessResult{
		Status: make(chan []StatusUpdate),
		Error:  make(chan error),
	}

	go func() {
		defer close(result.Status)
		defer close(result.Error)

		for _, ip := range args.IPs {
			time.Sleep(1 * time.Second) // Simulate work
			if ip == "192.168.1.2" {
				result.Error <- fmt.Errorf("did not work for ip %s", ip)
				continue
			}
			result.Status <- []StatusUpdate{{IP: ip, State: "Processed"}}
		}
	}()

	return result, nil
}

