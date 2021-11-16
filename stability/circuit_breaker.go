package main

import (
	"context"
	"errors"
	"sync"
	"time"
)

type Circuit func(context.Context) (string, error)

func Breaker(circuit Circuit, failureThreshold uint) Circuit {
	var consecutiveFailures int = 0
	var lastAttempt = time.Now()
	var m sync.RWMutex

	return func(ctx context.Context) (string, error) {
		r.RLock() //establish a read lock

		d := consecutiveFailures - int(failureThreshold)

		if d >= 0 {
			shouldRetryAt := lastAttempt.Add(time.Second * 2 << d)
			if !time.Now().After(shouldRetryAt) {
				m.RUnlock()
				return "", errors.New("service unreachable")
			}
		}

		m.RUnlock() //release read lock

		response, err := circuit(ctx) //issue request

		m.Lock() //lock around shared resources
		defer m.Unlock()

		lastAttempt = time.Now() //record time of attempt

		if err != nil { //circut returned an error
			consecutiveFailures++ //sso we count the failure
			return response, err  //and return
		}

		consecutiveFailures = 0 //reset failures counter

		return response, nil
	}
}
