package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type Pair struct {
	word  string
	count int
}

func main() {
	readTerminal := flag.Bool("terminal", false, "read from terminal")
	flag.Parse()
	if *readTerminal {
		//for testing with terminal run the command "go run *.go -terminal"
		readFromTerminal()
	} else {
		//else "go run *.go"
		readFromDirectory()
	}
}

func listAllFilesInDirectory(root string) ([]string, error) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return nil, err
	}
	var filePaths []string
	for _, file := range files {
		filePaths = append(filePaths, file.Name())
	}
	return filePaths, nil
}

func readFromDirectory() {
	var wgProducer sync.WaitGroup
	var wgConsumer sync.WaitGroup
	pipe := make(chan map[string]int)

	go consumer(pipe, &wgConsumer)

	var root string
	fmt.Print("Please type the directory: ")
	fmt.Scanln(&root)
	filePaths, err := listAllFilesInDirectory(root)
	if err != nil {
		log.Fatal("ERROR:", err)
		return
	}
	fmt.Println("There are file paths: ", filePaths)

	for _, filePath := range filePaths {
		wgProducer.Add(1)
		text, err := ioutil.ReadFile(filePath)
		if err == nil {
			go producer(string(text), pipe, &wgProducer)
		} else {
			fmt.Println("Cannot read text from", filePath)
			wgProducer.Done()
		}
	}
	wgProducer.Wait()
	close(pipe)
	waitResult(&wgConsumer)
}

func readFromTerminal() {
	var wgProducer sync.WaitGroup
	var wgConsumer sync.WaitGroup
	pipe := make(chan map[string]int)

	go consumer(pipe, &wgConsumer)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		wgProducer.Add(1)
		go producer(text, pipe, &wgProducer)
	}
	wgProducer.Wait()
	close(pipe)
	waitResult(&wgConsumer)
}
