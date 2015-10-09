package main

import "strings"

type Sentence struct {
	sentence []string
	buckets  []BucketIndex
}

func (s *Sentence) createBuckets() {
	s.buckets = []BucketIndex{{strings.Join(s.sentence[0:4], " "), 0, len(s.sentence)}, {strings.Join(s.sentence[len(s.sentence)-4:len(s.sentence)], " "), 1, len(s.sentence)}}
}

func (s *Sentence) compareWithSameLength(target Sentence, bucketLocation int) bool {
	//note: this is to prevent double counting
	//thus when bucketLocation=1 it can return false event when sentences are similar
	if bucketLocation == 1 {
		if s.buckets[0] == target.buckets[0] {
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

func (s *Sentence) compareWithLonger(target Sentence, bucketLocation int) bool {
	// if bucketLocation == 1 {
	// 	if s.buckets[0] == target.buckets[0] {
	// 		//these were already compared in a different bucket
	// 		return false
	// 	}
	// }

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
