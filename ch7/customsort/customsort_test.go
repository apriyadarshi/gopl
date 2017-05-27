package customsort

import (
	"fmt"
	"sort"
	"testing"
)

func TestStateMultSort(t *testing.T) {
	tracks := []*Track{
		{"Go", "Delilah", "From the roots Up", 2012},
		{"Go", "Moby", "Moby", 1992},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007},
		{"Ready 2 Go", "Martin Solevig", "Smash", 2011},
	}
	keys := []string{"Artist", "Title", "Album"}

	ts := TrackSort{Keys: keys, Tracks: tracks}

	sort.Sort(&ts)
	fmt.Println("Sort Result")
	for i, v := range ts.Tracks {
		fmt.Printf("%d. %+v\n", i, v)
	}

}

func BenchmarkStateMultSort(b *testing.B) {

	for i := 0; i < b.N; i++ {

		tracks := []*Track{
			{"Go", "Delilah", "From the roots Up", 2012},
			{"Go", "Moby", "Moby", 1992},
			{"Go Ahead", "Alicia Keys", "As I Am", 2007},
			{"Ready 2 Go", "Martin Solevig", "Smash", 2011},
		}
		keys := []string{"Artist", "Title", "Album"}

		ts := TrackSort{Keys: keys, Tracks: tracks}

		sort.Sort(&ts)
	}
}

func TestStableSort(t *testing.T) {
	tracks := []*Track{
		{"Go", "Delilah", "From the roots Up", 2012},
		{"Ready 2 Go", "Martin Solevig", "Smash", 2011},
		{"Go", "Moby", "Moby", 1992},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007},
	}
	keys := []string{"Album"}

	ts := TrackSort{Keys: keys, Tracks: tracks}

	sort.Stable(&ts)
	ts.Keys = []string{"Title"}
	sort.Stable(&ts)
	ts.Keys = []string{"Artist"}
	sort.Stable(&ts)

	fmt.Println("Stable Sort Result")
	for i, v := range ts.Tracks {
		fmt.Printf("%d. %+v\n", i, v)
	}

}

func BenchmarkStableSort(b *testing.B) {

	for i := 0; i < b.N; i++ {

		tracks := []*Track{
			{"Go", "Delilah", "From the roots Up", 2012},
			{"Ready 2 Go", "Martin Solevig", "Smash", 2011},
			{"Go", "Moby", "Moby", 1992},
			{"Go Ahead", "Alicia Keys", "As I Am", 2007},
		}
		keys := []string{"Album"}

		ts := TrackSort{Keys: keys, Tracks: tracks}

		sort.Stable(&ts)
		ts.Keys = []string{"Title"}
		sort.Stable(&ts)
		ts.Keys = []string{"Artist"}
		sort.Stable(&ts)
	}

}
