package main

import (
	"fmt"
	"image"
	"strings"
	"sync"
	"time"

	processimage "process-image/internal/process-image"
)

type Job struct {
	InputPath  string
	OutputPath string
	Image      image.Image
}

type ReadJobChan <-chan Job

type ImagePaths []string

func loadImage(paths ImagePaths) ReadJobChan {
	out := make(chan Job)

	go func() {
		var wg sync.WaitGroup

		wg.Add(len(paths))
		for _, path := range paths {
			go func(path string) {
				defer wg.Done()
				job := Job{
					InputPath:  path,
					OutputPath: strings.Replace(path, "images/", "images/output/", 1),
					Image:      processimage.Read(path),
				}
				out <- job
			}(path)
		}
		wg.Wait()

		close(out)
	}()

	return out
}

func resizeImage(input ReadJobChan) ReadJobChan {
	out := make(chan Job)

	go func() {
		for job := range input {
			job.Image = processimage.Resize(job.Image, 200, 200)
			out <- job
		}
		close(out)
	}()

	return out
}

func grayscaleImage(input ReadJobChan) ReadJobChan {
	out := make(chan Job)

	go func() {
		for job := range input {
			job.Image = processimage.Grayscale(job.Image)
			out <- job
		}
		close(out)
	}()

	return out
}

func saveImage(input ReadJobChan) <-chan bool {
	out := make(chan bool)

	go func() {
		for job := range input {
			processimage.Write(job.OutputPath, job.Image)
			out <- true
		}
		close(out)
	}()

	return out
}

func main() {
	paths := ImagePaths{
		"images/1.jpeg",
		"images/2.jpeg",
		"images/3.jpg",
	}

	now := time.Now()

	c1 := loadImage(paths)
	c2 := resizeImage(c1)
	c3 := grayscaleImage(c2)

	for success := range saveImage(c3) {
		if success {
			fmt.Println("Success")
		} else {

			fmt.Println("Failed")
		}
	}

	fmt.Println("done", time.Since(now))
}
