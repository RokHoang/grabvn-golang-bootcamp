package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	word  string
	count int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	pipe := make(chan map[string]int)
	finalCounter := make(map[string]int)
	go consumer(pipe, finalCounter)

	fmt.Print("> ")
	for scanner.Scan() {
		text := scanner.Text()
		go producer(text, pipe)

	}
}

func consumer(pipe chan map[string]int, finalCounter map[string]int) {
	for wordCount := range pipe {
		for word, count := range wordCount {
			finalCounter[word] += count
		}
		fmt.Println("Result:", finalCounter)
		fmt.Print("> ")
	}
}

func producer(text string, pipe chan map[string]int) {
	wordCount := countWord(text)
	fmt.Println(wordCount)
	sendWordCount(wordCount, pipe)
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
