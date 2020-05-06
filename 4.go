package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type resp struct {
	k string
	v string
}

func main() {
	res,_ := fetchData()
	log.Print(res)
}

func rpcwork(k, v string) resp {
	// do some rpc work
	return resp{
		k,v,
	}
}

func fetchData() (map[string]string, error) {
	var result = map[string]string{} // result is k -> v
	var keys = []string{"a", "b", "c"}
	var keys1 = []string{"d", "e", "f"}
	var wg sync.WaitGroup
	var m sync.Mutex
	fmt.Println(len(keys))
	for i := 0; i < len(keys)+100000; i++ {

		wg.Add(1)

		go func(i int) {
			m.Lock()
			defer m.Unlock()
			defer wg.Done()

			// do some rpc
			resp := rpcwork(keys[i%3], keys1[i%3])

			result[resp.k] = resp.v
		}(i)
	}

	waitTimeout(&wg, time.Millisecond)
	return result, nil
}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}