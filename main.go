package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

func main() {
	countBytes := flag.Bool("c", false, "The number of bytes in the input file")
	countLines := flag.Bool("l", false, "The number of lines in the input file")
	countWords := flag.Bool("w", false, "The number of words in the input file")
	countChars := flag.Bool("m", false, "The number of words in the input file")
	flag.Parse()

	var fileData []byte
	if filename := flag.Arg(0); filename != "" {
		f, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println("error opening the file: ", err)
			os.Exit(1)
		}
		fileData = f
	} else {
		fmt.Println("Please give a filename")
		return
	}

	if *countBytes {
		fmt.Println(len(fileData))
	} else if *countLines {
		fmt.Println(bytes.Count(fileData, []byte("\n")))
	} else if *countWords {
		fmt.Println(len(bytes.Fields(fileData)))
	} else if *countChars {
		fmt.Println(len(bytes.Runes(fileData)))
	}
}
