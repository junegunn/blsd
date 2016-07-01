package main

import (
	"fmt"
	"os"
	"path"
)

var blacklist map[string]bool

func init() {
	blacklist = make(map[string]bool)
	// Deal with it
	for _, name := range []string{".git", "target", "node_modules"} {
		blacklist[name] = true
	}
}

func ignore(name string) bool {
	_, contained := blacklist[name]
	return contained
}

func isDir(name string) bool {
	f, err := os.Open(name)
	if err != nil {
		return false
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return fi.Mode().IsDir()
}

func bfsd(queue []string) []string {
	newQueue := []string{}
	for _, dir := range queue {
		if ignore(dir) {
			continue
		}

		f, err := os.Open(dir)
		if err != nil {
			continue
		}

		fis, err := f.Readdir(-1)
		if err != nil {
			f.Close()
			continue
		}
		f.Close()

		for _, fi := range fis {
			name := fi.Name()
			if fi.Mode().IsDir() && !ignore(name) {
				path := path.Join(dir, name)
				fmt.Println(path)
				newQueue = append(newQueue, path)
			}
		}
	}
	return newQueue
}

func main() {
	var queue []string
	if len(os.Args) == 1 {
		queue = []string{"."}
	} else {
		for _, name := range os.Args[1:] {
			if isDir(name) {
				fmt.Println(name)
				queue = append(queue, name)
			}
		}
	}
	for len(queue) > 0 {
		queue = bfsd(queue)
	}
}
