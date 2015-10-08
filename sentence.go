package main

import "fmt"

type Sentence struct {
	sentence              []string
	frequencySortedPrefix [2]string
}

func (s *Sentence) sort(freq map[string]int) {
	var min string
	var secondMin string

	var minVal int
	var secondMinVal int

	if freq[s.sentence[0]] < freq[s.sentence[1]] {
		min = s.sentence[0]
		minVal = freq[s.sentence[0]]

		secondMin = s.sentence[1]
		secondMinVal = freq[s.sentence[1]]
	} else {
		min = s.sentence[1]
		minVal = freq[s.sentence[1]]

		secondMin = s.sentence[0]
		secondMinVal = freq[s.sentence[0]]
	}

	for _, v := range s.sentence[2:] {
		if freq[v] < minVal {
			secondMin = min
			secondMinVal = minVal

			min = v
			minVal = freq[v]
		} else if freq[v] < secondMinVal {
			secondMin = v
			secondMinVal = freq[v]
		}
	}

	s.frequencySortedPrefix = [2]string{min, secondMin}
	fmt.Println(s.frequencySortedPrefix)
}

func (s *Sentence) buckets() []BucketIndex {
	if s.frequencySortedPrefix[0] == s.frequencySortedPrefix[1] {
		fmt.Println(BucketIndex{s.frequencySortedPrefix[0], 0, len(s.sentence) - 1})
		return []BucketIndex{{s.frequencySortedPrefix[0], 0, len(s.sentence) - 1}}
	}
	fmt.Println(2)
	return []BucketIndex{{s.frequencySortedPrefix[0], 0, len(s.sentence) - 1}, {s.frequencySortedPrefix[1], 1, len(s.sentence) - 2}}
}

func (s *Sentence) compareWithSameLength(target Sentence, bucketLocation int) bool {
	//note: this is to prevent double counting
	//thus when bucketLocation=1 it can return false event when sentences are similar
	if bucketLocation == 1 {
		if s.frequencySortedPrefix[0] == target.frequencySortedPrefix[0] {
			//these were already compared in a different bucket
			return false
		}
	}

	mismatchAvailable := true
	for i, v := range s.sentence {
		if v != target.sentence[i] {
			if !mismatchAvailable {
				return false
			}
			mismatchAvailable = false
		}
	}
	// if !mismatchAvailable {
	// 	fmt.Println("similar pair:")
	// 	fmt.Println("   ", s)
	// 	fmt.Println("   ", target)
	// }
	return true
}

func (s *Sentence) compareWithLonger(target Sentence) bool {
	offset := 0
	for i, v := range s.sentence {
		if v != target.sentence[i+offset] {
			if offset == 1 {
				return false
			}
			offset++
			if v != target.sentence[i+offset] {
				return false
			}
		}
	}
	// fmt.Println("similar pair:")
	// fmt.Println("   ", s)
	// fmt.Println("   ", target)
	return true
}
