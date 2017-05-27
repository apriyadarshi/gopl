package main

import (
	"encoding/json"
	"fmt"
	"gopl/ch7/customsort"
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"
	"sync"
)

type Result struct {
	Field string `json:Field`
}

var mu sync.Mutex
var ts customsort.TrackSort

func main() {
	tracks := []*customsort.Track{
		{"Go", "Delilah", "From the roots Up", 2012},
		{"Go", "Moby", "Moby", 1992},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007},
		{"Ready 2 Go", "Martin Solevig", "Smash", 2011},
	}
	keys := []string{}

	ts = customsort.TrackSort{Keys: keys, Tracks: tracks}

	http.HandleFunc("/tracks", handler)
	http.HandleFunc("/tracksort", handlerSort)
	log.Fatal(http.ListenAndServe("192.168.1.3:8080", nil)) //nil => DefaultServeMux handler handles the function
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprint(w, `
		<html>
			<head>
				<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
			</head>
			<body>
				<h1>Tracks</h1>
				<div id="tableDiv">
			`)
	FPrintTracksHTML(w, ts)
	fmt.Fprint(w, `
			</div>
			<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
				<script>
					var sort = function(column) {
						var xhr = new XMLHttpRequest();
						xhr.open('POST', 'http://192.168.1.3:8080/tracksort');
						//xhr.setRequestHeader('Content-Type', 'application/json');
						xhr.onload = function() {
				    			if (xhr.status === 200) {
				         		document.getElementById("tableDiv").innerHTML = xhr.responseText
				    		}
						};

						xhr.send(JSON.stringify({
							Field: column
						}));
					}
				</script>
			</body>
		</html>`)
	mu.Unlock()
}

func handlerSort(w http.ResponseWriter, r *http.Request) {
	var result Result
	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		fmt.Println(err)
	}
	mu.Lock()
	insertTSKey(result.Field)
	sort.Sort(&ts)
	FPrintTracksHTML(w, ts)
	mu.Unlock()
}

func insertTSKey(k string) {
	if ts.Keys == nil || len(ts.Keys) == 0 {
		ts.Keys = append(ts.Keys, k)
		return
	}
	newSlice := make([]string, 3, 3)
	newSlice[0] = k
	for i, key := range ts.Keys {
		if i > 1 {
			break
		}
		newSlice[i+1] = key
	}
	ts.Keys = newSlice
}

func FPrintTracksHTML(w io.Writer, trs customsort.TrackSort) error {

	var tracklist = template.Must(template.New("tracklist").Parse(`
	<table class ="table table-bordered">
		<tr>
			<th onclick="sort('Title')">Title ↑↓</th>
			<th onclick="sort('Artist')">Artist ↑↓</th>
			<th onclick="sort('Album')">Album ↑↓</th>
			<th onclick="sort('Year')">Year ↑↓</th>
		</tr>
		{{range .Tracks}}
		<tr>
			<td>{{.Title}}</td>
			<td>{{.Artist}}</td>
			<td>{{.Album}}</td>
			<td>{{.Year}}</td>
		</tr>
		{{end}}
	</table>
	`))

	if err := tracklist.Execute(w, trs); err != nil {
		return err
	}
	return nil
}
