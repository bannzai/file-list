package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func contains(array []string, s string) bool {
	for _, v := range array {
		if s == v {
			return true
		}
	}
	return false
}

func parseFileContent(content []byte) []string {
	return split(string(content), "\n")
}

func split(content string, separator string) []string {
	return strings.Split(content, separator)
}

func convertToFileNames(argument string) []string {
	data, _ := ioutil.ReadFile(argument)
	if data != nil {
		return parseFileContent(data)
	}
	return split(argument, ",")
}

func main() {
	var (
		ignore string
		only   string
	)

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fs.StringVar(&ignore, "ignore", "", "Do not output file-list")
	fs.StringVar(&only, "only", "", "Show only file-list.")
	fs.Parse(os.Args[2:])

	fmt.Printf("os.Args() = %+v\n", os.Args)
	fmt.Printf("ignore = %+v\n", ignore)
	fmt.Printf("only = %+v\n", only)

	if len(os.Args) < 2 {
		panic("Should specifity directory. e.g file-list /path/to/directory")
	}

	directory, err := ioutil.ReadDir(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("Can not read directory: %v", os.Args[1]))
	}

	ignoreFiles := convertToFileNames(ignore)
	onlyFiles := convertToFileNames(only)
	hasIgnore := len(ignore) > 0
	hasOnly := len(only) > 0

	fmt.Printf("ignoreFiles = %+v\n", ignoreFiles)
	fmt.Printf("hasIgnore = %+v\n", hasIgnore)
	fmt.Printf("hasOnly = %+v\n", hasOnly)
	if hasIgnore && hasOnly {
		panic("Can not be shared option. --ignore and --only")
	}

	var files []string
	for _, f := range directory {
		fileName := f.Name()

		if contains(ignoreFiles, fileName) {
			continue
		}

		if hasOnly && contains(onlyFiles, fileName) {
			files = append(files, fileName)
			continue
		}

		files = append(files, fileName)
	}

	fmt.Printf("files = %+v\n", files)
}
