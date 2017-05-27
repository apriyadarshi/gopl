package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var rwMu sync.RWMutex

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {

	var template = template.Must(template.New("dblist").Parse(`
		<html>
			<head></head>
			<body>
				<h1>Item List</h1>
				<table>
					<thead>
						<tr>
							<th>Name</th>
							<th>Price</th>
						</tr>
					</thead>
					<tbody>
						{{range $k,$v := .}}
						<tr>
							<td>{{$k}}</td>
							<td>{{$v}}</td>
						</tr>
						{{end}}
					</tbody>
					<tfoot>
					</tfoot>
				</table>
			</body>
		</html>`))

	rwMu.RLock()
	if err := template.Execute(w, db); err != nil {
		http.Error(w, err.Error(), http.StatusExpectationFailed)
	}
	rwMu.RUnlock()
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	rwMu.RLock()
	price, ok := db[item]
	rwMu.RUnlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

/* /create?item=shirt&price=24
 */
func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	p, err := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
	if err != nil {
		http.Error(w, "strconv: price in url:"+err.Error(), http.StatusBadRequest)
	}
	if item == "" {
		http.Error(w, "empty item in url", http.StatusBadRequest)
	}
	if _, ok := db[item]; ok {
		http.Error(w, "item already exists", http.StatusBadRequest)
	}

	rwMu.Lock()
	db[item] = dollars(p)
	w.WriteHeader(http.StatusOK)
	rwMu.Unlock()
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	p, err := strconv.ParseFloat(req.URL.Query().Get("price"), 32)
	if err != nil {
		http.Error(w, "strconv: price in url:"+err.Error(), http.StatusBadRequest)
	}
	if item == "" {
		http.Error(w, "empty item in url", http.StatusBadRequest)
	}
	if _, ok := db[item]; !ok {
		http.Error(w, "item doesn't exist", http.StatusBadRequest)
	}

	rwMu.Lock()
	db[item] = dollars(p)
	w.WriteHeader(http.StatusOK)
	rwMu.Unlock()
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		http.Error(w, "empty item in url", http.StatusBadRequest)
	}
	if _, ok := db[item]; !ok {
		http.Error(w, "item doesn't exist", http.StatusBadRequest)
	}

	rwMu.Lock()
	delete(db, item)
	w.WriteHeader(http.StatusNoContent)
	rwMu.Unlock()
}
