package main

import "time"
import "text/tabwriter"
import "os"
import (
	"fmt"
	"sort"

	"github.com/sszarek/go-programming-lang-excercises/c07/e7.8/sorting"
)

var tracks = []*sorting.Track{
	{"Go", "Delilah", "From the roots up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*sorting.Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Album", "Artist", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "------", "----", "------")
	sort.Sort(sorting.MultiColumn(tracks, []string{"Title", "Year"}))
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Album, t.Artist, t.Year, t.Length)
	}
	tw.Flush()
}

func main() {
	printTracks(tracks)
}
