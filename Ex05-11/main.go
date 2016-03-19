// Exercise 5.11: sorting courses topologically and detecting loops.
package main

import (
	"fmt"
	"log"
	"sort"
)

type CourseID int
type CourseIDs []CourseID

func (ids CourseIDs) Len() int           { return len(ids) }
func (ids CourseIDs) Swap(i, j int)      { ids[i], ids[j] = ids[j], ids[i] }
func (ids CourseIDs) Less(i, j int) bool { return ids[i] < ids[j] }

// CourseID -> Name of course
var courseNames = map[CourseID]string{
	0:  "algorithms",
	1:  "calculus",
	2:  "compilers",
	3:  "data structures",
	4:  "databases",
	5:  "discrete math",
	6:  "formal languages",
	7:  "networks",
	8:  "operating systems",
	9:  "programming languages",
	10: "linear algebra",
	11: "computer organization",
	12: "intro to programming"}

// CourseID -> IDs of prerequisite courses.
var prereqs = map[CourseID]CourseIDs{
	0: {3},
	1: {10},
	2: {3, 6, 11},
	3: {5},
	4: {3},
	5: {12},
	6: {5},
	7: {8},
	8: {3, 11},
	9: {3, 11}}

func main() {
	order, err := topologicalSort(prereqs)
	if err != nil {
		// Error when processing a topological sort.
		log.Fatalf("%s", err)
	}

	for i, courseID := range order {
		fmt.Printf("%d:\t%s\n", i+1, courseNames[courseID])
	}
}

// Sort m topologically, returning an error if an error occurs
// (for example, if there is a cycle in m)
func topologicalSort(m map[CourseID]CourseIDs) (CourseIDs, error) {
	var order CourseIDs
	seen := make(map[CourseID]bool)
	seenInCurrentPath := make(map[CourseID]bool)

	// visitAll performs a depth-first search on the supplied courseIDs,
	// using the graph defined by preqs. It appends the result to order.
	var visitAll func(courseIDs CourseIDs) error
	visitAll = func(courseIDs CourseIDs) error {
		for _, courseID := range courseIDs {
			// If this CourseID is already in the current path, then we have a loop.
			if seenInCurrentPath[courseID] {
				return fmt.Errorf("%s", courseNames[courseID])
			}
			if !seen[courseID] {
				seen[courseID] = true
				seenInCurrentPath[courseID] = true
				if err := visitAll(m[courseID]); err != nil {
					// Loop found. Propogate courseID to report the chain.
					return fmt.Errorf("%s: %s", courseNames[courseID], err)
				}

				order = append(order, courseID)
				seenInCurrentPath[courseID] = false
			}
		}
		return nil
	}

	// Sort the keys to ensure a deterministic ordering.
	var keys CourseIDs
	for key := range m {
		keys = append(keys, key)
	}
	sort.Sort(keys)

	// Actually traverse the keys, and return the order.
	if err := visitAll(keys); err != nil {
		// Loop found.
		return nil, fmt.Errorf("Loop found: %s", err)
	}

	return order, nil
}
