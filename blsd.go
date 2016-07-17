package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/libgit2/git2go"
)

var blacklist map[string]bool

func init() {
	blacklist = make(map[string]bool)
	// Deal with it
	for _, name := range []string{".git", ".svn", ".hg"} {
		blacklist[name] = true
	}
}

func ignore(name string, repo *git.Repository) bool {
	_, contained := blacklist[name]
	if contained {
		return true
	} else if repo != nil {
		abs, err := filepath.Abs(name)
		if err != nil {
			return false
		}
		base := filepath.Clean(repo.Path() + "..")
		if abs == base {
			return false
		}
		rel, err := filepath.Rel(base, abs)
		if err != nil {
			return false
		}
		ignored, err := repo.IsPathIgnored(rel)
		return err == nil && ignored
	}
	return false
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
		repo, err := git.OpenRepository(dir)
		ignored := ignore(dir, repo)
		if err == nil {
			defer repo.Free()
		}
		if ignored {
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
			path := path.Join(dir, name)
			if fi.Mode().IsDir() && !ignore(path, repo) {
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
