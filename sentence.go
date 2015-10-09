package main

import "strings"

type UniqueSentence struct {
	sentence []string
	count    int
}

func (s *UniqueSentence) incr() {
	s.count++
}

func (s UniqueSentence) selfPairs() int {
	return s.count * (s.count - 1) / 2
}

func (s UniqueSentence) buckets() [2]BucketIndex {
	return [2]BucketIndex{{strings.Join(s.sentence[0:4], " "), 0, len(s.sentence)}, {strings.Join(s.sentence[len(s.sentence)-4:len(s.sentence)], " "), 1, len(s.sentence)}}
}

func (s UniqueSentence) compareWithSameLength(target UniqueSentence, bucketLocation int) int {
	mismatchAvailable := true

	for i, v := range s.sentence {
		//duplicate check condition
		if i == 4 && bucketLocation == 1 && mismatchAvailable {
			return 0
		}
		if v != target.sentence[i] {
			if !mismatchAvailable {
				return 0
			}
			mismatchAvailable = false
		}
	}
	// if !mismatchAvailable {
	// 	fmt.Println("similar pair:")
	// 	fmt.Println("   ", s)
	// 	fmt.Println("   ", target)
	// }
	return s.count * target.count
}

func (s UniqueSentence) compareWithLonger(target UniqueSentence, bucketLocation int) int {
	offset := 0
	for i, v := range s.sentence {
		//duplicate check condition
		if i == 4 && bucketLocation == 1 && offset == 0 {
			return 0
		}
		if v != target.sentence[i+offset] {
			if offset == 1 {
				return 0
			}
			offset++
			if v != target.sentence[i+offset] {
				return 0
			}
		}
	}
	// fmt.Println("similar pair:")
	// fmt.Println("   ", s)
	// fmt.Println("   ", target)
	return s.count * target.count
}
