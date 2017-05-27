package customsort

import (
	//"fmt"
	"reflect"
	"sort"
)

/*Stateful Multi-tier sort. 3 states - , , third last clicked
 *First key - last clicked
 *Secondary key - second last clicked
 *Tertiary key - third last clicked
 */
type SMSort interface {
	sort.Interface
	SetKeyOrder([]string) //sets the number of keys provided
	getKeyOrder() []string
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
}

type TrackSort struct {
	Keys   []string
	Tracks []*Track
}

func (ts *TrackSort) Len() int {
	return len(ts.Tracks)
}

func (ts *TrackSort) Less(i, j int) bool {
	for _, f := range ts.Keys {
		v1 := reflect.Indirect(reflect.ValueOf(ts.Tracks[i])).FieldByName(f)
		v2 := reflect.Indirect(reflect.ValueOf(ts.Tracks[j])).FieldByName(f)
		switch v1.Interface().(type) {
		case int:
			if v1.Int() != v2.Int() {

				return v1.Int() < v2.Int()
			}
		case string:
			if v1.String() != v2.String() {
				return v1.String() < v2.String()
			}
		}
	}
	return false

}

func (ts *TrackSort) Swap(i, j int) {
	ts.Tracks[i], ts.Tracks[j] = ts.Tracks[j], ts.Tracks[i]
}
