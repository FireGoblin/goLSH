package main

//import "fmt"

type Sentence []string

func (s Sentence) buckets() []BucketIndex {
	if s[0] == s[1] {
		return []BucketIndex{{s[0], 0, len(s) - 1}}
	}
	return []BucketIndex{{s[0], 0, len(s) - 1}, {s[1], 1, len(s) - 2}}
}

func (s Sentence) compareWithSameLength(target Sentence, bucketLocation int) bool {
	//note: this is to prevent double counting
	//thus when bucketLocation=1 it can return false event when sentences are similar
	if bucketLocation == 1 {
		if s[0] == target[0] {
			//these were already compared in a different bucket
			return false
		}
	}

	mismatchAvailable := true
	for i, _ := range s {
		if s[i] != target[i] {
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

func (s Sentence) compareWithLonger(target Sentence) bool {
	offset := 0
	for i, _ := range s {
		if s[i] != target[i+offset] {
			if offset == 1 {
				return false
			}
			offset++
			if s[i] != target[i+offset] {
				return false
			}
		}
	}
	// fmt.Println("similar pair:")
	// fmt.Println("   ", s)
	// fmt.Println("   ", target)
	return true
}
