package main

// Example: go generate >example/myset.go

//go:generate go run bitset.go MySet MyType example int

import (
	"flag"
	"fmt"
	"os"
	"text/template"
)

func report(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(2)
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: ibitset Name Type PackageName Underlying\n")
	flag.PrintDefaults()
	os.Exit(2)
}

type BitSet struct {
	Name       string
	BaseType   string
	Underlying string
	Package    string
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 4 {
		usage()
	}
	var b BitSet
	b.Name = args[0]
	b.BaseType = args[1]
	b.Package = args[2]
	b.Underlying = args[3]
	t, err := template.ParseFiles("set-template.txt")
	if err != nil {
		report(err)
	}
	err = t.Execute(os.Stdout, b)
	if err != nil {
		report(err)
	}
}
