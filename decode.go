package main

import (
	"encoding/gob"
	"os"
)

// decode loads index fie into ram
func decode(filename string) []*DirEntry {
	file, err := os.Open(filename)
	fatal(err)
	defer file.Close()

	data := make([]*DirEntry, 0)
	enc := gob.NewDecoder(file)
	err = enc.Decode(&data)
	fatal(err)

	return data
}
