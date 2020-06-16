package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	ff := flag.String("file", "", "Name of the .txt file (including extension) to be read.")

	df := flag.String("dir", ".", "The directory to read text files from.")

	flag.Parse()

	var posts []string

	if *ff != "" {
		posts[0] = *ff
	} else {
		files, err := ioutil.ReadDir(*df)
		if err != nil {
			fmt.Print(err)
		} else {
			for _, f := range files {
				if f.IsDir() {
					subfiles, err := ioutil.ReadDir(f.Name())
					if err != nil {
						fmt.Print(err)
					} else {
						for _, sub := range subfiles {
							files = append(files, sub)
						}
					}
				} else if strings.HasSuffix(f.Name(), ".txt") {
					fmt.Println(f.Name())
					posts = append(posts, f.Name())
				}
			}
		}
	}

	for _, post := range posts {
		content := readFile(post)
		newName := strings.Split(post, ".")[0] + ".html"
		writeFile(newName, content)
	}
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
