package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			res, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(i))

			if err != nil {
				log.Fatalln(err)
				return
			}

			defer func(Body io.ReadCloser) {
				err = Body.Close()
				if err != nil {
					log.Fatalln(err)
					return
				}
			}(res.Body)

			body, err := io.ReadAll(res.Body)
			if err != nil {
				log.Fatalln(err)
				return
			}

			sb := string(body)

			ch <- sb
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for b := range ch {
		fmt.Println(b)
	}
}
