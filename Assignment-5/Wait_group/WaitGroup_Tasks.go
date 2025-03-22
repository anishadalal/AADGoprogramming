package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadFile(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Downloading file %d...\n", id)
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("File %d downloaded.\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go downloadFile(i, &wg)
	}

	wg.Wait()
	fmt.Println("All files downloaded.")
}
