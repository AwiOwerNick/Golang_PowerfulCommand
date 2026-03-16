package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	lines := flag.Bool("l", false, "Count lines")
	bytesCount := flag.Bool("b", false, "Count bytes")
	flag.Parse()
	if flag.NArg() <= 1 {
		fmt.Println(count(os.Stdin, *lines, *bytesCount))
	} else {
		fileNames := flag.Args()
		for _, v := range fileNames {
			input, err := ioutil.ReadFile(v)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			reader := bytes.NewReader(input)
			fmt.Printf("File path: %s\n", v)
			fmt.Println(count(reader, *lines, *bytesCount))
		}
	}

}

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}
	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}

	return wc
}
