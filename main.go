package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	lineNumber bool
)

func init() {
	flag.BoolVar(&lineNumber, "b", false, "number noempty output lines, overrides -n")
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Printf("invalid args count %d\n", len(args))
		os.Exit(1)
	}

	if err := run(args); err != nil {
		log.Fatal(err)
	}
}

func run(args []string) error {
	f, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}

	var res string
	if lineNumber {
		res = readAllWithLineNumber(f)
	} else {
		res, err = readAll(f)
		if err != nil {
			return err
		}
	}
	f.Close()

	fmt.Print(res)
	return nil
}

func readAll(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func readAllWithLineNumber(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	var lines []string
	lineNumber := 1
	for scanner.Scan() {
		lines = append(lines, fmt.Sprintf("%d\t", lineNumber)+scanner.Text())
		lineNumber++
	}
	return strings.Join(lines, "\n") + "\n"
}
