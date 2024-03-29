package main

import (
	"fmt"
	"strings"
	"sync"
)

type Pair struct {
	counter int
	safer   sync.Mutex
}

func closePipe(pipe chan map[string]int, wg *sync.WaitGroup) {
	wg.Wait()
	close(pipe)
}

func waitResult(wg *sync.WaitGroup) {
	wg.Add(1)
	wg.Wait()
}

func consumer(pipe chan map[string]int, wg *sync.WaitGroup) {
	finalCounter := make(map[string]int)

	for wordCount := range pipe {
		for word, count := range wordCount {
			finalCounter[word] += count
		}
		fmt.Println("Result:", finalCounter)
	}
	fmt.Println("Final Result:", finalCounter)
	wg.Done()
}

func producer(text string, pipe chan map[string]int, wg *sync.WaitGroup) {
	wordCount := countWord(text)
	//fmt.Println("Raw:", text)
	//fmt.Println("Count:", wordCount)
	pipe <- wordCount
	wg.Done()
}

type Producer struct {
	action func(string, chan map[string]int, *sync.WaitGroup)
	text   string
	pipe   chan map[string]int
	wg     *sync.WaitGroup
}

func (p Producer) Do() {
	p.action(p.text, p.pipe, p.wg)
}

func countWord(text string) map[string]int {
	words := strings.Fields(text)
	wordCountMap := make(map[string]int)

	for _, word := range words {

		wordCountMap[word]++

	}
	return wordCountMap
}
