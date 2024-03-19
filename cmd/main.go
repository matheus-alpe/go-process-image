package main

import (
	processimage "process-image/pkg/process-image"
	"strings"
	"sync"
)

var paths = []string{
	"images/1.jpeg",
	"images/2.jpeg",
	"images/3.jpg",
}

func main() {
	var wg sync.WaitGroup

	wg.Add(len(paths))
	for _, path := range paths {
		go func(path string) {
			defer wg.Done()
			img := processimage.Read(path)
			resizedImg := processimage.Resize(img)
			processimage.Write(strings.Replace(path, "images/", "images/output/", 1), resizedImg)
		}(path)
	}

	wg.Wait()
	println("done")
}
