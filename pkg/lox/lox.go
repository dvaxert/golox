package lox

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func RunFile(path string) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.Buffer{}
	_, err = buf.ReadFrom(f)
	if err != nil {
		log.Fatal(err)
	}

	if err = f.Close(); err != nil {
		log.Fatal(err)
	}

	Run(buf.String())
}

func RunPromt() {
	scanner := bufio.NewScanner(os.Stdin) // TODO
	writer := bufio.NewWriter(os.Stdout)

	for {
		_, err := writer.WriteString("> ")
		if err != nil {
			log.Fatal(err)
		}

		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		Run(line)
	}
}

func Run(s string) {
	scanner := new(Scanner)
	tokens, err := scanner.ScanTokens(s)

	if err != nil {
		log.Fatal(err)
	}

	for t := range tokens {
		fmt.Println(t)
	}
}
