package main

import (
	"fmt"
	"strings"
	"sync"
)

func closePipe(pipe chan map[string]int, wg *sync.WaitGroup) {
	wg.Wait()
	close(pipe)
}

func waitResult(wg *sync.WaitGroup) {
	wg.Wait()
}

func consumer(pipe chan map[string]int, wg *sync.WaitGroup) {
	finalCounter := make(map[string]int)

	for wordCount := range pipe {
		for word, count := range wordCount {
			finalCounter[word] += count
		}
		fmt.Println("+++Result:", finalCounter)
	}
	fmt.Println("++++++Final Result:", finalCounter)
	wg.Done()
}

func producer(text string, pipe chan map[string]int, wg *sync.WaitGroup) {
	wordCount := countWord(text)
	fmt.Println("Raw:", text)
	fmt.Println("Count:", wordCount)
	sendWordCount(wordCount, pipe)
	wg.Done()
}

func countWord(text string) map[string]int {
	words := strings.Fields(text)
	wordCountMap := make(map[string]int)

	for _, word := range words {
		wordCountMap[word]++
	}
	return wordCountMap
}

func sendWordCount(wordCountMap map[string]int, pipe chan map[string]int) {
	pipe <- wordCountMap
}
