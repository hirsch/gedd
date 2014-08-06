package main

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fatal(errors.New("no search query"))
	}

	query := strings.Join(os.Args[1:], " ") // Only for Windows?
	fmt.Println("Searching for:", query)

	userdata, err := user.Current()
	fatal(err)
	wd, err := os.Getwd()
	fatal(err)

	savedir := userdata.HomeDir
	scandir := filepath.VolumeName(wd)
	if filepath.VolumeName(savedir) != scandir {
		savedir = scandir
	}
	savedir += string(filepath.Separator) + "index.fed"

	data := decode(savedir)
	lookup(data, query)
}

func fatal(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
