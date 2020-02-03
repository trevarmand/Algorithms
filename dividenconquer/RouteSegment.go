package dividenconquer

import (
	"fmt"
)

/**
With a one directional graph of n nodes, generate a set of edges in O(log n) that connects each node such that you can
travel from any node i to any larger node j by traversing at most 2 edges.
 */

func main() {
	for _, segment := range getSegments(6) {
		fmt.Print(segment, "\n")
	}
}

type Segment struct {
	firstStop int
	secondStop int
}

func getSegments(stops int) []Segment {
	return splitAssignment(1, stops)
}

func splitAssignment(min, max int) []Segment {
	segmentSet := make([]Segment, 0) // Builds a dynamically sized slice
	spread := max - min
	if spread == 0 {
		return segmentSet
	}
	if spread == 1 {
		return []Segment{Segment{min, max}}
	} else {
		midpoint := (max + min) / 2

		for start := min; start < midpoint; start++ {
			segmentSet = append(segmentSet, Segment{start, midpoint})
		}
		for end := midpoint + 1; end <= max; end++ {
			segmentSet = append(segmentSet, Segment{midpoint, end})
		}

		leftHalf := splitAssignment(min, midpoint-1)
		rightHalf := splitAssignment(midpoint + 1, max)

		return append(segmentSet, append(rightHalf, leftHalf...)...)
	}
}
