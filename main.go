package main

func check(err error) {
	if err != nil && err != io.EOF {
		panic(err)
	}
}

var filename = flag.String("file", "/Users/AnimotoOverstreet/go/bin/web-Google.txt", "file to parse on")

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

	lines := make([]Sentence, sentenceCount)

	index := 0

	for scanner.Scan() {
		lines[index] = scanner.Text().split(' ')[1:]
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

	alreadyCompared := make(map[struct{sentence, sentence}]bool)
	similarPairsCount := 0

	for k, sentences := range lshBuckets {
		for i, sentence := range sentences {
			for j := i+1; j < len(sentences); j++ {
				if !alradyCompared[{sentence, sentences[j]}] {
					alreadyCompared[{sentence, sentences[j]}] = true
					alreadyCompared[{sentences[j], sentence}] = true
					if sentence.compareWithSameLength(sentences[j]) {
						similarPairsCount++
					}
				}
			}

			for _, otherSentences := range k.largerNeighbors() {
				for _, otherSentence := range otherSentences {
					if !alreadyCompared[{sentence, sentences[j]}] {
						alreadyCompared[{sentence, otherSentence}] = true
						alreadyCompared[{otherSentence, sentence}] = true
						if sentence.compareWithLonger(otherSentence) {
							similarPairsCount++
						}
					}
				}
			}
		}
	}

	finish = time.Since(start)
	fmt.Println("algorithm time:", finish)

	fmt.Println("number of similar pairs:", similarPairsCount)

}
