package main

import (
	"github.com/sszarek/go-programming-lang-excercises/c07/e7.8/sorting"
	"time"
	"net/http"
	"html/template"
	"sort"
	"strings"
)

var report = template.Must(template.New("tracks").Parse(`
		<html>
			<head>
				<title>Tracks</title>
			</head>
			<body>
			<table>
				<thead>
					<tr>
						<th><a href="/?sort=Title">Title</a></th>
						<th><a href="/?sort=Artist">Artist</a></th>
						<th><a href="/?sort=Album">Album</a></th>
						<th><a href="/?sort=Year">Year</a></th>
						<th><a href="/?sort=Length">Length</a></th>
					</tr>
				</thead>
				<tbody>
				{{range .Tracks}}
					<tr>
						<td>{{.Title}}</td>
						<td>{{.Artist}}</td>
						<td>{{.Album}}</td>
						<td>{{.Year}}</td>
						<td>{{.Length}}</td>
					</tr>
				{{end}}
				</tbody>
			</table>
			</body>
		</html>
		`))

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, request *http.Request) {
		var tracks = []*sorting.Track{
			{"Go", "Delilah", "From the roots up", 2012, length("3m38s")},
			{"Go", "Moby", "Moby", 1992, length("3m37s")},
			{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
			{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		}

		query := request.URL.Query()
		col := query.Get("sort")
		if col != "" {
			sort.Sort(sorting.MultiColumn(tracks, strings.Split(col, ",")))
		}

		report.Execute(resp, struct {
			Tracks []*sorting.Track
		}{ tracks})
	})
	http.ListenAndServe(":8080", nil)
}
