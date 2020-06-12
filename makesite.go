package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type entry struct {
	Name string
	Done bool
}

type ToDo struct {
	User string
	List []entry
}

func main() {
	sp := flag.String("file", "first-post.txt", "Name of the .txt file (including extension) to be read.")
	flag.Parse()

	post := readFile(*sp)

	newName := strings.Split(*sp, ".")[0] + ".html"

	writeFile(newName, post)
}

func readFile(fileName string) string {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func writeFile(fileName string, text string) {
	// Files are provided as a slice of strings.
	paths := []string{
		"template.tmpl",
	}

	t := template.Must(template.New("template.tmpl").ParseFiles(paths...))
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	err = t.Execute(file, text)
	if err != nil {
		panic(err)
	}
}
