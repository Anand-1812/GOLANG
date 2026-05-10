package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

type Track struct {
	Title string
	Artist string
	Album string
	Year int
}

type byArtist []*Track

func (ba byArtist) Len() int {
	return len(ba)
}

func (ba byArtist) Less(i, j int) bool {
	return ba[i].Artist < ba[j].Artist
}

func (ba byArtist) Swap(i, j int) {
	ba[i], ba[j] = ba[j], ba[i]
}

func (s StringSlice) Len() int {
	return len(s)
}

func (s StringSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s StringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	
	names := []string{
		"Zoro",
		"Luffy",
		"Anand",
	}

	tracks := []*Track{
		{"Go", "Delilah", "Roots", 2012},
		{"Go", "Moby", "Mpby", 1992},
	}

	fmt.Println("--- Basic example ---")
	sort.Sort(StringSlice(names))
	fmt.Println(names)

	fmt.Println("-- Real world example ---")
	sort.Sort(byArtist(tracks))

	for _, t := range tracks {
		fmt.Printf(
			"Title: %-5s Artist: %-10s Album: %-5s Year: %d\n",
			t.Title,
			t.Artist,
			t.Album,
			t.Year,
		)
	}
}

