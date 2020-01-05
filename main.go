package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	file_name := os.Args[1]
	fmt.Println("is ths working")
	fmt.Println(file_name)

	file, err := os.Open(file_name)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	input, err := ioutil.ReadFile(file_name)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
	for i, line := range lines {
		strTrim := strings.Trim(line, "  \t")
		strTrimRight := strings.TrimRight(line, "  \t")
		if len(strTrim) == 0 {
			lines[i] = ""
		} else if len(strTrimRight) != len(line) {
			lines[i] = strTrimRight
		}
	}

	output := strings.Join(lines, "\n")
	if err = ioutil.WriteFile(file_name, []byte(output), 0644); err != nil {
		log.Fatal(err)
	}
}
