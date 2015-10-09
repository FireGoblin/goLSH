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
	"sync"
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
	//sentenceCount = 2000000

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

	concurrencyCount := runtime.GOMAXPROCS(0) * 8

	same, diff := gen(lshBuckets)

	responseChans := make([]<-chan int, concurrencyCount*2)

	for i := 0; i < concurrencyCount; i++ {
		responseChans[2*i] = sameWorker(same)
		responseChans[2*i+1] = diffWorker(diff)
	}

	for i := range merge(responseChans...) {
		similarPairsCount += i
	}

	finish = time.Since(start)
	fmt.Println("algorithm time:", finish)

	fmt.Println("number of similar pairs:", similarPairsCount)

}

func gen(lshBuckets map[BucketIndex][]*UniqueSentence) (<-chan Request, <-chan Request) {
	sameWorkChan := make(chan Request)
	diffWorkChan := make(chan Request)
	go func() {
		for k, sentences := range lshBuckets {
			for i, sentence := range sentences {
				for j := i + 1; j < len(sentences); j++ {
					sameWorkChan <- Request{sentence, sentences[j], k.location}
					//similarPairsCount += sentence.compareWithSameLength(*sentences[j], k.location)
				}

				for _, otherSentence := range lshBuckets[k.largerNeighbor()] {
					diffWorkChan <- Request{sentence, otherSentence, k.location}
					//similarPairsCount += sentence.compareWithLonger(*otherSentence, k.location)
				}
			}
		}
		close(sameWorkChan)
		close(diffWorkChan)
	}()
	return sameWorkChan, diffWorkChan
}

func sameWorker(in <-chan Request) <-chan int {
	out := make(chan int)
	go func() {
		for x := range in {
			out <- x.s1.compareWithSameLength(x.s2, x.location)
		}
		close(out)
	}()
	return out
}

func diffWorker(in <-chan Request) <-chan int {
	out := make(chan int)
	go func() {
		for x := range in {
			out <- x.s1.compareWithLonger(x.s2, x.location)
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
