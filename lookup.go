package main

import (
	"fmt"
	"path/filepath"
	"time"
)

var (
	esc   = string(byte(27))
	blue  = esc + "[34m"
	red   = esc + "[31m"
	reset = esc + "[0m"
)

type FileEntry struct {
	Name string // File name
}

type DirEntry struct {
	Files []*FileEntry // Files in the directory
	Path  string       // Absolute path of the directory
}

// lookup prints all filenames matching the pattern
func lookup(data []*DirEntry, query string) {
	elapsedTime := time.Now()
	results := 0
	checked := 0

	for _, dir := range data {

		match, err := filepath.Match(query, filepath.Base(dir.Path))
		fatal(err)
		if match {
			fmt.Println(blue + dir.Path + string(filepath.Separator) + reset)
			results++
		}
		checked++

		for _, file := range dir.Files {
			match, err := filepath.Match(query, file.Name)
			fatal(err)
			if match {
				fmt.Println(filepath.Join(dir.Path, red+file.Name+reset))
				results++
			}
			checked++
		}
	}
	fmt.Println("Results:", results, "/", checked, "Time:", time.Since(elapsedTime))
}
