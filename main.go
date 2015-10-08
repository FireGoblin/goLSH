package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(err error) {
	if err != nil && err != io.EOF {
		panic(err)
	}
}

var filename = flag.String("file", "/Users/AnimotoOverstreet/go/bin/sentences.txt", "file to parse on")

func main() {
	flag.Parse()
	fmt.Println(*filename)

	start := time.Now()

	file, err := os.Open(*filename)
	check(err)

	scanner := bufio.NewScanner(bufio.NewReader(file))

	scanner.Scan()
	sentenceCount, err := strconv.Atoi(scanner.Text())
	check(err)
	sentenceCount = 1000000

	lines := make([]Sentence, sentenceCount)

	index := 0

	for index < sentenceCount && scanner.Scan() {
		lines[index] = strings.Split(scanner.Text(), " ")[1:]
		index++
	}

	if scanner.Err() != nil {
		panic("reading input exited with non-EOF error")
	}

	file.Close()

	finish := time.Since(start)
	//reading file done
	fmt.Println("time to read file:", finish)

	start = time.Now()

	lshBuckets := Buckets(make(map[BucketIndex][]Sentence))

	for _, v := range lines {
		for _, b := range v.buckets() {
			lshBuckets[b] = append(lshBuckets[b], v)
		}
	}

	finish = time.Since(start)
	fmt.Println("time to make buckets:", finish)
	fmt.Println("number of buckets:", len(lshBuckets))

	similarPairsCount := 0
	checks := 0
	buckets := 0

	for k, sentences := range lshBuckets {
		for i, sentence := range sentences {
			for j := i + 1; j < len(sentences); j++ {
				checks++
				if sentence.compareWithSameLength(sentences[j], k.location) {
					similarPairsCount++
				}
			}

			for _, otherSentences := range k.largerNeighbors() {
				for _, otherSentence := range lshBuckets[otherSentences] {
					checks++
					if sentence.compareWithLonger(otherSentence) {
						similarPairsCount++
					}
				}
			}
		}
		buckets++
		if buckets%10000 == 0 {
			fmt.Println("buckets:", buckets)
			fmt.Println("checks:", checks)
			fmt.Println("similarPairsCount:", similarPairsCount)
		}
	}

	finish = time.Since(start)
	fmt.Println("algorithm time:", finish)

	fmt.Println("number of similar pairs:", similarPairsCount)

}
