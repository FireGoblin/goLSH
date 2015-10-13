package main

import "strings"

type uniqueSentence struct {
	sentence []string
	count    int
}

func (s *uniqueSentence) incr() {
	s.count++
}

func (s uniqueSentence) selfPairs() int {
	return s.count * (s.count - 1) / 2
}

func (s uniqueSentence) buckets() [2]bucketIndex {
	return [2]bucketIndex{{strings.Join(s.sentence[0:4], " "), 0, len(s.sentence)}, {strings.Join(s.sentence[len(s.sentence)-4:len(s.sentence)], " "), 1, len(s.sentence)}}
}

func (s uniqueSentence) compareWithSameLength(target uniqueSentence, bucketLocation int) int {
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

func (s uniqueSentence) compareWithLonger(target uniqueSentence, bucketLocation int) int {
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
