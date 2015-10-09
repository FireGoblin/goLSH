package main

import "strings"

type Sentence struct {
	sentence []string
}

func (s Sentence) buckets() [2]BucketIndex {
	return [2]BucketIndex{{strings.Join(s.sentence[0:4], " "), 0, len(s.sentence)}, {strings.Join(s.sentence[len(s.sentence)-4:len(s.sentence)], " "), 1, len(s.sentence)}}
}

func (s Sentence) compareWithSameLength(target Sentence, bucketLocation int) bool {
	mismatchAvailable := true

	for i, v := range s.sentence {
		//duplicate check condition
		if i == 4 && bucketLocation == 1 && mismatchAvailable {
			return false
		}
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

func (s Sentence) compareWithLonger(target Sentence, bucketLocation int) bool {
	offset := 0
	for i, v := range s.sentence {
		//duplicate check condition
		if i == 4 && bucketLocation == 1 && offset == 0 {
			return false
		}
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
