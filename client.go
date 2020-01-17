package main

import (
	"log"
	"net/http"
	"sync"
)

func main() {
	errno := 0
	log.Println("begin:")
	w := &sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		w.Add(1)
		go func() {
			for t := 0; t < 1000; t++ {
				func() {
					defer func() {
						if r := recover(); r != nil {
							log.Fatal("recover error", r)
						}
					}()
					r, err := http.Get("http://49.234.232.77")
					if err != nil {
						log.Printf("Error: %s", err.Error())
						return
					}
					r.Body.Close()
					if r.StatusCode != 200 {
						errno++
						log.Printf("ResponseErr:%d\n", errno)
					}
				}()

			}
			w.Done()
		}()
	}
	w.Wait()

}
