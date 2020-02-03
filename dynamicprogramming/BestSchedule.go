package main

import "fmt"

/**
There are two types of events, A and B. In general, I prefer to attend event B.
Every weekend there is an event A and an event B. I'd like to go to an event every weekend,
but unfortunately if I chose to attend event B, then I have to take the next weekend off.

Determine the best schedule given a set of n weekends
 */

//var A = []int {3,6,5,4}
//var B = []int {7, 12, 8, 6}

var Jazz = []int {3,2,4,5,4,2,3,4}
var Metal = []int {11,12,13,16,12,15,19,21}

var results = make(map[int]int)

// Determine the value of the optimal schedule for range i...n
func bestValue(i, n int) int {
	for j := n; j >= i; j-- {
		if j == n {
			results[j] =  max(Jazz[j], Metal[j])
		}
		if j == (n - 1) {
			results[j] = max(Metal[j], Jazz[j] + results[j + 1])
		} else {
			results[j] = max(Metal[j] + results[j + 2], Jazz[j] + results[j + 1])
		}
	}
	return results[i]
}


func bestSchedule(i, n int) map[int]string {
	var concerts = make(map[int]string)
	for j := n; j >= i; j-- {
		if j == n {
			results[j] =  max(Jazz[j], Metal[j])
			concerts[j] = compare(Jazz[j], Metal[j])
		}
		if j == (n - 1) {
			results[j] = max(Metal[j], Jazz[j] + results[j + 1])
			concerts[j] = compare(Jazz[j] + results[j + 1], Metal[j])
		} else {
			results[j] = max(Metal[j] + results[j + 2], Jazz[j] + results[j + 1])
			if compare( Jazz[j] + results[j + 1], Metal[j] + results[j + 2]) == "jazz" {
				concerts[j] = "jazz"
			} else {
				concerts[j+1] = "rest"
				concerts[j] = "metal"
			}
		}
	}
	return concerts
}


func compare(jazz, metal int) string {
	if jazz >= metal {
		return "jazz"
	}
		return "metal"
}

func main() {
	print(bestValue(0, 6))
	bestScheduleMap := bestSchedule(0, 6)
	for i := 0; i < len(bestScheduleMap); i++ {
		fmt.Println("Took:", bestScheduleMap[i], "from options:  Jazz: ", Jazz[i], " Metal: ", Metal[i])
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}