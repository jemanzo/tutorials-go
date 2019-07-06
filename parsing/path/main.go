package main

import (
	"fmt"
	"path"
	"strings"
)

func main() {
	fmt.Println()
	DirtyCleanPaths()
	DirFileExtension()
	MatchPattern()
	JoinAndSplit()
}

func DirtyCleanPaths() {
	fmt.Println(" -- DirtyCleanPaths ---------")
	defer fmt.Printf(" %s\n\n", strings.Repeat("-", 30))

	dirtyPath := "/my//Folder    Name///main.doc.go   "
	trimPath := strings.TrimSpace(dirtyPath)
	fmt.Printf(" %-12s : %q\n", "DirtyPath", dirtyPath)
	fmt.Printf(" %-12s : %q\n", "CleanPath1", path.Clean(dirtyPath))
	fmt.Printf(" %-12s : %q\n", "TrimPath", trimPath)
	fmt.Printf(" %-12s : %q\n", "CleanPath2", path.Clean(trimPath))
}

func DirFileExtension() {
	fmt.Println(" -- DirFileExtension --------")
	defer fmt.Printf(" %s\n\n", strings.Repeat("-", 30))

	fullPath := "/user//folder1/myfile1.doc.pdf"
	fmt.Printf(" %-12s : %q\n", "fullPath", fullPath)
	fmt.Printf(" %-12s : %q\n", "path.Dir()", path.Dir(fullPath))
	fmt.Printf(" %-12s : %q\n", "path.Base()", path.Base(fullPath))
	fmt.Printf(" %-12s : %q\n", "path.Ext()", path.Ext(fullPath))

	filename := path.Base(fullPath)
	fileNoExt := filename[:len(filename)-len(path.Ext(filename))]
	fmt.Printf(" %-12s : %q\n", "fileNoExt", fileNoExt)
}

func MatchPattern() {
	fmt.Println(" -- MatchPattern -------------")
	defer fmt.Printf(" %s\n\n", strings.Repeat("-", 30))

	pattern := "[/]*[/]*[/]main.go"
	// pattern := "[a-z0-9]*"
	fname := "/my/path/main.go"
	match, err := path.Match(pattern, fname)
	fmt.Printf(" %-12s : Pattern %q Path %q Match %v Error %v\n", "Match", pattern, fname, match, err)
}

func JoinAndSplit() {
	fmt.Println(" -- JoinAndSplit -------------")
	defer fmt.Printf(" %s\n\n", strings.Repeat("-", 30))

	joinPath := path.Join("~/Documents//", "/myfile.doc.pdf")
	splitDir, splitFile := path.Split(joinPath)

	fmt.Printf(" %-12s : %q\n", "Join", joinPath)
	fmt.Printf(" %-12s : Dir %q File %q\n", "Split", splitDir, splitFile)
	fmt.Printf(" %-12s : %v and %v\n", "IsAbs", path.IsAbs("/a/b/c/main.go"), path.IsAbs("~/Documents"))
}
