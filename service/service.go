package service

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sync"
)

var exitChan = make(chan struct{}, 1)

func FetchWebpages() {
	reader := bufio.NewReader(os.Stdin)
	wg := sync.WaitGroup{}
	for {
		select {
		case <-exitChan:
			return
		default:
			arg, err := readInput(reader)
			if err != nil {
				if !errors.Is(err, EmptyInputError) && !errors.Is(err, ExitError) {
					fmt.Println("Error reading input from command line: ", err)
				}
				continue
			}
			wg.Add(len(arg.urls))
			for _, url := range arg.urls {
				go fetch(&wg, arg.downloadPath, arg.strategy, url)
			}
			wg.Wait()
			fmt.Println("===============DONE======================")
		}
	}
}
