package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
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
	sentenceCount = 2000000

	//lines := make([]Sentence, sentenceCount)

	lines := make(map[string]*UniqueSentence)

	index := 0

	for index < sentenceCount && scanner.Scan() {
		splitLine := strings.Split(scanner.Text(), " ")[1:]
		line := strings.Join(splitLine, " ")

		_, ok := lines[line]
		if !ok {
			lines[line] = &UniqueSentence{splitLine, 0}
		}
		lines[line].incr()

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

	lshBuckets := make(map[BucketIndex][]*UniqueSentence)

	similarPairsCount := 0

	for i, v := range lines {
		similarPairsCount += v.selfPairs()
		for _, b := range v.buckets() {
			lshBuckets[b] = append(lshBuckets[b], lines[i])
		}
	}

	finish = time.Since(start)
	fmt.Println("time to make buckets:", finish)
	fmt.Println("number of buckets:", len(lshBuckets))

	concurrencyCount := runtime.GOMAXPROCS(0)

	sameWorkChan := make(chan Request, 1000000)
	diffWorkChan := make(chan Request, 1000000)
	responseChan := make(chan int, 1000000)
	checks := 0

	for i := 0; i < concurrencyCount; i++ {
		go sameWorker(sameWorkChan, responseChan)
		go diffWorker(diffWorkChan, responseChan)
	}

	for k, sentences := range lshBuckets {
		for i, sentence := range sentences {
			for j := i + 1; j < len(sentences); j++ {
				checks++
				sameWorkChan <- Request{sentence, sentences[j], k.location}
				//similarPairsCount += sentence.compareWithSameLength(*sentences[j], k.location)
			}

			for _, otherSentence := range lshBuckets[k.largerNeighbor()] {
				checks++
				diffWorkChan <- Request{sentence, otherSentence, k.location}
				//similarPairsCount += sentence.compareWithLonger(*otherSentence, k.location)
			}
		}
	}

	responses := 0

	for i := range responseChan {
		similarPairsCount += i
		responses++
		if responses == checks {
			break
		}
	}

	finish = time.Since(start)
	fmt.Println("algorithm time:", finish)

	fmt.Println("number of similar pairs:", similarPairsCount)

}

func sameWorker(in chan Request, out chan int) {
	for {
		select {
		case x := <-in:
			out <- x.s1.compareWithSameLength(x.s2, x.location)
		}
	}
}

func diffWorker(in chan Request, out chan int) {
	for {
		select {
		case x := <-in:
			out <- x.s1.compareWithLonger(x.s2, x.location)
		}
	}
}
