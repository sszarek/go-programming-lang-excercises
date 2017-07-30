package sorting

import (
	"sort"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func ByArtist(t []*Track) sort.Interface {
	return byArtist(t)
}

type multiColumn struct {
	tracks  []*Track
	columns []string
}

func (x multiColumn) Len() int { return len(x.tracks) }
func (x multiColumn) Less(i, j int) bool {
	for _, col := range x.columns {
		first := x.tracks[i]
		second := x.tracks[j]

		switch col {
		case "Title":
			if first.Title != second.Title {
				return first.Title < second.Title
			}
		case "Artist":
			if first.Artist != second.Artist {
				return first.Artist < second.Artist
			}
		case "Album":
			if first.Album != second.Album {
				return first.Album < second.Album
			}
		case "Year":
			if first.Year != second.Year {
				return first.Year < second.Year
			}
		case "Length":
			if first.Length != second.Length {
				return first.Length < second.Length
			}
		}
	}

	return false
}
func (x multiColumn) Swap(i, j int) {
	x.tracks[i], x.tracks[j] = x.tracks[j], x.tracks[i]
}

func MultiColumn(t []*Track, columns []string) sort.Interface {
	return multiColumn{t, columns}
}
