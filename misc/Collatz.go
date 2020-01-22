package main

func main() {
	print(collatzCallCounter(3))
	// We can process in batches, by running 1,000 routines simultaneously.
	//for set := 0; set < 1000000; set++ {
	//	for routine := 0; routine < 1000; routine++ {
	//		go collatzCallCounter(routine * set)
	//	}
	//	// Need to determine a way to wait for all 1000 goroutines to finish before starting a new batch.
	//}
}

var results = make(map[int]int)

func collatzCallCounter(target int) int {
	originalTarget := target
	count := 1
	for target != 1 {
		if results[target] != 0 {
			count += results[target]
			break
		} else {
			if target % 2 == 0 {
				count += 1
				target = target / 2
			} else {
				count += 1
				target = (target * 3) + 1
			}
		}
	}
	results[originalTarget] = count
	return count
}

/**
def collatzCallCounter(self, target):
        count = 1
        while target != 1:
            if target in self.callCounts:
                count += self.callCounts[target]
                break
            else:
                if target % 2 == 0:  # Even
                    count += 1
                    target = target // 2
                else:
                    count += 1
                    target = (target * 3) + 1
            continue
        return count
 */