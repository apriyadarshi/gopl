package ch5

import (
	"sort"
)

/*var prereqs = map[string][]string{
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
}*/

//prereqs mas CS courses to their prerequisites
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var visitAll func(items []string) //Declared to allow its recursive call
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys) //sortng isn't necessary
	visitAll(keys)
	return order
}

func topoSort2(m map[string][]string) []string {

	//Courses with no prereqs
	noPrereq := make(map[string]bool)

	//Map from prereqs to the courses dependent on it
	revMap := make(map[string]map[string]bool)

	levels := make(map[string]int)
	//find edges that don't direct to any others
	for k, v := range m {
		for _, prereq := range v {
			// If course not in prereqs keys, it doesn't have any prereq.
			if _, ok := m[prereq]; !ok {
				noPrereq[prereq] = true
				levels[prereq] = 0
			}
			if _, ok := revMap[prereq]; ok {
				revMap[prereq][k] = true
			} else {
				revMap[prereq] = make(map[string]bool)
				revMap[prereq][k] = true
			}
		}
	}

	var traverse func(prereq string)
	traverse = func(prereq string) {
		if courses, ok := revMap[prereq]; ok {
			for course := range courses {

				//If level already set assign the greater level
				newlevel := levels[prereq] + 1
				if oldlevel, ok := levels[course]; ok && newlevel < oldlevel {
					levels[course] = oldlevel
				} else {
					levels[course] = newlevel
				}
				traverse(course)
			}
		}
	}

	//First traverse from zero level to top
	for k := range noPrereq {
		traverse(k)
	}
	//Now traverse secondary paths
	for k := range revMap {
		if _, ok := noPrereq[k]; !ok {
			traverse(k)
		}
	}

	//Levels is populated now
	maxLevel := 0
	levelGroup := make(map[int][]string)
	for k, v := range levels {
		if v > maxLevel {
			maxLevel = v
		}
		levelGroup[v] = append(levelGroup[v], k)
	}

	result := []string{}
	for i := 0; i < maxLevel+1; i++ {
		result = append(result, levelGroup[i]...)
	}
	return result

}

func topoSort3(m map[string][]string) ([]string, bool) {

	var cycle bool

	//Courses with no prereqs
	noPrereq := make(map[string]bool)

	//Map from prereqs to the courses dependent on it
	revMap := make(map[string]map[string]bool)

	levels := make(map[string]int)
	//find edges that don't direct to any others
	for k, v := range m {
		for _, prereq := range v {
			// If course not in prereqs keys, it doesn't have any prereq.
			if _, ok := m[prereq]; !ok {
				noPrereq[prereq] = true
				levels[prereq] = 0
			}
			if _, ok := revMap[prereq]; ok {
				revMap[prereq][k] = true
			} else {
				revMap[prereq] = make(map[string]bool)
				revMap[prereq][k] = true
			}
		}
	}

	seen := make(map[string]bool)
	var traverse func(prereq string)
	traverse = func(prereq string) {
		if courses, ok := revMap[prereq]; ok {
			for course := range courses {
				//If level already set assign the greater level
				newlevel := levels[prereq] + 1
				if oldlevel, ok := levels[course]; ok && newlevel < oldlevel {
					levels[course] = oldlevel
				} else {
					levels[course] = newlevel
				}
				//check if same point is reached in this path
				if _, ok := seen[prereq]; ok {
					cycle = true
				} else {
					seen[prereq] = true
					traverse(course)
				}
			}
		}
	}

	//First traverse from zero level to top
	for k := range noPrereq {
		seen = make(map[string]bool)
		traverse(k)
	}
	//Now traverse secondary paths
	for k := range revMap {
		if _, ok := noPrereq[k]; !ok {
			seen = make(map[string]bool)
			traverse(k)
		}
	}

	//Levels is populated now
	maxLevel := 0
	levelGroup := make(map[int][]string)
	for k, v := range levels {
		if v > maxLevel {
			maxLevel = v
		}
		levelGroup[v] = append(levelGroup[v], k)
	}

	result := []string{}
	for i := 0; i < maxLevel+1; i++ {
		result = append(result, levelGroup[i]...)
	}
	return result, cycle

}
