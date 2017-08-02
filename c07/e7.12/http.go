package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"sync"
)

var itemsTemplate = template.Must(template.New("items").Parse(`
<html>
			<head>
				<title>Tracks</title>
			</head>
			<body>
			<table>
				<thead>
					<tr>
						<th>Name</th>
						<th>Price</th>
						<th> </th>
					</tr>
				</thead>
				<tbody>
				{{range $key, $value := .}}
					<tr>
						<td>{{ $key}}</td>
						<td>{{ $value}}</td>
						<td><a href="/delete?item={{ $key }}">Delete</a></td>
					</tr>
				{{end}}
				</tbody>
			</table>
			</body>
		</html>
`))

func main() {
	db := &database{data: map[string]dollars{"buty": 50, "skarpety": 20}}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/add", db.add)
	mux.HandleFunc("/delete", db.delete)
	http.ListenAndServe(":8080", mux)
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("%.2f $", d) }

type database struct {
	sync.Mutex
	data map[string]dollars
}

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	itemsTemplate.Execute(w, db.data)
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db.data[item]
	if !ok {
		writeNotFound(w, item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db.data[item]; !ok {
		writeNotFound(w, item)
		return
	}

	priceStr := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid price value: %s\n", priceStr)
	}

	db.Lock()
	db.data[item] = dollars(price)
	db.Unlock()
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db.data[item]; !ok {
		writeNotFound(w, item)
		return
	}

	db.Lock()
	delete(db.data, item)
	db.Unlock()
}

func (db *database) add(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db.data[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Entry already exists: %s\n", item)
		return
	}

	priceStr := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid price value: %s\n", priceStr)
	}

	db.Lock()
	db.data[item] = dollars(price)
	db.Unlock()
}

func writeNotFound(w http.ResponseWriter, item string) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "There is not such entry: %q\n", item)
}
