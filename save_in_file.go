package nix

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func SaveInFile(i int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/" + strconv.Itoa(i))

	if err != nil {
		log.Fatalln(err)
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
	}

	sb := string(body)

	message := []byte(sb)
	name := "storage/posts/" + strconv.Itoa(i) + ".txt"

	err = os.WriteFile(name, message, 0644)
	if err != nil {
		log.Fatalln(err)
	}

	ch <- sb
}
