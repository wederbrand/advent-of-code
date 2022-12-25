package main

import (
	"fmt"
	"github.com/wederbrand/advent-of-code/github.com/wederbrand/advent-of-code/util"
	"regexp"
	"strconv"
)

type dir struct {
	name      string
	parent    *dir
	sub       map[string]*dir
	files     map[string]int
	totalSize int
}

func newDir(name string) *dir {
	d := new(dir)
	d.name = name
	d.sub = make(map[string]*dir)
	d.files = make(map[string]int)
	return d
}

func (d *dir) calculateAllSizes() {
	if d.totalSize != 0 {
		return
	}
	for _, d2 := range d.sub {
		d2.calculateAllSizes()
		d.totalSize += d2.totalSize
	}

	for _, s := range d.files {
		d.totalSize += s
	}
}

func (d *dir) getAllLessThan(maxSize int) []*dir {
	all := make([]*dir, 0)
	if d.totalSize <= maxSize {
		all = append(all, d)
	}
	for _, d2 := range d.sub {
		all = append(all, d2.getAllLessThan(maxSize)...)
	}

	return all
}

func (d *dir) getAll() (result []*dir) {
	result = append(result, d)

	for _, d2 := range d.sub {
		result = append(result, d2.getAll()...)
	}

	return
}

func main() {
	inFile := util.GetFileContents("2022/07/input.txt", "\n")

	root := newDir("/")

	regexCd := regexp.MustCompile("\\$ cd (.+)")
	regexLs := regexp.MustCompile("\\$ ls")
	regexDir := regexp.MustCompile("dir (.+)")
	regexFile := regexp.MustCompile("(.+) (.+)")
	current := root
	for i := 0; i < len(inFile); i++ {
		s := inFile[i]
		match := regexCd.FindStringSubmatch(s)
		if match != nil {
			// handle cd
			d := match[1]

			if d == "/" {
				current = root
			} else if d == ".." {
				current = current.parent
			} else {
				current = current.sub[d]
			}

			continue
		}

		match = regexLs.FindStringSubmatch(s)
		if match != nil {
			// handle ls (do nothing)
			continue
		}

		match = regexDir.FindStringSubmatch(s)
		if match != nil {
			// handle dir
			d := match[1]
			_, found := current.sub[d]
			if !found {
				current.sub[d] = newDir(d)
				current.sub[d].parent = current
			}
			continue
		}

		match = regexFile.FindStringSubmatch(s)
		if match != nil {
			// handle file
			f := match[2]
			size, _ := strconv.Atoi(match[1])
			current.files[f] = size
			continue
		}
	}

	root.calculateAllSizes()
	all := root.getAllLessThan(100000)

	total := 0
	for _, d := range all {
		total += d.totalSize
	}

	fmt.Println("part 1:", total)

	freeUp := 30_000_000 - (70_000_000 - root.totalSize)

	all = root.getAll()

	smallest := root
	for _, d := range all {
		if d.totalSize < smallest.totalSize && d.totalSize >= freeUp {
			smallest = d
		}
	}

	fmt.Println("part 2:", smallest.totalSize)
}
