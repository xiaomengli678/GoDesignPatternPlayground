package main

import "fmt"

type File struct {
	name      string
	extension string
}

type SearchFilter interface {
	Apply(file *File) bool
}

type NameFilter struct {
	targetName string
}

func (nf *NameFilter) Apply(f *File) bool {
	return f.name == nf.targetName
}

type ExtensionFilter struct {
	extenion string
}

func (ef *ExtensionFilter) Apply(f *File) bool {
	return f.extension == ef.extenion
}

type AndFilter struct {
	Filters []SearchFilter
}

func (af *AndFilter) Apply(f *File) bool {
	for _, filter_ := range af.Filters {
		if !filter_.Apply(f) {
			return false
		}
	}
	return true
}

type OrFilter struct {
	Filters []SearchFilter
}

func (of *OrFilter) Apply(f *File) bool {
	for _, filter_ := range of.Filters {
		if filter_.Apply(f) {
			return true
		}
	}
	return false
}

type SearchEngine struct {
	files []*File
}

func (s *SearchEngine) Search(filter SearchFilter) []*File {
	ans := []*File{}

	for _, file := range s.files {
		if filter.Apply(file) {
			ans = append(ans, file)
		}
	}
	return ans
}

func main() {
	files := []*File{
		&File{name: "report", extension: ".log"},
		&File{name: "notes", extension: ".pdf"}}
	engine := SearchEngine{files}
	nf := &NameFilter{targetName: "report"}
	ef := &ExtensionFilter{extenion: ".pdf"}
	af := &AndFilter{Filters: []SearchFilter{nf, ef}}
	of := &OrFilter{Filters: []SearchFilter{nf, ef}}
	matches_end1 := engine.Search(af)
	for _, m := range matches_end1 {
		fmt.Println(m.name, m.extension)
	}
	matches_end2 := engine.Search(of)
	for _, m := range matches_end2 {
		fmt.Println(m.name, m.extension)
	}

}
