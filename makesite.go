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

	if *ff != "" {
		makePost(*ff)
	} else {
		parseDir(*df)
	}
}

func parseDir(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Print("Error reading files: ")
		fmt.Println(err)
	} else {
		for _, f := range files {
			if f.IsDir() {
				parseDir(fmt.Sprintf("%s/%s", dir, f.Name()))
			} else if strings.HasSuffix(f.Name(), ".txt") {
				fmt.Println(f.Name())
				fmt.Println(dir + "/" + f.Name())
				makePost(dir + "/" + f.Name())
			}
		}
	}
}

func makePost(name string) {
	content := readFile(name)
	newName := strings.Split(name, ".txt")[0] + ".html"
	writeFile(newName, content)
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
