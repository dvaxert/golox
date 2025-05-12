package lox

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/dvaxert/golox/pkg/lox/scanner"
)

func NewLox() *Lox {
	return &Lox{
		had_error: false,
	}
}

//----------------------------------------------------------------------------------------------------------------------

type Lox struct {
	had_error bool
}

//----------------------------------------------------------------------------------------------------------------------

func RunFile(path string) {
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("error when opening a file: %v", err)
		return
	}

	buf := bytes.Buffer{}
	_, err = buf.ReadFrom(f)
	if err != nil {
		fmt.Printf("file reading error: %v", err)
		return
	}

	if err = f.Close(); err != nil {
		fmt.Printf("file closing error: %v", err)
		return
	}

	lox := NewLox()

	errs := lox.Run([]rune(buf.String()))
	if len(errs) != 0 {
		for _, err := range errs {
			os.Stderr.WriteString(err.Error())
		}
		os.Exit(1)
	}
}

//----------------------------------------------------------------------------------------------------------------------

func RunPromt() {
	reader := bufio.NewReader(os.Stdin)
	lox := NewLox()

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')

		if len(line) > 0 {
			errs := lox.Run([]rune(line))
			if len(errs) != 0 {
				for _, e := range errs {
					os.Stderr.WriteString(fmt.Sprintf("%s\n", e.Error()))
				}
			}
		}

		if err != nil {
			if err == io.EOF {
				os.Stdout.WriteString("bye bye")
				break
			}

			log.Fatal(err)
		}
	}
}

//----------------------------------------------------------------------------------------------------------------------

func (l *Lox) Run(s []rune) []error {
	scanner := scanner.New(s)
	tokens, err := scanner.ScanTokens()

	if err != nil {
		return err
	}

	for _, t := range tokens {
		os.Stdout.WriteString(fmt.Sprintf("%s\n", t.String()))
	}

	return nil
}
