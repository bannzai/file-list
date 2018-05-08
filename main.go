package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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

}

func convertToFileNames(argument string) []string {
	data, err := ioutil.ReadFile(argument)
	if err == nil && data != nil {
		return
	}

}

func main() {
	var (
		fs     = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		ignore = fs.String("ignore", "", "Do not output file-list")
		only   = fs.String("only", "", "Show only file-list.")
	)

	fs.Parse(os.Args[2:])
	fmt.Printf("os.Args() = %+v\n", os.Args)

	if len(os.Args) < 2 {
		panic("Should specifity directory. e.g file-list /path/to/directory")
	}

	fmt.Printf("ignore = %+v\n", *ignore)
	fmt.Printf("only = %+v\n", *only)

	// directory := os.Args[0]

	cmd := exec.Command("find", "./", "-maxdepth", "1", "-type", "f", "-not", "-name", "'user.go'")
	fmt.Printf("cmd = %+v\n", cmd)
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Printf("output = %+v\n", string(output))
}
