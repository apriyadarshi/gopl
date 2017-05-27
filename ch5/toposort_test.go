package ch5

import (
	//"fmt"
	"testing"
)

func TestTopoSort(t *testing.T) {
	sorted := topoSort(prereqs)
	/*for i, course := range sorted {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}*/
	pos := make(map[string]int)
	for i, v := range sorted {
		pos[v] = i
	}

	for course, precourses := range prereqs {
		for _, precourse := range precourses {
			if pos[precourse] > pos[course] {
				t.Errorf("Incorrect sequence: %s should come after %s ", course, precourse)
			}
		}
	}
}

/*func TestTopoSort2(t *testing.T) {
	sorted := topoSort2(prereqs)
	pos := make(map[string]int)
	for i, v := range sorted {
		pos[v] = i
	}

	for course, precourses := range prereqs {
		for _, precourse := range precourses {
			if pos[precourse] > pos[course] {
				t.Errorf("Incorrect sequence: %s should come after %s ", course, precourse)
			}
		}
	}
}*/

func TestTopoSort3(t *testing.T) {
	sorted, _ := topoSort3(prereqs)

	/*for i, course := range sorted {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}*/
	pos := make(map[string]int)
	for i, v := range sorted {
		pos[v] = i
	}

	for course, precourses := range prereqs {
		for _, precourse := range precourses {
			if pos[precourse] > pos[course] {
				t.Errorf("Incorrect sequence: %s should come after %s ", course, precourse)
			}
		}
	}
}
