package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func contains(array []string, s string) bool {
	for _, v := range array {
		if s == v {
			return true
		}
	}
	return false
}

func parseFileContent(filePath string) []string {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(errors.Wrap(err, "Failed read file"))
	}
	return split(string(data), "\n")
}

func split(content string, separator string) []string {
	return strings.Split(content, separator)
}

func main() {
	var (
		ignoreWithFile string
		ignoreList     string
		onlyWithFile   string
		onlyList       string
	)

	flag.Usage = func() {
		name := os.Args[0]
		fmt.Fprintf(
			os.Stderr, `
Usage of %s: 
%s [DIRECTORY] ARGS...
Options: 
	ignore-list                 "Do not output file-list"
	only-list                   "Show only file-list"
	ignore-with-file            "Do not output file-list"
	only-with-file              "Show only file-list" 
	 `,
			name,
			name,
		)
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	fs.StringVar(&ignoreList, "ignore-list", "", "Do not output file-list")
	fs.StringVar(&onlyList, "only-list", "", "Show only file-list.")
	fs.StringVar(&ignoreWithFile, "ignore-with-file", "", "Do not output file-list")
	fs.StringVar(&onlyWithFile, "only-with-file", "", "Show only file-list.")
	fs.Parse(os.Args[2:])

	if len(os.Args) < 2 {
		panic("Should specifity directory. e.g file-list /path/to/directory")
	}

	var ignore []string
	var only []string

	if len(ignoreList) > 0 {
		ignore = append(ignore, split(ignoreList, ",")...)
	}
	if len(onlyList) > 0 {
		only = append(only, split(onlyList, ",")...)
	}
	if len(ignoreWithFile) > 0 {
		ignore = append(ignore, parseFileContent(ignoreWithFile)...)
	}
	if len(onlyWithFile) > 0 {
		only = append(only, parseFileContent(onlyWithFile)...)
	}

	hasIgnore := len(ignore) > 0
	hasOnly := len(only) > 0

	if hasIgnore && hasOnly {
		panic("Can not be shared option. --ignore and --only")
	}

	dirName := os.Args[1]

	var files []string
	err := filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("Can not read directory: %v, info: %v", dirName, info))
		}

		dir := filepath.Dir(path)
		if dir != dirName {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		fileName := filepath.Join(dir, info.Name())
		if contains(ignore, fileName) {
			return nil
		}

		if !hasOnly {
			files = append(files, fileName)
			return nil
		}

		if contains(only, fileName) {
			files = append(files, fileName)
			return nil
		}

		return nil
	})

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Printf("%+v\n", file)
	}
}
